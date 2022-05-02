package condition

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/samy-dougui/tftest/cli/internal/loader"
	"github.com/zclconf/go-cty/cty"
)

type Condition struct {
	Attribute string
	Operator  string
	Values    cty.Value
}

func (c *Condition) Init(block *hcl.Block) hcl.Diagnostics {
	var diags hcl.Diagnostics
	conditionContent, _, diag := block.Body.PartialContent(Schema)
	diags = append(diags, diag...)
	for _, conditionAttribute := range conditionContent.Attributes {
		switch conditionAttribute.Name {
		case "attribute":
			conditionAttribute, _ := conditionAttribute.Expr.Value(nil)
			c.Attribute = conditionAttribute.AsString()
		case "values":
			conditionValues, _ := conditionAttribute.Expr.Value(nil)
			c.Values = conditionValues
		case "operator":
			conditionOperator, _ := conditionAttribute.Expr.Value(nil)
			c.Operator = conditionOperator.AsString()
		}
	}
	return diags
}

func (c *Condition) Check(resource *loader.ResourceChange) bool {
	var attribute = resource.GetAttribute(c.Attribute)
	if attribute != nil {
		var operatorCheck = OperatorMap[c.Operator](attribute, c.Values)
		return operatorCheck
	}
	return false
}

var conditionAttributes = []hcl.AttributeSchema{
	{
		Name: "attribute",
	},
	{
		Name: "values",
	},
	{
		Name: "operator",
	},
}

var Schema = &hcl.BodySchema{
	Attributes: conditionAttributes,
}
