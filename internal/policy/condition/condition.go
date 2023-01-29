package condition

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/samy-dougui/ptf/internal/ports"
	"github.com/zclconf/go-cty/cty"
	"log"
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

func (c *Condition) Check(attributes []interface{}) []ports.InvalidAttribute {
	var invalidAttributes []ports.InvalidAttribute
	for _, attribute := range attributes {
		isValid, invalidAttribute, err := OperatorMap[c.Operator](attribute, c.Values)
		if err != nil {
			// TODO: if err, we should return err (or list err) and return it to user
			log.Println(err) // TODO: use logger
		} else {
			if !isValid {
				invalidAttributes = append(invalidAttributes, invalidAttribute)
			}
		}
	}
	return invalidAttributes
}
