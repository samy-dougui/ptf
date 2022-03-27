package condition

import "github.com/hashicorp/hcl/v2"

type Condition struct {
	Attribute string
	Operator  string
	Values    interface{}
}

var conditionAttributes = []hcl.AttributeSchema{
	{
		Name: "attributes",
	},
}

var Schema = &hcl.BodySchema{
	Attributes: conditionAttributes,
}
