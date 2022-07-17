package utils

import (
	"github.com/hashicorp/hcl/v2"
)

func ConcatDiagsDetail(diags *hcl.Diagnostics) string {
	var diagsDetail string
	for _, diag := range *diags {
		diagsDetail += "- " + diag.Detail + "\n"
	}
	return diagsDetail
}
