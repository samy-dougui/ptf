package policy

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/samy-dougui/ptf/internal/policy/condition"
	"github.com/samy-dougui/ptf/internal/policy/filter"
	"github.com/samy-dougui/ptf/internal/ports"
)

const (
	WARNING = "warning"
	ERROR   = "error"
)

type Policy struct {
	Name         string
	Target       string
	Severity     string
	ErrorMessage string
	Disabled     bool
	Filter       filter.Filter
	Condition    condition.Condition
}

type Output struct {
	Name             string                  `json:"name"`
	Validated        bool                    `json:"is_valid"`
	Severity         string                  `json:"severity,omitempty"`
	InvalidResources []ports.InvalidResource `json:"invalid_resources,omitempty"`
}

func (p *Policy) Init(policyBlock *hcl.Block) {
	p.Name = policyBlock.Labels[0]
	p.defaultInit()
	policyBody, _, _ := policyBlock.Body.PartialContent(policyHclSchema) // TODO: add error handling with diags

	for _, attribute := range policyBody.Attributes {
		switch attribute.Name {
		case "severity":
			severity, _ := attribute.Expr.Value(nil)
			p.Severity = severity.AsString()
		case "error_message":
			errorMessage, _ := attribute.Expr.Value(nil)
			p.ErrorMessage = errorMessage.AsString()
		case "disabled":
			disabled, _ := attribute.Expr.Value(nil)
			p.Disabled = disabled.True()
		default:
			continue
		}
	}
	for _, block := range policyBody.Blocks {
		switch block.Type {
		case "filter":
			p.Filter.Init(block)
		case "condition":
			p.Condition.Init(block)
		default:
			fmt.Printf("Unknown block inside policy: %v\n", block.Type)
			continue
		}
	}
}

func (p *Policy) defaultInit() {
	p.Disabled = false
	p.Severity = ERROR
}
