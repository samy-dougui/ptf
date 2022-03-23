package filter

import "github.com/hashicorp/hcl/v2"

var filterAttributes = []hcl.AttributeSchema{
	{
		Name: "type",
	},
}
var Schema = &hcl.BodySchema{
	Attributes: filterAttributes,
}
