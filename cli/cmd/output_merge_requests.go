package cmd

import (
	"fmt"

	"github.com/olekukonko/tablewriter"
	"github.com/plouc/go-gitlab-client/gitlab"
)

func mergeRequestsOutput(mergeRequests []*gitlab.MergeRequest) {
	if outputFormat == "json" {
		jsonOutput(mergeRequests)
	} else if outputFormat == "yaml" {
		yamlOutput(mergeRequests)
	} else {
		fmt.Fprintln(output, "")
		table := tablewriter.NewWriter(output)
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
		fmt.Fprintln(output, "")
	}
}
