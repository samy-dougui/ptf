package ux

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/samy-dougui/ptf/internal/old/policy"
	"testing"
)

func TestGetnumberofwarnings_1(t *testing.T) {
	diags := hcl.Diagnostics{
		&hcl.Diagnostic{
			Severity: hcl.DiagWarning,
		},
		&hcl.Diagnostic{
			Severity: hcl.DiagError,
		},
	}
	got := getNumberOfWarnings(&diags)
	expected := 1
	if got != expected {
		t.Errorf("getNumberOfWarnings(diags) = %d; want %d", got, expected)
	}
}

func TestGetnumberofwarnings_2(t *testing.T) {
	diags := hcl.Diagnostics{
		&hcl.Diagnostic{
			Severity: hcl.DiagError,
		},
		&hcl.Diagnostic{
			Severity: hcl.DiagError,
		},
	}
	got := getNumberOfWarnings(&diags)
	expected := 0
	if got != expected {
		t.Errorf("getNumberOfWarnings(diags) = %d; want %d", got, expected)
	}
}

func TestGetnumberoferrors_1(t *testing.T) {
	diags := hcl.Diagnostics{
		&hcl.Diagnostic{
			Severity: hcl.DiagWarning,
		},
		&hcl.Diagnostic{
			Severity: hcl.DiagError,
		},
	}
	got := getNumberOfErrors(&diags)
	expected := 1
	if got != expected {
		t.Errorf("getNumberOfErrors(diags) = %d; want %d", got, expected)
	}
}

func TestGetnumberoferrors_2(t *testing.T) {
	diags := hcl.Diagnostics{
		&hcl.Diagnostic{
			Severity: hcl.DiagWarning,
		},
		&hcl.Diagnostic{
			Severity: hcl.DiagWarning,
		},
	}
	got := getNumberOfErrors(&diags)
	expected := 0
	if got != expected {
		t.Errorf("getNumberOfErrors(diags) = %d; want %d", got, expected)
	}
}

func TestGetnumberofpass(t *testing.T) {
	policies := []*policy.Policy{
		{
			Name:     "policy_1",
			Severity: "warn",
			Disabled: false,
			Passed:   false,
		},
		{
			Name:     "policy_1",
			Severity: "warn",
			Disabled: false,
			Passed:   true,
		},
	}
	got := getNumberOfPass(policies)
	expected := 1
	if got != expected {
		t.Errorf("getNumberOfPass(policies) = %d; want %d", got, expected)
	}
}
