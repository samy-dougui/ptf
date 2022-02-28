package main

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/samy-dougui/tftest/internal/filter"
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
	diags = append(diags, diagHCLFile...)

	var filterCondition string
	content, _ := body.Content(configFileSchema)
	for _, block := range content.Blocks {
		switch block.Type {
		case "rule":
			partialContent, _, resourceDiags := block.Body.PartialContent(rule.BlockSchema)
			diags = append(diags, resourceDiags...)
			//attrs := partialContent.Attributes
			//for name, attr := range attrs {
			//	value, _ := attr.Expr.Value(nil)
			//	fmt.Printf("The parameters %v for the rule %v has the value %v\n", name, block.Labels[0], value.AsString())
			//}
			for _, a := range partialContent.Blocks {
				filterContent, _, _ := a.Body.PartialContent(filter.Schema)
				for name, attr := range filterContent.Attributes {
					if name == "type" {
						val, _ := attr.Expr.Value(nil)
						filterCondition = val.AsString()
					}
				}
			}
		default:
			continue
		}
	}
	errDiag := wr.WriteDiagnostics(diags)
	if errDiag != nil {
		fmt.Printf("Error while writing the diagnostics: %v", errDiag)
	}

	plan, _ := loader.LoadPlan("tfplan.json")
	for _, ressourceChange := range plan.ResourceChanges {
		if ressourceChange.Type == filterCondition {
			fmt.Println(ressourceChange.Address)
		}
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
