package ux

import (
	"fmt"
	"github.com/samy-dougui/ptf/internal/ports"
)

func DisplayOutputPolicies(outputs *[]ports.PolicyOutput) {
	for _, output := range *outputs {
		if !output.Validated {
			formattedOutput := formatOutput(output)
			fmt.Println(formattedOutput)
		}
	}
}

func formatOutput(output ports.PolicyOutput) string {
	header := fmt.Sprintf("The policy %s has not been validated. The following resources are not respecting it: ", output.Name)
	var body string
	for _, invalidResource := range output.InvalidResourceList {
		body = body + fmt.Sprintf("\t- The resource %s was expecting the attribute '%s' to be '%s' but it's equal to '%s'\n", invalidResource.Address, invalidResource.AttributeName, invalidResource.ExpectedAttribute, invalidResource.ReceivedAttribute)
	}
	return fmt.Sprintf("%s\n%s", header, body)
}
