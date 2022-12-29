package filter

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/samy-dougui/ptf/internal/ports"
)

type Filter struct {
	Type string
}

func (f *Filter) Init(block *hcl.Block) {
	var diags hcl.Diagnostics
	content, _, _ := block.Body.PartialContent(hclSchema)
	for _, attribute := range content.Attributes {
		switch attribute.Name {
		case "type":
			filterType, diag := attribute.Expr.Value(nil)
			f.Type = filterType.AsString()
			diags = append(diags, diag...)
		}
	}
}

func (f *Filter) Apply(resource *ports.Resource) bool {
	// NOTE: it only works with resource type for now
	return resource.Type == f.Type
}
