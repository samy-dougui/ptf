package rule

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/samy-dougui/ptf/internal/loader"
	"github.com/samy-dougui/ptf/internal/rule/condition"
	"github.com/samy-dougui/ptf/internal/rule/filter"
)

type Rule struct {
	Name         string
	Severity     string
	ErrorMessage string
	Filter       filter.Filter
	Condition    condition.Condition
	Disabled     bool
}

// TODO: add more logging in init rule

func (r *Rule) Init(block *hcl.Block) hcl.Diagnostics {
	var diags hcl.Diagnostics
	var ruleFilter filter.Filter
	var ruleCondition condition.Condition

	r.Name = block.Labels[0]
	r.initDefaultValues()
	ruleBody, _, diagInitRule := block.Body.PartialContent(BlockSchema)
	diags = append(diags, diagInitRule...)

	for _, attribute := range ruleBody.Attributes {
		switch attribute.Name {
		case "severity":
			severity, diagSeverity := attribute.Expr.Value(nil)
			diags = append(diags, diagSeverity...)
			r.Severity = severity.AsString()
		case "error_message":
			errorMessage, diagErrorMessage := attribute.Expr.Value(nil)
			diags = append(diags, diagErrorMessage...)
			r.ErrorMessage = errorMessage.AsString()
		case "disabled":
			disabled, diagDisabled := attribute.Expr.Value(nil)
			diags = append(diags, diagDisabled...)
			r.Disabled = disabled.True()
		default:
			continue
		}
	}
	for _, myBlock := range ruleBody.Blocks {
		switch myBlock.Type {
		case "filter":
			diagInitFilter := ruleFilter.Init(myBlock)
			diags = append(diags, diagInitFilter...)
		case "condition":
			diagInitCondition := ruleCondition.Init(myBlock)
			diags = append(diags, diagInitCondition...)
		}
	}
	r.Filter = ruleFilter
	r.Condition = ruleCondition
	return diags
}

func (r *Rule) initDefaultValues() {
	r.Disabled = false
	r.Severity = "error"
}

func (r *Rule) Apply(plan *loader.Plan) hcl.Diagnostics {
	var diags hcl.Diagnostics
	for _, resourceChange := range plan.ResourceChanges {
		if isCapturedByFilter := r.Filter.Apply(&resourceChange); isCapturedByFilter {
			isValid, ruleDiag := r.Condition.Check(&resourceChange)
			if !isValid {
				r.FormatDiag(&resourceChange, &ruleDiag)
				diags = append(diags, &ruleDiag)
			}
		}
	}
	return diags
}

func (r *Rule) FormatDiag(resource *loader.ResourceChange, diag *hcl.Diagnostic) {
	var severity hcl.DiagnosticSeverity
	if r.Severity == "warning" {
		severity = hcl.DiagWarning
	} else {
		severity = hcl.DiagError
	}
	diag.Severity = severity
	diag.Summary = fmt.Sprintf("Resource %v doesn't follow the rule %v", resource.Address, r.Name)

}

func (r *Rule) IsDisabled() bool {
	return r.Disabled
}

var ruleAttributes = []hcl.AttributeSchema{
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

var ruleBlock = []hcl.BlockHeaderSchema{
	{
		Type: "filter",
	},
	{
		Type: "condition",
	},
}

var BlockSchema = &hcl.BodySchema{
	Attributes: ruleAttributes,
	Blocks:     ruleBlock,
}
