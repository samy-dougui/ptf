package utils

import (
	"github.com/hashicorp/hcl/v2"
	"testing"
)

func TestConcatdiagsdetail(t *testing.T) {
	diags := &hcl.Diagnostics{
		&hcl.Diagnostic{Detail: "item 1"},
		&hcl.Diagnostic{Detail: "item 2"},
		&hcl.Diagnostic{Detail: "item 3"},
	}
	got := ConcatDiagsDetail(diags)
	expected := "- item 1\n- item 2\n- item 3\n"
	if got != expected {
		t.Errorf("ConcatDiagsDetail(diags) = %v; want %v", got, expected)
	}
}
