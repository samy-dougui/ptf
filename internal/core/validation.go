package core

import (
	p "github.com/samy-dougui/ptf/internal/policy"
	"github.com/samy-dougui/ptf/internal/ports"
)

func Validate(policies *[]*p.Policy, resources *[]*ports.Resource, configuration *ports.Configuration) []p.Output {
	// TODO: one go-routine per policy with a max concurrency arg
	var output []p.Output
	for _, policy := range *policies {
		policyOutput := policy.Apply(resources, configuration)
		output = append(output, policyOutput)
	}
	return output
}
