package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/plouc/go-gitlab-client/gitlab"
)

func namespacesOutput(namespaces []*gitlab.Namespace) {
	if outputFormat == "json" {
		jsonOutput(namespaces)
	} else if outputFormat == "yaml" {
		yamlOutput(namespaces)
	} else {
		fmt.Fprintln(output, "")
		table := tablewriter.NewWriter(output)
		table.SetHeader([]string{
			"Id",
			"Name",
			"Path",
			"Kind",
			"Full path",
		})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		for _, namespace := range namespaces {
			table.Append([]string{
				fmt.Sprintf("%d", namespace.Id),
				namespace.Name,
				namespace.Path,
				namespace.Kind,
				namespace.FullPath,
			})
		}
		table.Render()
		fmt.Fprintln(output, "")
	}
}

func namespaceOutput(namespace *gitlab.Namespace) {
	if outputFormat == "json" {
		jsonOutput(namespace)
	} else {
		fmt.Fprintln(output, "")
		fmt.Fprintf(output, "  Id                           %s\n", color.YellowString("%d", namespace.Id))
		fmt.Fprintf(output, "  Name                         %s\n", color.YellowString(namespace.Name))
		fmt.Fprintf(output, "  Path                         %s\n", color.YellowString(namespace.Path))
		fmt.Fprintf(output, "  Kind                         %s\n", color.YellowString(namespace.Kind))
		fmt.Fprintf(output, "  FullPath                     %s\n", color.YellowString(namespace.FullPath))
		fmt.Fprintf(output, "  ParentId                     %s\n", color.YellowString("%d", namespace.ParentId))
		fmt.Fprintf(output, "  MembersCountWithDescendants  %s\n", color.YellowString("%d", namespace.MembersCountWithDescendants))
		fmt.Fprintf(output, "  Plan                         %s\n", color.YellowString(namespace.Plan))
		fmt.Fprintf(output, "  Description                  %s\n", color.YellowString(namespace.Description))
		fmt.Fprintf(output, "  OwnerId                      %s\n", color.YellowString("%d", namespace.OwnerId))
		fmt.Fprintf(output, "  CreatedAt                    %s\n", color.YellowString(namespace.CreatedAt))
		fmt.Fprintf(output, "  UpdatedAt                    %s\n", color.YellowString(namespace.UpdatedAt))
		fmt.Fprintln(output, "")
	}
}
