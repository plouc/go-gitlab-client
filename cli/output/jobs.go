package output

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/plouc/textree"
	"io"
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

func Jobs(w io.Writer, format string, collection *gitlab.JobCollection, pretty bool) {
	if pretty {
		agg := gitlab.AggregateJobs(collection.Items)

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
		root.Render(w, o)

	} else if format == "json" {
		collection.RenderJson(w)
	} else if format == "yaml" {
		collection.RenderYaml(w)
	} else {
		fmt.Fprintln(w, "")
		table := tablewriter.NewWriter(w)
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
		for _, job := range collection.Items {
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
		fmt.Fprintln(w, "")
	}
}

func Job(w io.Writer, format string, job *gitlab.Job) {
	if format == "json" {
		job.RenderJson(w)
	} else if format == "yaml" {
		job.RenderYaml(w)
	} else {
		fmt.Fprintln(w, "")
		fmt.Fprintf(w, "  Id          %s\n", color.YellowString("%d", job.Id))
		fmt.Fprintf(w, "  PipelineId  %s\n", color.YellowString("%d", job.Pipeline.Id))
		fmt.Fprintf(w, "  Stage       %s\n", color.YellowString(job.Stage))
		fmt.Fprintf(w, "  Name        %s\n", color.YellowString(job.Name))
		fmt.Fprintf(w, "  Status      %s\n", color.YellowString(job.Status))
		fmt.Fprintf(w, "  StartedAt   %s\n", color.YellowString(job.StartedAt))
		fmt.Fprintf(w, "  FinishedAt  %s\n", color.YellowString(job.FinishedAt))
		fmt.Fprintf(w, "  Duration    %s\n", color.YellowString("%f", job.Duration))
		fmt.Fprintln(w, "")
	}
}
