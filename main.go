package main

import (
	"github.com/samy-dougui/ptf/internal/core"
	"github.com/samy-dougui/ptf/internal/loader"
	"github.com/samy-dougui/ptf/internal/ports"
	"github.com/samy-dougui/ptf/internal/ux"
	"strings"
)

//func main() {
//	setUp()
//	cmd.Execute()
//}
//
//func setUp() {
//	logging.SetUpLogger()
//	logging.SetUpDiagnosticLogger()
//}

func main() {
	policies := loadPolicies("./data/new_policies.hcl")
	resources, configuration := loadResources("./data/tfplan.json")
	validationOutput := core.Validate(&policies, &resources, &configuration)
	ux.DisplayOutputPolicies(&validationOutput)
	return
}

func loadPolicies(path string) []*ports.Policy {
	// TODO: move to another package and split
	var policies []*ports.Policy
	var fileLoader loader.Loader
	fileLoader.Init()

	policiesBody, _ := fileLoader.LoadHCLFile(path)
	policiesBlock, _ := policiesBody.Content(ports.PolicyFileSchema)
	for _, policyBlock := range policiesBlock.Blocks {
		switch policyBlock.Type {
		case "policy":
			var policy ports.Policy
			policy.Init(policyBlock)
			policies = append(policies, &policy)
		default:
			continue
		}
	}
	return policies
}

func loadResources(path string) ([]*ports.Resource, ports.Configuration) {
	plan := GetPlan(path)
	resources, configuration := AdaptPlan(plan)
	return resources, configuration
}

func GetPlan(path string) *ports.Plan {
	var fileLoader loader.Loader
	fileLoader.Init()
	plan, _ := fileLoader.LoadPlan(path)
	return plan
}

func AdaptPlan(plan *ports.Plan) ([]*ports.Resource, ports.Configuration) {
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
