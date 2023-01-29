package filter

import "github.com/hashicorp/hcl/v2"

var attributes = []hcl.AttributeSchema{
	{
		Name: "type",
	},
}

var hclSchema = &hcl.BodySchema{
	Attributes: attributes,
}
