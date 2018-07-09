package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/plouc/go-gitlab-client/gitlab"
)

func runnersOutput(runners []*gitlab.Runner) {
	if outputFormat == "json" {
		jsonOutput(runners)
	} else if outputFormat == "yaml" {
		yamlOutput(runners)
	} else {
		fmt.Fprintln(output, "")
		table := tablewriter.NewWriter(output)
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
		for _, runner := range runners {
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
		fmt.Fprintln(output, "")
	}
}

func runnerOutput(runner *gitlab.RunnerWithDetails) {
	if outputFormat == "json" {
		jsonOutput(runner)
	} else if outputFormat == "yaml" {
		yamlOutput(runner)
	} else {
		fmt.Fprintln(output, "")

		fmt.Fprintf(output, "  Id              %s\n", color.YellowString("%d", runner.Id))
		fmt.Fprintf(output, "  Name            %s\n", color.YellowString(runner.Name))
		fmt.Fprintf(output, "  Active          %s\n", color.YellowString("%t", runner.Active))
		fmt.Fprintf(output, "  Description     %s\n", color.YellowString(runner.Description))
		fmt.Fprintf(output, "  IpAddress       %s\n", color.YellowString(runner.IpAddress))
		fmt.Fprintf(output, "  IsShared        %s\n", color.YellowString("%t", runner.IsShared))
		fmt.Fprintf(output, "  Online          %s\n", color.YellowString("%t", runner.Online))
		fmt.Fprintf(output, "  Status          %s\n", color.YellowString(runner.Status))
		fmt.Fprintf(output, "  Architecture    %s\n", color.YellowString(runner.Architecture))
		fmt.Fprintf(output, "  Platform        %s\n", color.YellowString(runner.Platform))
		fmt.Fprintf(output, "  Token           %s\n", color.YellowString(runner.Token))
		fmt.Fprintf(output, "  Revision        %s\n", color.YellowString(runner.Revision))
		fmt.Fprintf(output, "  ContactedAt     %s\n", color.YellowString(runner.ContactedAt))
		fmt.Fprintf(output, "  Version         %s\n", color.YellowString(runner.Version))
		fmt.Fprintf(output, "  AccessLevel     %s\n", color.YellowString(runner.AccessLevel))
		fmt.Fprintf(output, "  MaximumTimeout  %s\n", color.YellowString("%d", runner.MaximumTimeout))
		fmt.Fprintf(output, "  TagList         %s\n", color.YellowString(strings.Join(runner.TagList, ", ")))

		fmt.Fprintln(output, "  Projects")
		if len(runner.Projects) > 0 {
			for _, project := range runner.Projects {
				fmt.Fprintf(output, "    %-32s %s\n", color.YellowString(project.Name), fmt.Sprintf("%d", project.Id))
			}
		} else {
			fmt.Fprintln(output, "    No project")
		}

		fmt.Fprintln(output, "")
	}
}
