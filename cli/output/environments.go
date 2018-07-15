package output

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/plouc/go-gitlab-client/gitlab"
	"io"
)

func Environments(w io.Writer, format string, collection *gitlab.EnvironmentCollection) {
	if format == "json" {
		collection.RenderJson(w)
	} else if format == "yaml" {
		collection.RenderYaml(w)
	} else {
		fmt.Fprintln(w, "")
		table := tablewriter.NewWriter(w)
		table.SetHeader([]string{
			"Id",
			"Name",
			"Slug",
			"External URL",
		})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		for _, environment := range collection.Items {
			table.Append([]string{
				fmt.Sprintf("%d", environment.Id),
				environment.Name,
				environment.Slug,
				environment.ExternalUrl,
			})
		}
		table.Render()
		fmt.Fprintln(w, "")
	}
}

func Environment(w io.Writer, format string, environment *gitlab.Environment) {
	if format == "json" {
		environment.RenderJson(w)
	} else if format == "yaml" {
		environment.RenderYaml(w)
	} else {
		fmt.Fprintln(w, "")
		fmt.Fprintf(w, "  Id            %s\n", color.YellowString("%d", environment.Id))
		fmt.Fprintf(w, "  Name          %s\n", color.YellowString(environment.Name))
		fmt.Fprintf(w, "  Slug          %s\n", color.YellowString(environment.Slug))
		fmt.Fprintf(w, "  External URL  %s\n", color.YellowString(environment.ExternalUrl))
		fmt.Fprintln(w, "")
	}
}
