package config

import "github.com/hashicorp/hcl/v2"

var ConfigFileSchema = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type:       "policy",
			LabelNames: []string{"name"},
		},
	},
}
