package ports

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
)

type Policy struct {
	Name         string
	Type         string
	Severity     string
	ErrorMessage string
	Disabled     bool
	Filter       Filter
	Condition    Condition
}

type Filter struct {
	Type string
}

type Condition struct {
	Attribute string
	Operator  string
	Values    cty.Value
}

type PolicyOutput struct {
	Name                string            `json:"name"`
	Validated           bool              `json:"is_valid"`
	Severity            string            `json:"severity,omitempty"`
	InvalidResourceList []InvalidResource `json:"invalid_resource,omitempty"`
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
			p.Filter.init(block)
		case "condition":
			p.Condition.init(block)
		default:
			fmt.Printf("Unknown block inside policy: %v\n", block.Type)
			continue
		}
	}
}

func (p *Policy) defaultInit() {
	p.Disabled = false
	p.Severity = "error"
}

func (f *Filter) init(block *hcl.Block) {
	var diags hcl.Diagnostics
	content, _, _ := block.Body.PartialContent(filterHclSchema)
	for _, attribute := range content.Attributes {
		switch attribute.Name {
		case "type":
			filterType, diag := attribute.Expr.Value(nil)
			f.Type = filterType.AsString()
			diags = append(diags, diag...)
		}
	}
}

func (c *Condition) init(block *hcl.Block) {
	content, _, _ := block.Body.PartialContent(conditionHclSchema)
	for _, attribute := range content.Attributes {
		switch attribute.Name {
		case "attribute":
			conditionAttribute, _ := attribute.Expr.Value(nil)
			c.Attribute = conditionAttribute.AsString()
		case "values":
			conditionValues, _ := attribute.Expr.Value(nil)
			c.Values = conditionValues
		case "operator":
			conditionOperator, _ := attribute.Expr.Value(nil)
			c.Operator = conditionOperator.AsString()
		}
	}
}
