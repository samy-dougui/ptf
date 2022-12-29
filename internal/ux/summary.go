package ux

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/samy-dougui/ptf/internal/logging"
	"github.com/samy-dougui/ptf/internal/old/policy"
)

func WriteSummary(diags *hcl.Diagnostics, policies []*policy.Policy) {
	logger := logging.GetLogger()
	numberOfPolicies := len(policies)
	numberOfWarnings := getNumberOfWarnings(diags)
	numberOfErrors := getNumberOfErrors(diags)
	numberOfPass := getNumberOfPass(policies)
	logger.Infof("Number of Policies: %v", numberOfPolicies)
	logger.Infof("Number of Pass: %v", numberOfPass)
	logger.Infof("Number of Warnings: %v", numberOfWarnings)
	logger.Infof("Number of Errors: %v", numberOfErrors)
}

func getNumberOfWarnings(diags *hcl.Diagnostics) int {
	var numberOfWarnings int
	for _, diag := range *diags {
		if diag.Severity == hcl.DiagWarning {
			numberOfWarnings++
		}
	}
	return numberOfWarnings
}

func getNumberOfErrors(diags *hcl.Diagnostics) int {
	var numberOfErrors int
	for _, diag := range *diags {
		if diag.Severity == hcl.DiagError {
			numberOfErrors++
		}
	}
	return numberOfErrors
}

func getNumberOfPass(policies []*policy.Policy) int {
	numberOfPass := 0
	for _, policy := range policies {
		if policy.Passed {
			numberOfPass++
		}
	}
	return numberOfPass
}
