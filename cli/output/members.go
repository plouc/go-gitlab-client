package output

import (
	"fmt"
	"io"

	"github.com/olekukonko/tablewriter"
	"github.com/plouc/go-gitlab-client/gitlab"
)

func Members(w io.Writer, format string, collection *gitlab.MemberCollection) {
	if format == "json" {
		collection.RenderJson(w)
	} else if format == "yaml" {
		collection.RenderYaml(w)
	} else {
		fmt.Fprintln(w, "")
		table := tablewriter.NewWriter(w)
		table.SetHeader([]string{
			"Id",
			"Username",
			"Name",
			"State",
			"Expires at",
			"Access level",
			"Web url",
		})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		for _, member := range collection.Items {
			table.Append([]string{
				fmt.Sprintf("%d", member.Id),
				member.Username,
				member.Name,
				member.State,
				member.ExpiresAt,
				fmt.Sprintf("%d", member.AccessLevel),
				member.WebUrl,
			})
		}
		table.Render()
		fmt.Fprintln(w, "")
	}
}

func Member(w io.Writer, format string, member *gitlab.Member) {
	if format == "json" {
		member.RenderJson(w)
	} else if format == "yaml" {
		member.RenderYaml(w)
	} else {
		// @todo
	}
}
