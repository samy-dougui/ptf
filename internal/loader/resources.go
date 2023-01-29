package loader

import (
	"github.com/samy-dougui/ptf/internal/ports"
	"strings"
)

func LoadLocalResources(path string) ([]*ports.Resource, ports.Configuration) {
	plan := getPlan(path)
	resources, configuration := adaptPlan(plan)
	return resources, configuration
}

func getPlan(path string) *ports.Plan {
	var fileLoader Loader
	fileLoader.Init()
	plan, _ := fileLoader.LoadPlan(path)
	return plan
}

func adaptPlan(plan *ports.Plan) ([]*ports.Resource, ports.Configuration) {
	var resources []*ports.Resource
	for _, resourceChange := range plan.ResourceChanges {
		resource := ports.Resource{
			Address:  resourceChange.Address,
			Type:     resourceChange.Type,
			Provider: resourceChange.ProviderName,
			Action:   resourceChange.ActionReason,
			Values:   resourceChange.Change.After, // TODO: merge all the after attributes
		}
		resources = append(resources, &resource)
	}
	var variables []ports.Variable
	for variableName, variableValue := range plan.Variables {
		variable := ports.Variable{
			Name:  variableName,
			Value: variableValue,
		}
		variables = append(variables, variable)
	}
	var providers []ports.Provider
	for providerName, providerAttribute := range plan.Configuration.Providers {
		if !strings.HasPrefix(providerName, "module") {
			providers = append(providers, providerAttribute)
		}
	}
	configuration := ports.Configuration{
		Variables:        variables,
		TerraformVersion: plan.TerraformVersion,
		Providers:        providers,
	}
	return resources, configuration
}
