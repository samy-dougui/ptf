package rule

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	loader2 "github.com/samy-dougui/tftest/cli/internal/loader"
	"github.com/samy-dougui/tftest/cli/internal/rule/filter"
)

type Rule struct {
	Name         string
	Severity     string
	ErrorMessage string
	Filter       Filter
	Condition    string
}

type Filter struct {
	Type string
}

func (r *Rule) Init(block *hcl.Block) hcl.Diagnostics {
	var diags hcl.Diagnostics
	r.Name = block.Labels[0]
	ruleBody, _, _ := block.Body.PartialContent(BlockSchema)

	for _, attribute := range ruleBody.Attributes {
		switch attribute.Name {
		case "condition":
			condition, _ := attribute.Expr.Value(nil)
			r.Condition = condition.AsString()
		case "severity":
			severity, _ := attribute.Expr.Value(nil)
			r.Severity = severity.AsString()
		case "error_message":
			errorMessage, _ := attribute.Expr.Value(nil)
			r.ErrorMessage = errorMessage.AsString()
		default:
			continue
		}
	}

	var my_filter Filter
	for _, my_block := range ruleBody.Blocks {
		switch my_block.Type {
		case "filter":
			filterContent, _, _ := my_block.Body.PartialContent(filter.Schema)
			for _, filterAttribute := range filterContent.Attributes {
				switch filterAttribute.Name {
				case "type":
					filterType, _ := filterAttribute.Expr.Value(nil)
					my_filter.Type = filterType.AsString()
				}
			}
		}
	}
	r.Filter = my_filter
	return diags
}

func (r *Rule) Apply(plan *loader2.Plan) hcl.Diagnostics {
	for _, ressourceChange := range plan.ResourceChanges {
		if ressourceChange.Type == r.Filter.Type {
			fmt.Printf("Resource %v, captured by ryle %v, on filter %v\n", ressourceChange.Address, r.Name, r.Filter.Type)
		}
	}
	return nil
}

var ruleAttributes = []hcl.AttributeSchema{
	{
		Name: "filter",
	},
	{
		Name: "condition",
	},
	{
		Name: "severity",
	}, {
		Name: "error_message",
	},
}

var filterBlock = []hcl.BlockHeaderSchema{
	{Type: "filter"},
}

var BlockSchema = &hcl.BodySchema{
	Attributes: ruleAttributes,
	Blocks:     filterBlock,
}
