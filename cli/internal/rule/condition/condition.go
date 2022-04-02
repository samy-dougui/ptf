package condition

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/samy-dougui/tftest/cli/internal/loader"
	"strings"
)

type Condition struct {
	Attributes string
	Operator   string
	Values     string
}

var conditionAttributes = []hcl.AttributeSchema{
	{
		Name: "attributes",
	},
	{
		Name: "values",
	},
}

var Schema = &hcl.BodySchema{
	Attributes: conditionAttributes,
}

func (c *Condition) Init(block *hcl.Block) hcl.Diagnostics {
	var diags hcl.Diagnostics
	conditionContent, _, diag := block.Body.PartialContent(Schema)
	diags = append(diags, diag...)
	for _, conditionAttribute := range conditionContent.Attributes {
		switch conditionAttribute.Name {
		case "attributes":
			conditionAttributes, _ := conditionAttribute.Expr.Value(nil)
			c.Attributes = conditionAttributes.AsString()
		case "values":
			conditionValues, _ := conditionAttribute.Expr.Value(nil)
			c.Values = conditionValues.AsString()
		}
	}
	return diags
}

func (c *Condition) Check(resource *loader.ResourceChange) bool {
	// TODO: Replace following block by: ressource.getAttribute(c.Attributes)
	// TODO: it's only for the After, should be configurable

	var _attribute = resource.Change.After
	var nestedAttributes = strings.Split(c.Attributes, ".")
	for _, nestedAttribute := range nestedAttributes[:len(nestedAttributes)-1] {
		_attribute = _attribute[nestedAttribute].(map[string]interface{})
	}
	// TODO: warning or error if the attribute is not defined
	var attribute = _attribute[nestedAttributes[len(nestedAttributes)-1]]
	return attribute == c.Values
}
