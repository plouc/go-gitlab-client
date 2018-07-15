package output

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/plouc/go-gitlab-client/gitlab"
	"io"
)

func Pipeline(w io.Writer, format string, pipeline *gitlab.PipelineWithDetails) {
	if format == "json" {
		Json(w, pipeline)
	} else if format == "yaml" {
		Yaml(w, pipeline)
	} else {
		fmt.Fprintln(w, "")

		fmt.Fprintf(w, "  Id           %s\n", color.YellowString("%d", pipeline.Id))
		fmt.Fprintf(w, "  Status       %s\n", color.YellowString(pipeline.Status))
		fmt.Fprintf(w, "  Ref          %s\n", color.YellowString(pipeline.Ref))
		fmt.Fprintf(w, "  Sha          %s\n", color.YellowString(pipeline.Sha))
		fmt.Fprintf(w, "  BeforeSha    %s\n", color.YellowString(pipeline.BeforeSha))
		fmt.Fprintf(w, "  Tag          %s\n", color.YellowString("%t", pipeline.Tag))
		fmt.Fprintf(w, "  CreatedAt    %s\n", color.YellowString(pipeline.CreatedAt))
		fmt.Fprintf(w, "  UpdatedAt    %s\n", color.YellowString(pipeline.UpdatedAt))
		fmt.Fprintf(w, "  StartedAt    %s\n", color.YellowString(pipeline.StartedAt))
		fmt.Fprintf(w, "  FinishedAt   %s\n", color.YellowString(pipeline.FinishedAt))
		fmt.Fprintf(w, "  CommittedAt  %s\n", color.YellowString(pipeline.CommittedAt))
		fmt.Fprintf(w, "  Duration     %s\n", color.YellowString("%d", pipeline.Duration))
		fmt.Fprintf(w, "  Coverage     %s\n", color.YellowString(pipeline.Coverage))
		fmt.Fprintf(w, "  YamlErrors   %s\n", color.YellowString(pipeline.YamlErrors))

		fmt.Fprintln(w, "  User")
		fmt.Fprintf(w, "    Id         %s\n", color.YellowString("%d", pipeline.User.Id))
		fmt.Fprintf(w, "    Username   %s\n", color.YellowString(pipeline.User.Username))
		fmt.Fprintf(w, "    Name       %s\n", color.YellowString(pipeline.User.Name))
		fmt.Fprintf(w, "    State      %s\n", color.YellowString(pipeline.User.State))
		fmt.Fprintf(w, "    AvatarUrl  %s\n", color.YellowString(pipeline.User.AvatarUrl))
		fmt.Fprintf(w, "    WebUrl     %s\n", color.YellowString(pipeline.User.WebUrl))

		fmt.Fprintln(w, "")
	}
}

func Pipelines(w io.Writer, format string, pipelines []*gitlab.Pipeline) {
	if format == "json" {
		Json(w, pipelines)
	} else if format == "yaml" {
		Yaml(w, pipelines)
	} else {
		fmt.Fprintln(w, "")
		table := tablewriter.NewWriter(w)
		table.SetHeader([]string{
			"Id",
			"Ref",
			"Sha",
			"Status",
		})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		for _, pipeline := range pipelines {
			table.Append([]string{
				fmt.Sprintf("%d", pipeline.Id),
				pipeline.Ref,
				pipeline.Sha,
				pipeline.Status,
			})
		}
		table.Render()
		fmt.Fprintln(w, "")
	}
}
