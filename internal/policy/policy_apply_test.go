package policy

import (
	"github.com/samy-dougui/ptf/internal/policy/filter"
	"github.com/samy-dougui/ptf/internal/ports"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testPolicyFilterValues = []struct {
	TestPolicy        Policy
	InputResources    []*ports.Resource
	ExpectedResources []*ports.Resource
}{
	{
		TestPolicy: Policy{
			Name: "test_policy_1",
			Filter: filter.Filter{
				Type: "type_1",
			},
		},
		InputResources: []*ports.Resource{
			{
				Address: "address_to_my_resource_1",
				Type:    "type_1",
			},
			{
				Address: "address_to_my_resource_2",
				Type:    "type_2",
			},
		},
		ExpectedResources: []*ports.Resource{
			{
				Address: "address_to_my_resource_1",
				Type:    "type_1",
			},
		},
	},
	{
		TestPolicy: Policy{
			Name: "test_policy_1",
			Filter: filter.Filter{
				Type: "type_2",
			},
		},
		InputResources: []*ports.Resource{
			{
				Address: "address_to_my_resource_1",
				Type:    "type_1",
			},
			{
				Address: "address_to_my_resource_2",
				Type:    "type_2",
			},
		},
		ExpectedResources: []*ports.Resource{
			{
				Address: "address_to_my_resource_2",
				Type:    "type_2",
			},
		},
	},
}

func TestPolicy_filter(t *testing.T) {
	for _, testValue := range testPolicyFilterValues {
		assert.Equal(t, testValue.ExpectedResources, testValue.TestPolicy.filter(&testValue.InputResources))
	}
}
