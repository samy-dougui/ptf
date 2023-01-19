package policy

import (
	"errors"
	"fmt"
	"github.com/samy-dougui/ptf/internal/ports"
	"github.com/samy-dougui/ptf/internal/utils"
)

func (p *Policy) Apply(resources *[]*ports.Resource, configuration *ports.Configuration) Output {
	var invalidResources []ports.InvalidResource
	filteredResources := p.filter(resources)
	for _, resource := range filteredResources {
		attributes, err := utils.GetResourceAttribute(resource.Values, p.Condition.Attribute)
		if err != nil {
			// TODO: make error message more explicit
			if errors.As(err, &utils.MissingAttributeError{}) {
				invalidResources = append(invalidResources, ports.InvalidResource{
					Address:       resource.Address,
					AttributeName: p.Condition.Attribute,
					ErrorMessage:  err.Error(),
				})
			} else {
				invalidResources = append(invalidResources, ports.InvalidResource{
					Address:       resource.Address,
					AttributeName: p.Condition.Attribute,
					ErrorMessage:  fmt.Sprintf("error retrieving attribute %v", p.Condition.Attribute),
				})
			}
		} else {
			invalidAttributes := p.Condition.Check(attributes) // TODO: return list of invalid attributes, if list is nil => valid resource
			if len(invalidAttributes) >= 1 {
				invalidResources = append(invalidResources, ports.InvalidResource{
					Address:           resource.Address,
					AttributeName:     p.Condition.Attribute,
					ErrorMessage:      "Invalid attribute",
					InvalidAttributes: invalidAttributes,
				})
			}
			//if !invalidAttributes {
			//	invalidResources = append(invalidResources, ports.InvalidResource{
			//		Address:           resource.Address,
			//		AttributeName:     p.Condition.Attribute,
			//		ReceivedAttribute: attributes,
			//		ErrorMessage:      p.ErrorMessage,
			//	})
			//}
		}
	}

	validPolicy := !(len(invalidResources) >= 1)

	var severity string
	if !validPolicy {
		severity = p.Severity
	}
	return Output{
		Name:             p.Name,
		Validated:        validPolicy,
		InvalidResources: invalidResources,
		Severity:         severity,
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
