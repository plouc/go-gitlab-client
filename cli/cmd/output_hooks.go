package cmd

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/plouc/go-gitlab-client/gitlab"
)

func hooksOutput(hooks []*gitlab.Hook) {
	if outputFormat == "json" {
		jsonOutput(hooks)
	} else {
		fmt.Fprintln(output, "")
		table := tablewriter.NewWriter(output)
		table.SetHeader([]string{
			"Id",
			"Url",
			"Push",
			"Issues",
			"Confidential Issues",
			"Merge Requests",
			"Tag Push",
			"Notes",
			"Pipelines",
			"Wiki Pages",
			"SSL check",
		})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		for _, hook := range hooks {
			table.Append([]string{
				strconv.Itoa(hook.Id),
				hook.Url,
				fmt.Sprintf("%t", hook.PushEvents),
				fmt.Sprintf("%t", hook.IssuesEvents),
				fmt.Sprintf("%t", hook.ConfidentialIssuesEvents),
				fmt.Sprintf("%t", hook.MergeRequestsEvents),
				fmt.Sprintf("%t", hook.TagPushEvents),
				fmt.Sprintf("%t", hook.NoteEvents),
				fmt.Sprintf("%t", hook.PipelineEvents),
				fmt.Sprintf("%t", hook.WikiPageEvents),
				fmt.Sprintf("%t", hook.EnableSslVerification),
			})
		}
		table.Render()
		fmt.Fprintln(output, "")
	}
}

func hookOutput(hook *gitlab.Hook) {
	if outputFormat == "json" {
		jsonOutput(hook)
	} else {
		fmt.Fprintln(output, "")
		fmt.Fprintf(output, "  Id                        %s\n", color.YellowString("%d", hook.Id))
		fmt.Fprintf(output, "  ProjectId                 %s\n", color.YellowString("%d", hook.ProjectId))
		fmt.Fprintf(output, "  Url                       %s\n", color.YellowString(hook.Url))
		fmt.Fprintf(output, "  PushEvents                %s\n", color.YellowString("%t", hook.PushEvents))
		fmt.Fprintf(output, "  IssuesEvents              %s\n", color.YellowString("%t", hook.IssuesEvents))
		fmt.Fprintf(output, "  ConfidentialIssuesEvents  %s\n", color.YellowString("%t", hook.ConfidentialIssuesEvents))
		fmt.Fprintf(output, "  MergeRequestsEvents       %s\n", color.YellowString("%t", hook.MergeRequestsEvents))
		fmt.Fprintf(output, "  TagPushEvents             %s\n", color.YellowString("%t", hook.TagPushEvents))
		fmt.Fprintf(output, "  NoteEvents                %s\n", color.YellowString("%t", hook.NoteEvents))
		fmt.Fprintf(output, "  JobEvents                 %s\n", color.YellowString("%t", hook.JobEvents))
		fmt.Fprintf(output, "  PipelineEvents            %s\n", color.YellowString("%t", hook.PipelineEvents))
		fmt.Fprintf(output, "  WikiPageEvents            %s\n", color.YellowString("%t", hook.WikiPageEvents))
		fmt.Fprintf(output, "  EnableSslVerification     %s\n", color.YellowString("%t", hook.EnableSslVerification))
		fmt.Fprintf(output, "  Token                     %s\n", color.YellowString(hook.Token))
		fmt.Fprintf(output, "  CreatedAt                 %s\n", color.YellowString(hook.CreatedAt))
		fmt.Fprintln(output, "")
	}
}
