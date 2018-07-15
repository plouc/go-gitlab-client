package output

import (
	"fmt"

	"github.com/olekukonko/tablewriter"
	"github.com/plouc/go-gitlab-client/gitlab"
	"io"
)

func SshKeys(w io.Writer, format string, collection *gitlab.SshKeyCollection) {
	if format == "json" {
		collection.RenderJson(w)
	} else if format == "yaml" {
		collection.RenderYaml(w)
	} else {
		fmt.Fprintln(w, "")
		table := tablewriter.NewWriter(w)
		table.SetHeader([]string{
			"Id",
			"Title",
			"Key",
			"Created at",
		})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		for _, key := range collection.Items {
			table.Append([]string{
				fmt.Sprintf("%d", key.Id),
				key.Title,
				key.Key[:16] + "…",
				key.CreatedAtRaw,
			})
		}
		table.Render()
		fmt.Fprintln(w, "")
	}
}
