package filter

import (
	"github.com/hashicorp/hcl/v2"
	loader2 "github.com/samy-dougui/ptf/internal/loader"
)

type Filter struct {
	Type string
}

func (f *Filter) Init(block *hcl.Block) hcl.Diagnostics {
	var diags hcl.Diagnostics
	filterContent, _, diag := block.Body.PartialContent(Schema)
	diags = append(diags, diag...)
	for _, filterAttribute := range filterContent.Attributes {
		switch filterAttribute.Name {
		case "type":
			filterType, diag := filterAttribute.Expr.Value(nil)
			f.Type = filterType.AsString()
			diags = append(diags, diag...)
		}
	}
	return diags
}

// The Filtering logic should be included in this function
func (f *Filter) Apply(resource *loader2.ResourceChange) bool {
	return f.Type == resource.Type
}

var filterAttributes = []hcl.AttributeSchema{
	{
		Name: "type",
	},
}
var Schema = &hcl.BodySchema{
	Attributes: filterAttributes,
}