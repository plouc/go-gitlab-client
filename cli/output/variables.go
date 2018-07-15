package output

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/plouc/go-gitlab-client/gitlab"
	"io"
)

func Variables(w io.Writer, format string, collection *gitlab.VariableCollection) {
	if format == "json" {
		collection.RenderJson(w)
	} else if format == "yaml" {
		collection.RenderYaml(w)
	} else {
		fmt.Fprintln(w, "")
		table := tablewriter.NewWriter(w)
		table.SetHeader([]string{
			"Key",
			"Value",
			"Protected",
			"Environment scope",
		})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		for _, variable := range collection.Items {
			table.Append([]string{
				variable.Key,
				variable.Value,
				fmt.Sprintf("%t", variable.Protected),
				variable.EnvironmentScope,
			})
		}
		table.Render()
		fmt.Fprintln(w, "")
	}
}

func Variable(w io.Writer, format string, variable *gitlab.Variable) {
	if format == "json" {
		variable.RenderJson(w)
	} else if format == "yaml" {
		variable.RenderYaml(w)
	} else {
		fmt.Fprintln(w, "")

		fmt.Fprintf(w, "  Key               %s\n", color.YellowString(variable.Key))
		fmt.Fprintf(w, "  Value             %s\n", color.YellowString(variable.Value))
		fmt.Fprintf(w, "  Protected         %s\n", color.YellowString("%t", variable.Protected))
		fmt.Fprintf(w, "  EnvironmentScope  %s\n", color.YellowString(variable.EnvironmentScope))

		fmt.Fprintln(w, "")
	}
}
