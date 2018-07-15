package output

import (
	"fmt"

	"github.com/olekukonko/tablewriter"
	"github.com/plouc/go-gitlab-client/gitlab"
	"io"
)

func Environments(w io.Writer, format string, environments []*gitlab.Environment) {
	if format == "json" {
		Json(w, environments)
	} else if format == "yaml" {
		Yaml(w, environments)
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
		for _, environment := range environments {
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
		Json(w, environment)
	} else if format == "yaml" {
		Yaml(w, environment)
	} else {

	}
}
