package main

import (
	"github.com/samy-dougui/ptf/internal/core"
	"github.com/samy-dougui/ptf/internal/loader"
	"github.com/samy-dougui/ptf/internal/ports"
	"github.com/samy-dougui/ptf/internal/ux"
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
	resources := loadResources()
	configuration := ports.Configuration{
		Variables:        nil,
		TerraformVersion: "1.12.0",
		ProvidersConfiguration: []ports.ProviderConfiguration{
			{
				Name:              "azure",
				VersionConstraint: "3.15.0",
			},
		},
	}
	validationOutput := core.Validate(&policies, &resources, &configuration)
	ux.DisplayOutputPolicies(&validationOutput)
}

func loadPolicies(path string) []*ports.Policy {
	// TODO: move to another package and split
	var policies []*ports.Policy
	var loader loader.Loader
	loader.Init()

	policiesBody, _ := loader.LoadHCLFile(path)
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

func loadResources() []*ports.Resource {
	// TODO: 1. LoadPlan and return interface "json"
	// TODO: 2. Transform plan to get list of resources and configuration


	var resources []*ports.Resource
	// load plan
	// for resource in after plan
	// adapt to new struct type

	resource1 := ports.Resource{
		Address: "my_resource_1",
		Type:    "managed",
		Values:  map[string]string{"foo": "bar3"},
		Action:  "delete",
	}
	resource2 := ports.Resource{
		Address: "my_resource_2",
		Type:    "managed",
		Values:  map[string]string{"foo": "bar2"},
		Action:  "delete",
	}
	return resources
}

func GetPlan(path string) *ports.Plan {}

func SplitPlan(plan *ports.Plan) ([]*ports.Resource, ports.Configuration) {}
