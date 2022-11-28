package core

import "github.com/samy-dougui/ptf/internal/ports"

func apply(policy *ports.Policy, resources *[]*ports.Resource, configuration *ports.Configuration) ports.PolicyOutput {
	var invalidResources []ports.InvalidResource
	for _, resource := range *resources {
		if policy.Condition.Values.AsString() != resource.Values[policy.Condition.Attribute] {
			invalidResources = append(invalidResources, ports.InvalidResource{
				Address:           resource.Address,
				AttributeName:     policy.Condition.Attribute,
				ExpectedAttribute: policy.Condition.Values.AsString(),
				ReceivedAttribute: resource.Values[policy.Condition.Attribute],
				ErrorMessage:      policy.ErrorMessage,
			})
		}
	}

	var validPolicy bool
	if len(invalidResources) >= 1 {
		validPolicy = false
	} else {
		validPolicy = true
	}
	var severity string
	if !validPolicy {
		severity = policy.Severity
	}
	return ports.PolicyOutput{
		Name:                policy.Name,
		Validated:           validPolicy,
		InvalidResourceList: invalidResources,
		Severity:            severity,
	}
}
