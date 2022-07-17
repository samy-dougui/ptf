package control

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/samy-dougui/ptf/internal/config"
	"github.com/samy-dougui/ptf/internal/loader"
	"github.com/samy-dougui/ptf/internal/logging"
	p "github.com/samy-dougui/ptf/internal/policy"
	"github.com/samy-dougui/ptf/internal/ux"
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
	Long:  ``, // TODO: Add long description to control command

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

	policies, initDiags := p.InitPolicies(content.Blocks)
	diags = append(diags, initDiags...)

	policyDiag := p.ApplyPolicies(policies, plan)
	diags = append(diags, policyDiag...)
	ux.WriteSummary(&diags, policies)

	if diags.HasErrors() {
		os.Exit(1)
	}

}
