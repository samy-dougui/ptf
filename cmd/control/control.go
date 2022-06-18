package control

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/samy-dougui/ptf/internal/config"
	"github.com/samy-dougui/ptf/internal/loader"
	"github.com/samy-dougui/ptf/internal/logging"
	"github.com/samy-dougui/ptf/internal/policy"
	"github.com/spf13/cobra"
	"os"
	"path"
)

var (
	planPath string
	dirPath  string
)

var ControlCmd = &cobra.Command{
	Use:   "control",
	Short: "Control your Terraform plan and Terraform state.",
	Long:  ``, // TODO: Add long description to control comdand

	Run: func(cmd *cobra.Command, args []string) {
		run(planPath, dirPath)
	},
}

func init() {
	ControlCmd.Flags().StringVarP(&planPath, "plan", "p", "tfplan.json", "Path to the Terraform plan that needs to be controlled.")
	ControlCmd.Flags().StringVarP(&dirPath, "chdir", "d", ".", "Directory where the policy files are defined.")
}

func run(planPath string, dirPath string) {
	// TODO: Tidy this function
	logger := logging.GetLogger()
	logger.Debugf("Run control command with planPath: %v and chdir: %v", planPath, dirPath)
	var diags hcl.Diagnostics
	var loader loader.Loader // TODO: Maybe add global loader and init and setup
	loader.Init()

	body, diagHCLFile := loader.LoadConfigDir(path.Clean(dirPath))
	plan, diagPlanFile := loader.LoadPlan(planPath)

	diags = append(diags, diagHCLFile...)
	diags = append(diags, diagPlanFile...)

	if diags.HasErrors() {
		err := logging.WriteDiagnostics(diags) // TODO: Log using logger not diag
		if err != nil {
			logger.Fatalf("Unexpected error while writing diagnostics: %e", err)
		}
		logger.Fatal("Error while loading the files.")
	}

	content, bodyDiag := body.Content(config.ConfigFileSchema)
	diags = append(diags, bodyDiag...)
	logger.Debugf("Number of policy found: %v", len(content.Blocks))
	for _, block := range content.Blocks {
		switch block.Type {
		case "policy":
			var policy policy.Policy
			policyDiags := policy.Init(block)
			diags = append(diags, policyDiags...)
			if !policy.IsDisabled() {
				applyDiags := policy.Apply(plan)
				diags = append(diags, applyDiags...)
			}
		default:
			continue
		}
	}
	errDiag := logging.WriteDiagnostics(diags)
	if errDiag != nil {
		logger.Errorf("Error while writing the diagnostics: %v", errDiag)
	}

	if diags.HasErrors() {
		os.Exit(1)
	}

}
