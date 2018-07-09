package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/plouc/go-gitlab-client/gitlab"
)

func pipelineOutput(pipeline *gitlab.PipelineWithDetails) {
	if outputFormat == "json" {
		jsonOutput(pipeline)
	} else {
		fmt.Fprintln(output, "")

		fmt.Fprintf(output, "  Id           %s\n", color.YellowString("%d", pipeline.Id))
		fmt.Fprintf(output, "  Status       %s\n", color.YellowString(pipeline.Status))
		fmt.Fprintf(output, "  Ref          %s\n", color.YellowString(pipeline.Ref))
		fmt.Fprintf(output, "  Sha          %s\n", color.YellowString(pipeline.Sha))
		fmt.Fprintf(output, "  BeforeSha    %s\n", color.YellowString(pipeline.BeforeSha))
		fmt.Fprintf(output, "  Tag          %s\n", color.YellowString("%t", pipeline.Tag))
		fmt.Fprintf(output, "  CreatedAt    %s\n", color.YellowString(pipeline.CreatedAt))
		fmt.Fprintf(output, "  UpdatedAt    %s\n", color.YellowString(pipeline.UpdatedAt))
		fmt.Fprintf(output, "  StartedAt    %s\n", color.YellowString(pipeline.StartedAt))
		fmt.Fprintf(output, "  FinishedAt   %s\n", color.YellowString(pipeline.FinishedAt))
		fmt.Fprintf(output, "  CommittedAt  %s\n", color.YellowString(pipeline.CommittedAt))
		fmt.Fprintf(output, "  Duration     %s\n", color.YellowString("%d", pipeline.Duration))
		fmt.Fprintf(output, "  Coverage     %s\n", color.YellowString(pipeline.Coverage))
		fmt.Fprintf(output, "  YamlErrors   %s\n", color.YellowString(pipeline.YamlErrors))

		fmt.Fprintln(output, "  User")
		fmt.Fprintf(output, "    Id         %s\n", color.YellowString("%d", pipeline.User.Id))
		fmt.Fprintf(output, "    Username   %s\n", color.YellowString(pipeline.User.Username))
		fmt.Fprintf(output, "    Name       %s\n", color.YellowString(pipeline.User.Name))
		fmt.Fprintf(output, "    State      %s\n", color.YellowString(pipeline.User.State))
		fmt.Fprintf(output, "    AvatarUrl  %s\n", color.YellowString(pipeline.User.AvatarUrl))
		fmt.Fprintf(output, "    WebUrl     %s\n", color.YellowString(pipeline.User.WebUrl))

		fmt.Fprintln(output, "")
	}
}
