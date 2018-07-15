package output

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/plouc/go-gitlab-client/gitlab"
	"io"
)

func Runners(w io.Writer, format string, collection *gitlab.RunnerCollection) {
	if format == "json" {
		collection.RenderJson(w)
	} else if format == "yaml" {
		collection.RenderYaml(w)
	} else {
		fmt.Fprintln(w, "")
		table := tablewriter.NewWriter(w)
		table.SetHeader([]string{
			"Id",
			"Name",
			"Active",
			"Online",
			"Status",
			"Ip",
			"Shared",
		})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		for _, runner := range collection.Items {
			table.Append([]string{
				strconv.Itoa(runner.Id),
				runner.Name,
				fmt.Sprintf("%t", runner.Active),
				fmt.Sprintf("%t", runner.Online),
				runner.Status,
				runner.IpAddress,
				fmt.Sprintf("%t", runner.IsShared),
			})
		}
		table.Render()
		fmt.Fprintln(w, "")
	}
}

func Runner(w io.Writer, format string, runner *gitlab.RunnerWithDetails) {
	if format == "json" {
		runner.RenderJson(w)
	} else if format == "yaml" {
		runner.RenderYaml(w)
	} else {
		fmt.Fprintln(w, "")

		fmt.Fprintf(w, "  Id              %s\n", color.YellowString("%d", runner.Id))
		fmt.Fprintf(w, "  Name            %s\n", color.YellowString(runner.Name))
		fmt.Fprintf(w, "  Active          %s\n", color.YellowString("%t", runner.Active))
		fmt.Fprintf(w, "  Description     %s\n", color.YellowString(runner.Description))
		fmt.Fprintf(w, "  IpAddress       %s\n", color.YellowString(runner.IpAddress))
		fmt.Fprintf(w, "  IsShared        %s\n", color.YellowString("%t", runner.IsShared))
		fmt.Fprintf(w, "  Online          %s\n", color.YellowString("%t", runner.Online))
		fmt.Fprintf(w, "  Status          %s\n", color.YellowString(runner.Status))
		fmt.Fprintf(w, "  Architecture    %s\n", color.YellowString(runner.Architecture))
		fmt.Fprintf(w, "  Platform        %s\n", color.YellowString(runner.Platform))
		fmt.Fprintf(w, "  Token           %s\n", color.YellowString(runner.Token))
		fmt.Fprintf(w, "  Revision        %s\n", color.YellowString(runner.Revision))
		fmt.Fprintf(w, "  ContactedAt     %s\n", color.YellowString(runner.ContactedAt))
		fmt.Fprintf(w, "  Version         %s\n", color.YellowString(runner.Version))
		fmt.Fprintf(w, "  AccessLevel     %s\n", color.YellowString(runner.AccessLevel))
		fmt.Fprintf(w, "  MaximumTimeout  %s\n", color.YellowString("%d", runner.MaximumTimeout))
		fmt.Fprintf(w, "  TagList         %s\n", color.YellowString(strings.Join(runner.TagList, ", ")))

		fmt.Fprintln(w, "  Projects")
		if len(runner.Projects) > 0 {
			for _, project := range runner.Projects {
				fmt.Fprintf(w, "    %-32s %s\n", color.YellowString(project.Name), fmt.Sprintf("%d", project.Id))
			}
		} else {
			fmt.Fprintln(w, "    No project")
		}

		fmt.Fprintln(w, "")
	}
}
