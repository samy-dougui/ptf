package main

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/samy-dougui/tftest/internal/loader"
	"github.com/samy-dougui/tftest/internal/rule"
	"os"
)

func main() {
	wr := hcl.NewDiagnosticTextWriter(
		os.Stdout,                    // writer to send messages to
		hclparse.NewParser().Files(), // the parser's file cache, for source snippets
		78,                           // wrapping width
		true,                         // generate colored/highlighted output
	)
	var diags hcl.Diagnostics
	var loader loader.Loader
	loader.Init()

	body, diagHCLFile := loader.LoadHCLFile("./data/policy.hcl")
	plan, _ := loader.LoadPlan("tfplan.json")
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
