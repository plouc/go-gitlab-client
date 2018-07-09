package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/plouc/go-gitlab-client/gitlab"
)

func varsOutput(variables []*gitlab.Variable) {
	if outputFormat == "json" {
		jsonOutput(variables)
	} else {
		fmt.Fprintln(output, "")
		table := tablewriter.NewWriter(output)
		table.SetHeader([]string{
			"Key",
			"Value",
			"Protected",
			"Environment scope",
		})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		for _, variable := range variables {
			table.Append([]string{
				variable.Key,
				variable.Value,
				fmt.Sprintf("%t", variable.Protected),
				variable.EnvironmentScope,
			})
		}
		table.Render()
		fmt.Fprintln(output, "")
	}
}

func varOutput(variable *gitlab.Variable) {
	if outputFormat == "json" {
		jsonOutput(variable)
	} else {
		fmt.Fprintln(output, "")

		fmt.Fprintf(output, "  Key               %s\n", color.YellowString(variable.Key))
		fmt.Fprintf(output, "  Value             %s\n", color.YellowString(variable.Value))
		fmt.Fprintf(output, "  Protected         %s\n", color.YellowString("%t", variable.Protected))
		fmt.Fprintf(output, "  EnvironmentScope  %s\n", color.YellowString(variable.EnvironmentScope))

		fmt.Fprintln(output, "")
	}
}
