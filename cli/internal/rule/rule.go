package rule

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	loader2 "github.com/samy-dougui/tftest/cli/internal/loader"
	"github.com/samy-dougui/tftest/cli/internal/rule/condition"
	"github.com/samy-dougui/tftest/cli/internal/rule/filter"
	"strings"
)

type Rule struct {
	Name         string
	Severity     string
	ErrorMessage string
	Filter       Filter
	Condition    Condition
}

type Filter struct {
	Type string
}

type Condition struct {
	Attributes string
}

func (r *Rule) Init(block *hcl.Block) hcl.Diagnostics {
	var diags hcl.Diagnostics
	r.Name = block.Labels[0]
	ruleBody, _, _ := block.Body.PartialContent(BlockSchema)

	for _, attribute := range ruleBody.Attributes {
		switch attribute.Name {
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
	var my_condition Condition
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
		case "condition":
			conditionContent, _, _ := my_block.Body.PartialContent(condition.Schema)
			for _, conditionAttribute := range conditionContent.Attributes {
				switch conditionAttribute.Name {
				case "attributes":
					conditionAttributes, _ := conditionAttribute.Expr.Value(nil)
					my_condition.Attributes = conditionAttributes.AsString()
				}
			}
		}
	}
	r.Filter = my_filter
	r.Condition = my_condition
	return diags
}

func (r *Rule) Apply(plan *loader2.Plan) hcl.Diagnostics {
	for _, ressourceChange := range plan.ResourceChanges {
		if ressourceChange.Type == r.Filter.Type {
			var _attribute = ressourceChange.Change.After
			var nestedAttributes = strings.Split(r.Condition.Attributes, ".")
			for _, nestedAttribute := range nestedAttributes[:len(nestedAttributes)-1] {
				_attribute = _attribute[nestedAttribute].(map[string]interface{})
			}
			var attribute = _attribute[nestedAttributes[len(nestedAttributes)-1]]
			fmt.Printf("Resource %v, captured by ryle %v, on filter %v\n", ressourceChange.Address, r.Name, r.Filter.Type)
			fmt.Printf("Attribute %v = %v\n", r.Condition.Attributes, attribute)
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

var ruleBlock = []hcl.BlockHeaderSchema{
	{
		Type: "filter",
	},
	{
		Type: "condition",
	},
}

var BlockSchema = &hcl.BodySchema{
	Attributes: ruleAttributes,
	Blocks:     ruleBlock,
}
