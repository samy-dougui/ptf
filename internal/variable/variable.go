package variable

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
)

type Variable struct {
	Name        string
	Description string
	Default     cty.Value
	Value       cty.Value
}

var BlockSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "description",
		},
		{
			Name: "default",
		},
	},
}
