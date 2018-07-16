package output

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/plouc/go-gitlab-client/gitlab"
	"io"
)

func MergeRequests(w io.Writer, format string, collection *gitlab.MergeRequestCollection) {
	if format == "json" {
		collection.RenderJson(w)
	} else if format == "yaml" {
		collection.RenderYaml(w)
	} else {
		fmt.Fprintln(w, "")
		table := tablewriter.NewWriter(w)
		table.SetHeader([]string{
			"Project Id",
			"Id",
			"Iid",
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
		for _, mergeRequest := range collection.Items {
			assignee := ""
			if mergeRequest.Assignee != nil {
				assignee = mergeRequest.Assignee.Username
			}

			table.Append([]string{
				fmt.Sprintf("%d", mergeRequest.ProjectId),
				fmt.Sprintf("%d", mergeRequest.Id),
				fmt.Sprintf("%d", mergeRequest.Iid),
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

func MergeRequest(w io.Writer, format string, mergeRequest *gitlab.MergeRequest) {
	if format == "json" {
		mergeRequest.RenderJson(w)
	} else if format == "yaml" {
		mergeRequest.RenderYaml(w)
	} else {
		fmt.Fprintln(w, "")
		fmt.Fprintf(w, "  Id            %s\n", color.YellowString("%d", mergeRequest.Id))
		fmt.Fprintf(w, "  Iid           %s\n", color.YellowString("%d", mergeRequest.Iid))
		fmt.Fprintf(w, "  Sha           %s\n", color.YellowString(mergeRequest.Sha))
		fmt.Fprintf(w, "  Title         %s\n", color.YellowString(mergeRequest.Title))
		fmt.Fprintf(w, "  SourceBranch  %s\n", color.YellowString(mergeRequest.SourceBranch))
		fmt.Fprintf(w, "  TargetBranch  %s\n", color.YellowString(mergeRequest.TargetBranch))
		fmt.Fprintf(w, "  State         %s\n", color.YellowString(mergeRequest.State))
		fmt.Fprintf(w, "  MergeStatus   %s\n", color.YellowString(mergeRequest.MergeStatus))
		fmt.Fprintln(w, "")
	}
}
