package cli

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	loader2 "github.com/samy-dougui/tftest/cli/internal/loader"
	"github.com/samy-dougui/tftest/cli/internal/rule"
	"github.com/samy-dougui/tftest/cli/logging"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path"
)

func init() {
	RootCmd.Flags().String("plan", "tfplan.json", "Terraform plan that needs to be tested")
	RootCmd.Flags().String("chdir", ".", "Directory where the policy files are")
}

var RootCmd = &cobra.Command{
	Use:   "tftest",
	Short: "Tftest helps you test your terraform plan",
	Long:  `Tftest is a cli tool that helps you test your terraform plan through HCL config file in a declarative manner.`,
	Run: func(cmd *cobra.Command, args []string) {
		planPath, _ := cmd.Flags().GetString("plan")
		dirPath, _ := cmd.Flags().GetString("chdir")
		run(planPath, dirPath)
	},
}

func run(planPath string, dirPath string) {
	logger := logging.GetLogger()
	logger.Info("My first logging !!")
	wr := hcl.NewDiagnosticTextWriter(
		os.Stdout,                    // writer to send messages to
		hclparse.NewParser().Files(), // the parser's file cache, for source snippets
		78,                           // wrapping width
		true,                         // generate colored/highlighted output
	)
	var diags hcl.Diagnostics
	var loader loader2.Loader
	loader.Init()

	body, diagHCLFile := loader.LoadConfigDir(path.Clean(dirPath))
	plan, diagPlanFile := loader.LoadPlan(planPath)

	diags = append(diags, diagHCLFile...)
	diags = append(diags, diagPlanFile...)

	if diags.HasErrors() {
		err := wr.WriteDiagnostics(diags)
		if err != nil {
			log.Fatal("Unexpected error")
		}
		os.Exit(1)
	}

	content, bodyDiag := body.Content(configFileSchema)
	diags = append(diags, bodyDiag...)
	logger.Info(fmt.Sprintf("Number of rules found: %v", len(content.Blocks)))
	for _, block := range content.Blocks {
		switch block.Type {
		case "rule":
			var rule rule.Rule
			ruleDiags := rule.Init(block)
			diags = append(diags, ruleDiags...)
			applyDiags := rule.Apply(plan)

			diags = append(diags, applyDiags...)

		default:
			continue
		}
	}
	errDiag := wr.WriteDiagnostics(diags)
	if errDiag != nil {
		fmt.Printf("Error while writing the diagnostics: %v", errDiag)
	}

	if diags.HasErrors() {
		os.Exit(1)
	}

}

var configFileSchema = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type:       "rule",
			LabelNames: []string{"name"},
		},
	},
}
