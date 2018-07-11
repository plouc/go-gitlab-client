package cmd

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/plouc/textree"
)

func jobLabel(job *gitlab.Job) string {
	statusIcon := color.YellowString("?")
	if job.Status == "skipped" {
		statusIcon = "-"
	} else if job.Status == "success" {
		statusIcon = color.GreenString("✔")
	} else if job.Status == "failed" {
		statusIcon = color.RedString("✘")
	}

	return fmt.Sprintf(
		"%s %-16s %3ds [%-7s] %d",
		statusIcon,
		job.Name,
		int64(job.Duration),
		job.Status,
		job.Id,
	)
}

func jobsOutput(jobs []*gitlab.Job, pretty bool) {
	if pretty {
		agg := client.AggregateJobs(jobs)

		root := textree.NewNode("PIPELINES")

		for pipelineId, stages := range agg {
			p := textree.NewNode(fmt.Sprintf("PIPELINE [%d]", pipelineId))
			root.Append(p)

			for stage, jobsByName := range stages {
				s := textree.NewNode(strings.ToUpper(stage))
				p.Append(s)

				for _, similarJobs := range jobsByName {
					if len(similarJobs) > 0 {
						n := textree.NewNode(jobLabel(similarJobs[0]))
						s.Append(n)

						for _, job := range similarJobs[1:] {
							n.Append(textree.NewNode(jobLabel(job)))
						}
					}
				}
			}
		}

		o := textree.NewRenderOptions()
		o.ChildrenMarginTop = 1
		enableCustomStyle := false
		if enableCustomStyle {
			o.HorizontalLink = color.YellowString(o.HorizontalLink)
			o.VerticalLink = color.YellowString(o.VerticalLink)
			o.ChildrenLink = color.YellowString(o.ChildrenLink)
			o.ChildLink = color.YellowString(o.ChildLink)
			o.LastChildLink = color.YellowString(o.LastChildLink)
		}
		root.Render(output, o)

	} else if outputFormat == "json" {
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
