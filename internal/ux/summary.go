package ux

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/samy-dougui/ptf/internal/logging"
	"github.com/samy-dougui/ptf/internal/policy"
)

func WriteSummary(diags *hcl.Diagnostics, policies *[]policy.Policy) {
	logger := logging.GetLogger()
	numberOfPolicies := len(*policies)
	numberOfWarnings := getNumberOfWarnings(diags)
	numberOfErrors := getNumberOfErrors(diags)
	numberOfPass := getNumberOfPass(policies)
	numberOfDisabled := numberOfPolicies - (numberOfErrors + numberOfPass + numberOfWarnings)
	logger.Infof("Number of Policies: %v", numberOfPolicies)
	logger.Infof("Number of Pass: %v", numberOfPass)
	logger.Infof("Number of Warnings: %v", numberOfWarnings)
	logger.Infof("Number of Errors: %v", numberOfErrors)
	logger.Infof("Number of Disabled: %v", numberOfDisabled)
}

func getNumberOfWarnings(diags *hcl.Diagnostics) int {
	var numberOfWarnings int
	for _, diag := range *diags {
		if diag.Severity == hcl.DiagWarning {
			numberOfWarnings += 1
		}
	}
	return numberOfWarnings
}

func getNumberOfErrors(diags *hcl.Diagnostics) int {
	var numberOfErrors int
	for _, diag := range *diags {
		if diag.Severity == hcl.DiagError {
			numberOfErrors += 1
		}
	}
	return numberOfErrors
}

func getNumberOfPass(policies *[]policy.Policy) int {
	var numberOfPass int
	for _, policy := range *policies {
		if policy.Passed {
			numberOfPass += 1
		}
	}
	return numberOfPass
}
