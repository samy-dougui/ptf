package cli

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	loader2 "github.com/samy-dougui/tftest/cli/internal/loader"
	"github.com/samy-dougui/tftest/cli/internal/rule"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "tftest",
	Short: "Tftest helps you test your terraform plan",
	Long:  `Tftest is a cli tool that helps you test your terraform plan through HCL config file in a declarative manner.`,
	Run: func(cmd *cobra.Command, args []string) {
		planPath, _ := cmd.Flags().GetString("plan")
		// TODO: add global dir flag
		//dirPath, _ := cmd.Flags().GetString("dir")
		run(planPath)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().String("plan", "tfplan.json", "Terraform plan that needs to be tested")
}

func run(planPath string) {
	wr := hcl.NewDiagnosticTextWriter(
		os.Stdout,                    // writer to send messages to
		hclparse.NewParser().Files(), // the parser's file cache, for source snippets
		78,                           // wrapping width
		true,                         // generate colored/highlighted output
	)
	var diags hcl.Diagnostics
	var loader loader2.Loader
	loader.Init()

	body, diagHCLFile := loader.LoadHCLFile("./data/policy.hcl")
	plan, _ := loader.LoadPlan(planPath)
	diags = append(diags, diagHCLFile...)

	content, _ := body.Content(configFileSchema)
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
}

var configFileSchema = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type:       "rule",
			LabelNames: []string{"name"},
		},
	},
}
