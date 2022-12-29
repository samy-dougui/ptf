package policy

import (
	"github.com/samy-dougui/ptf/internal/ports"
	"github.com/samy-dougui/ptf/internal/utils"
)

func (p *Policy) Apply(resources *[]*ports.Resource, configuration *ports.Configuration) Output {
	var invalidResources []ports.InvalidResource
	filteredResources := p.filter(resources)
	for _, resource := range filteredResources {
		attributes, err := utils.GetResourceAttribute(resource.Values, p.Condition.Attribute)
		if err != nil {
			// TODO: check if error is missing attributes => create real error
			// TODO: make error message more explicit
			invalidResources = append(invalidResources, ports.InvalidResource{
				Address:           resource.Address,
				AttributeName:     p.Condition.Attribute,
				ReceivedAttribute: attributes,
				ErrorMessage:      "error retrieving attribute",
			})
		} else {
			if validResource := p.Condition.Check(attributes); !validResource {
				invalidResources = append(invalidResources, ports.InvalidResource{
					Address:           resource.Address,
					AttributeName:     p.Condition.Attribute,
					ReceivedAttribute: attributes,
					ErrorMessage:      p.ErrorMessage,
				})
			}

		}
	}

	validPolicy := !(len(invalidResources) >= 1)

	var severity string
	if !validPolicy {
		severity = p.Severity
	}
	return Output{
		Name:                p.Name,
		Validated:           validPolicy,
		InvalidResourceList: invalidResources,
		Severity:            severity,
	}
}

func (p *Policy) filter(resources *[]*ports.Resource) []*ports.Resource {
	var selectedResources []*ports.Resource
	for _, resource := range *resources {
		if p.Filter.Apply(resource) {
			selectedResources = append(selectedResources, resource)
		}
	}
	return selectedResources
}
