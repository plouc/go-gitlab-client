package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/plouc/go-gitlab-client/gitlab"
)

func jobsOutput(jobs []*gitlab.Job) {
	if outputFormat == "json" {
		jsonOutput(jobs)
	} else if outputFormat == "yaml" {
		yamlOutput(jobs)
	} else {
		fmt.Fprintln(output, "")
		table := tablewriter.NewWriter(output)
		table.SetHeader([]string{
			"Pipeline id",
			"Stage",
			"Id",
			"Name",
			"Status",
			"Started at",
			"Finished at",
			"Duration",
		})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		for _, job := range jobs {
			table.Append([]string{
				fmt.Sprintf("%d", job.Pipeline.Id),
				job.Stage,
				fmt.Sprintf("%d", job.Id),
				job.Name,
				job.Status,
				job.StartedAt,
				job.FinishedAt,
				fmt.Sprintf("%f", job.Duration),
			})
		}
		table.Render()
		fmt.Fprintln(output, "")
	}
}

func jobOutput(job *gitlab.Job) {
	if outputFormat == "json" {
		jsonOutput(job)
	} else if outputFormat == "yaml" {
		yamlOutput(job)
	} else {
		fmt.Fprintln(output, "")
		fmt.Fprintf(output, "  Id          %s\n", color.YellowString("%d", job.Id))
		fmt.Fprintf(output, "  PipelineId  %s\n", color.YellowString("%d", job.Pipeline.Id))
		fmt.Fprintf(output, "  Stage       %s\n", color.YellowString(job.Stage))
		fmt.Fprintf(output, "  Name        %s\n", color.YellowString(job.Name))
		fmt.Fprintf(output, "  Status      %s\n", color.YellowString(job.Status))
		fmt.Fprintf(output, "  StartedAt   %s\n", color.YellowString(job.StartedAt))
		fmt.Fprintf(output, "  FinishedAt  %s\n", color.YellowString(job.FinishedAt))
		fmt.Fprintf(output, "  Duration    %s\n", color.YellowString("%f", job.Duration))
		fmt.Fprintln(output, "")
	}
}
