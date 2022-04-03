package rule

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	loader2 "github.com/samy-dougui/tftest/cli/internal/loader"
	"github.com/samy-dougui/tftest/cli/internal/rule/condition"
	"github.com/samy-dougui/tftest/cli/internal/rule/filter"
)

type Rule struct {
	Name         string
	Severity     string
	ErrorMessage string
	Filter       filter.Filter
	Condition    condition.Condition
}

func (r *Rule) Init(block *hcl.Block) hcl.Diagnostics {
	var diags hcl.Diagnostics
	var ruleFilter filter.Filter
	var ruleCondition condition.Condition

	r.Name = block.Labels[0]
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

func (r *Rule) Apply(plan *loader2.Plan) hcl.Diagnostics {
	var diags hcl.Diagnostics
	for _, resourceChange := range plan.ResourceChanges {
		if isCapturedByFilter := r.Filter.Apply(&resourceChange); isCapturedByFilter {
			isValid := r.Condition.Check(&resourceChange)
			if !isValid {
				ruleDiag := r.FormatError(&resourceChange)
				diags = append(diags, &ruleDiag)
			}
		}
	}
	return diags
}

func (r *Rule) FormatError(resource *loader2.ResourceChange) hcl.Diagnostic {
	var severity hcl.DiagnosticSeverity
	if r.Severity == "warning" {
		severity = hcl.DiagWarning
	} else {
		severity = hcl.DiagError
	}
	return hcl.Diagnostic{
		Severity: severity,
		Summary:  fmt.Sprintf("Resource %v doesn't follow the rule %v", resource.Address, r.Name),
		Detail:   fmt.Sprintf("The resource %v doesn't follow the rule %v. Its attribute %v should be %v.", resource.Address, r.Name, r.Condition.Attribute, r.Condition.Values),
	}
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
	}, {
		Name: "error_message",
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
