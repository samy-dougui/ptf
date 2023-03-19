package ux

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/samy-dougui/ptf/internal/policy"
	"github.com/samy-dougui/ptf/internal/utils"
	"os"
)

func Display(outputs *[]policy.Output, pretty bool, short bool) {
	if pretty {
		prettyDisplay(outputs, short)
	} else {
		rawDisplay(outputs)
	}
}
func rawDisplay(outputs *[]policy.Output) {
	ans, _ := utils.MarshalJson(outputs)
	fmt.Println(string(ans))
}

func prettyDisplay(outputs *[]policy.Output, short bool) {
	if !short {
		displaySummary(outputs)
	}
	displayTable(outputs)
}

func displaySummary(outputs *[]policy.Output) {
	// TODO: add the policy's operator in the output to have custom messages instead of just expected/received
	for _, output := range *outputs {
		fmt.Println(blue(fmt.Sprintf("Rule: %v", output.Name)))
		fmt.Printf("  Status: %v\n", getPrettyPolicyStatus(output.Validated, output.Severity))
		if !output.Validated {
			invalidResourceCount := len(output.InvalidResources)
			fmt.Printf("  Number of Invalid Resource: %v\n", invalidResourceCount)
			for _, invalidResource := range output.InvalidResources {
				fmt.Println(red(fmt.Sprintf("\tInvalid for resource: %v", invalidResource.Address)))
				fmt.Printf("\tCause: %v\n", invalidResource.ErrorMessage)
				invalidAttributeCount := len(invalidResource.InvalidAttributes)
				for index, invalidAttribute := range invalidResource.InvalidAttributes {
					fmt.Printf("\t%v.Expected Attribute: %v\n", index+1, invalidAttribute.ExpectedValue)
					fmt.Printf("\t  Received Attribute: %v\n", invalidAttribute.ReceivedValue)
					if invalidAttributeCount > 1 {
						fmt.Printf("\t  Index: %v\n\n", index)
					}
				}
			}
		}
		fmt.Println("")
	}
}

func displayTable(outputs *[]policy.Output) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Policy Name", "Status"})
	for _, output := range *outputs {
		status := getPrettyPolicyStatus(output.Validated, output.Severity)
		t.AppendRow(table.Row{output.Name, status})
		t.AppendSeparator()
	}
	t.SetStyle(table.StyleLight)
	t.SortBy([]table.SortBy{
		{Number: 2, Mode: table.Asc},
	})
	t.Render()
}

func getPrettyPolicyStatus(validPolicy bool, severity string) string {
	var status string
	if validPolicy {
		status = green("OK")
	} else {
		if severity == policy.ERROR {
			status = red("ERR")
		} else {
			status = yellow("WARN")
		}
	}
	return status
}
