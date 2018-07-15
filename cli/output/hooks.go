package output

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/plouc/go-gitlab-client/gitlab"
	"io"
)

func Hooks(w io.Writer, format string, collection *gitlab.HookCollection) {
	if format == "json" {
		collection.RenderJson(w)
	} else if format == "yaml" {
		collection.RenderYaml(w)
	} else {
		fmt.Fprintln(w, "")
		table := tablewriter.NewWriter(w)
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
		for _, hook := range collection.Items {
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
		fmt.Fprintln(w, "")
	}
}

func Hook(w io.Writer, format string, hook *gitlab.Hook) {
	if format == "json" {
		hook.RenderJson(w)
	} else if format == "yaml" {
		hook.RenderYaml(w)
	} else {
		fmt.Fprintln(w, "")
		fmt.Fprintf(w, "  Id                        %s\n", color.YellowString("%d", hook.Id))
		fmt.Fprintf(w, "  ProjectId                 %s\n", color.YellowString("%d", hook.ProjectId))
		fmt.Fprintf(w, "  Url                       %s\n", color.YellowString(hook.Url))
		fmt.Fprintf(w, "  PushEvents                %s\n", color.YellowString("%t", hook.PushEvents))
		fmt.Fprintf(w, "  IssuesEvents              %s\n", color.YellowString("%t", hook.IssuesEvents))
		fmt.Fprintf(w, "  ConfidentialIssuesEvents  %s\n", color.YellowString("%t", hook.ConfidentialIssuesEvents))
		fmt.Fprintf(w, "  MergeRequestsEvents       %s\n", color.YellowString("%t", hook.MergeRequestsEvents))
		fmt.Fprintf(w, "  TagPushEvents             %s\n", color.YellowString("%t", hook.TagPushEvents))
		fmt.Fprintf(w, "  NoteEvents                %s\n", color.YellowString("%t", hook.NoteEvents))
		fmt.Fprintf(w, "  JobEvents                 %s\n", color.YellowString("%t", hook.JobEvents))
		fmt.Fprintf(w, "  PipelineEvents            %s\n", color.YellowString("%t", hook.PipelineEvents))
		fmt.Fprintf(w, "  WikiPageEvents            %s\n", color.YellowString("%t", hook.WikiPageEvents))
		fmt.Fprintf(w, "  EnableSslVerification     %s\n", color.YellowString("%t", hook.EnableSslVerification))
		fmt.Fprintf(w, "  Token                     %s\n", color.YellowString(hook.Token))
		fmt.Fprintf(w, "  CreatedAt                 %s\n", color.YellowString(hook.CreatedAt))
		fmt.Fprintln(w, "")
	}
}
