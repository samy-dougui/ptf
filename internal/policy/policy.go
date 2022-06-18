package policy

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/samy-dougui/ptf/internal/loader"
	"github.com/samy-dougui/ptf/internal/policy/condition"
	"github.com/samy-dougui/ptf/internal/policy/filter"
)

type Policy struct {
	Name         string
	Severity     string
	ErrorMessage string
	Filter       filter.Filter
	Condition    condition.Condition
	Disabled     bool
	Passed       bool
}

func InitPolicies(blocks hcl.Blocks) ([]Policy, hcl.Diagnostics) {
	policies := make([]Policy, len(blocks))
	var diags hcl.Diagnostics
	for _, block := range blocks {
		switch block.Type {
		case "policy":
			var policy Policy
			policyDiags := policy.Init(block)
			diags = append(diags, policyDiags...)
			policies = append(policies, policy)
		default:
			continue
		}
	}
	return policies, diags
}

func ApplyPolicies(policies *[]Policy, plan *loader.Plan) hcl.Diagnostics {
	var diags hcl.Diagnostics
	for _, policy := range *policies {
		if !policy.IsDisabled() {
			applyDiags := policy.Apply(plan)
			diags = append(diags, applyDiags...)
		}
	}
	return diags
}

func (p *Policy) Init(block *hcl.Block) hcl.Diagnostics {
	var diags hcl.Diagnostics
	var policyFilter filter.Filter
	var policyCondition condition.Condition

	p.Name = block.Labels[0]
	p.initDefaultValues()
	policyBody, _, diagInitPolicy := block.Body.PartialContent(BlockSchema)
	diags = append(diags, diagInitPolicy...)

	for _, attribute := range policyBody.Attributes {
		switch attribute.Name {
		case "severity":
			severity, diagSeverity := attribute.Expr.Value(nil)
			diags = append(diags, diagSeverity...)
			p.Severity = severity.AsString()
		case "error_message":
			errorMessage, diagErrorMessage := attribute.Expr.Value(nil)
			diags = append(diags, diagErrorMessage...)
			p.ErrorMessage = errorMessage.AsString()
		case "disabled":
			disabled, diagDisabled := attribute.Expr.Value(nil)
			diags = append(diags, diagDisabled...)
			p.Disabled = disabled.True()
		default:
			continue
		}
	}
	for _, myBlock := range policyBody.Blocks {
		switch myBlock.Type {
		case "filter":
			diagInitFilter := policyFilter.Init(myBlock)
			diags = append(diags, diagInitFilter...)
		case "condition":
			diagInitCondition := policyCondition.Init(myBlock)
			diags = append(diags, diagInitCondition...)
		}
	}
	p.Filter = policyFilter
	p.Condition = policyCondition
	return diags
}

func (p *Policy) initDefaultValues() {
	p.Disabled = false
	p.Severity = "error"
}

func (p *Policy) Apply(plan *loader.Plan) hcl.Diagnostics {
	var diags hcl.Diagnostics
	for _, resourceChange := range plan.ResourceChanges {
		if isCapturedByFilter := p.Filter.Apply(&resourceChange); isCapturedByFilter {
			isValid, policyDiag := p.Condition.Check(&resourceChange)
			p.Passed = isValid
			if !isValid {
				p.FormatDiag(&resourceChange, &policyDiag)
				diags = append(diags, &policyDiag)
			}
		}
	}
	return diags
}

func (p *Policy) FormatDiag(resource *loader.ResourceChange, diag *hcl.Diagnostic) {
	var severity hcl.DiagnosticSeverity
	if p.Severity == "warning" {
		severity = hcl.DiagWarning
	} else {
		severity = hcl.DiagError
	}
	diag.Severity = severity
	diag.Summary = fmt.Sprintf("Resource %v doesn't follow the policy %v", resource.Address, p.Name)

}

func (p *Policy) IsDisabled() bool {
	return p.Disabled
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

var BlockSchema = &hcl.BodySchema{
	Attributes: policyAttributes,
	Blocks:     policyBlock,
}
