package condition

import "github.com/hashicorp/hcl/v2"

var attributes = []hcl.AttributeSchema{
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

var hclSchema = &hcl.BodySchema{
	Attributes: attributes,
}
