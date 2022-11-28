package core

import (
	"github.com/samy-dougui/ptf/internal/ports"
)

func Validate(policies *[]*ports.Policy, resources *[]*ports.Resource, configuration *ports.Configuration) []ports.PolicyOutput {
	var output []ports.PolicyOutput
	for _, policy := range *policies {
		policyOutput := apply(policy, resources, configuration)
		output = append(output, policyOutput)
	}
	return output
}
