package output

import (
	"fmt"

	"github.com/olekukonko/tablewriter"
	"github.com/plouc/go-gitlab-client/gitlab"
	"io"
)

func MergeRequests(w io.Writer, format string, mergeRequests []*gitlab.MergeRequest) {
	if format == "json" {
		Json(w, mergeRequests)
	} else if format == "yaml" {
		Yaml(w, mergeRequests)
	} else {
		fmt.Fprintln(w, "")
		table := tablewriter.NewWriter(w)
		table.SetHeader([]string{
			"Project Id",
			"Id",
			"Title",
			"Source",
			"Target",
			"State",
			"Assignee",
			"WIP",
			"Merge status",
			"Created at",
		})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		for _, mergeRequest := range mergeRequests {
			assignee := ""
			if mergeRequest.Assignee != nil {
				assignee = mergeRequest.Assignee.Username
			}

			table.Append([]string{
				fmt.Sprintf("%d", mergeRequest.ProjectId),
				fmt.Sprintf("%d", mergeRequest.Id),
				mergeRequest.Title,
				mergeRequest.SourceBranch,
				mergeRequest.TargetBranch,
				mergeRequest.State,
				assignee,
				fmt.Sprintf("%t", mergeRequest.WorkInProgress),
				mergeRequest.MergeStatus,
				mergeRequest.CreatedAt,
			})
		}
		table.Render()
		fmt.Fprintln(w, "")
	}
}
