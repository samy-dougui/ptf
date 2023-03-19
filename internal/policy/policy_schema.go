package policy

import "github.com/hashicorp/hcl/v2"

var PolicyFileSchema = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type:       "policy",
			LabelNames: []string{"name"},
		},
	},
}

var policyAttributes = []hcl.AttributeSchema{
	{
		Name: "filter",
	},
	{
		Name: "condition",
	},
	{
		Name: "severity",
	},
	{
		Name: "error_message",
	},
	{
		Name: "disabled",
	},
	{
		Name: "name",
	},
}

var policyBlock = []hcl.BlockHeaderSchema{
	{
		Type: "filter",
	},
	{
		Type: "condition",
	},
}

var policyHclSchema = &hcl.BodySchema{
	Attributes: policyAttributes,
	Blocks:     policyBlock,
}
