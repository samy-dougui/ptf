package resource

import "github.com/hashicorp/hcl/v2"

type Resource struct {
}

var commonResourceAttributes = []hcl.AttributeSchema{
	{
		Name: "count",
	},
	{
		Name: "for_each",
	},
	{
		Name: "provider",
	},
	{
		Name: "depends_on",
	},
}
var BlockSchema = &hcl.BodySchema{
	Attributes: commonResourceAttributes,
	Blocks: []hcl.BlockHeaderSchema{
		{Type: "locals"}, // reserved for future use
		{Type: "lifecycle"},
		{Type: "connection"},
		{Type: "provisioner", LabelNames: []string{"type"}},
		{Type: "_"}, // meta-argument escaping block
	},
}
