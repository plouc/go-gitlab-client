package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/plouc/go-gitlab-client/gogitlab"
)

func groupsOutput(groups []*gogitlab.Group) {
	if outputFormat == "json" {
		jsonOutput(groups)
	} else if outputFormat == "yaml" {
		yamlOutput(groups)
	} else {
		fmt.Fprintln(output, "")
		table := tablewriter.NewWriter(output)
		table.SetHeader([]string{
			"Id",
			"Name",
			"Path",
			"Visibility",
		})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		for _, group := range groups {
			table.Append([]string{
				fmt.Sprintf("%d", group.Id),
				group.Name,
				group.Path,
				fmt.Sprintf("%s", group.Visibility),
			})
		}
		table.Render()
		fmt.Fprintln(output, "")
	}
}

func groupOutput(group *gogitlab.GroupWithDetails) {
	if outputFormat == "json" {
		jsonOutput(group)
	} else if outputFormat == "yaml" {
		yamlOutput(group)
	} else {
		fmt.Fprintln(output, "")

		fmt.Fprintf(output, "  Id                    %s\n", color.YellowString("%d", group.Id))
		fmt.Fprintf(output, "  Name                  %s\n", color.YellowString(group.Name))
		fmt.Fprintf(output, "  Path                  %s\n", color.YellowString(group.Path))
		fmt.Fprintf(output, "  Description           %s\n", color.YellowString(group.Description))
		fmt.Fprintf(output, "  Visibility            %s\n", color.YellowString("%s", group.Visibility))
		fmt.Fprintf(output, "  LfsEnabled            %s\n", color.YellowString("%t", group.LfsEnabled))
		fmt.Fprintf(output, "  AvatarUrl             %s\n", color.YellowString(group.AvatarUrl))
		fmt.Fprintf(output, "  WebURL                %s\n", color.YellowString(group.WebURL))
		fmt.Fprintf(output, "  RequestAccessEnabled  %s\n", color.YellowString("%t", group.RequestAccessEnabled))
		fmt.Fprintf(output, "  FullName              %s\n", color.YellowString(group.FullName))
		fmt.Fprintf(output, "  FullPath              %s\n", color.YellowString(group.FullPath))
		fmt.Fprintf(output, "  ParentId              %s\n", color.YellowString("%d", group.ParentId))

		fmt.Fprintln(output, "  Projects")
		if len(group.Projects) > 0 {
			for _, project := range group.Projects {
				fmt.Fprintf(output, "    %-32s %s\n", color.YellowString(project.Name), fmt.Sprintf("%d", project.Id))
			}
		} else {
			fmt.Fprintln(output, "    No project")
		}

		fmt.Fprintln(output, "  SharedProjects")
		if len(group.SharedProjects) > 0 {
			for _, project := range group.SharedProjects {
				fmt.Fprintf(output, "    %-32s %s\n", color.YellowString(project.Name), fmt.Sprintf("%d", project.Id))
			}
		} else {
			fmt.Fprintln(output, "    No shared project")
		}

		fmt.Fprintln(output, "")
	}
}
