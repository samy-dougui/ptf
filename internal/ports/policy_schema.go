package ports

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

var filterAttributes = []hcl.AttributeSchema{
	{
		Name: "type",
	},
}

var filterHclSchema = &hcl.BodySchema{
	Attributes: filterAttributes,
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

var conditionHclSchema = &hcl.BodySchema{
	Attributes: conditionAttributes,
}
