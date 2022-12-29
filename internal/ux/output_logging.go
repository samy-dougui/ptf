package ux

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/samy-dougui/ptf/internal/policy"
	"os"
)

func PrettyDisplay(outputs *[]policy.Output) {
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
		if severity == "error" {
			status = red("ERR")
		} else {
			status = yellow("WARN")
		}
	}
	return status
}
