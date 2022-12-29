package condition

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
)

type Condition struct {
	Attribute string
	Operator  string
	Values    cty.Value
}

func (c *Condition) Init(block *hcl.Block) {
	content, _, _ := block.Body.PartialContent(hclSchema)
	for _, attribute := range content.Attributes {
		switch attribute.Name {
		case "attribute":
			conditionAttribute, _ := attribute.Expr.Value(nil)
			c.Attribute = conditionAttribute.AsString()
		case "values":
			conditionValues, _ := attribute.Expr.Value(nil)
			c.Values = conditionValues
		case "operator":
			conditionOperator, _ := attribute.Expr.Value(nil)
			c.Operator = conditionOperator.AsString()
		}
	}
}

func (c *Condition) Check(attributes []interface{}) bool {
	var validResource = true
	for _, attribute := range attributes {
		validAttribute, _ := OperatorMap[c.Operator](attribute, c.Values)
		validResource = validResource && validAttribute
	}
	return validResource
}
