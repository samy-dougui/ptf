package rule

import "github.com/hashicorp/hcl/v2"

type Resource struct {
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
