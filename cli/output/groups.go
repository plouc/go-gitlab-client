package output

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/plouc/go-gitlab-client/gitlab"
	"io"
)

func Groups(w io.Writer, format string, collection *gitlab.GroupCollection) {
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
			"Path",
			"Visibility",
		})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		for _, group := range collection.Items {
			table.Append([]string{
				fmt.Sprintf("%d", group.Id),
				group.Name,
				group.Path,
				fmt.Sprintf("%s", group.Visibility),
			})
		}
		table.Render()
		fmt.Fprintln(w, "")
	}
}

func Group(w io.Writer, format string, group *gitlab.GroupWithDetails) {
	if format == "json" {
		group.RenderJson(w)
	} else if format == "yaml" {
		group.RenderYaml(w)
	} else {
		fmt.Fprintln(w, "")

		fmt.Fprintf(w, "  Id                    %s\n", color.YellowString("%d", group.Id))
		fmt.Fprintf(w, "  Name                  %s\n", color.YellowString(group.Name))
		fmt.Fprintf(w, "  Path                  %s\n", color.YellowString(group.Path))
		fmt.Fprintf(w, "  Description           %s\n", color.YellowString(group.Description))
		fmt.Fprintf(w, "  Visibility            %s\n", color.YellowString("%s", group.Visibility))
		fmt.Fprintf(w, "  LfsEnabled            %s\n", color.YellowString("%t", group.LfsEnabled))
		fmt.Fprintf(w, "  AvatarUrl             %s\n", color.YellowString(group.AvatarUrl))
		fmt.Fprintf(w, "  WebURL                %s\n", color.YellowString(group.WebURL))
		fmt.Fprintf(w, "  RequestAccessEnabled  %s\n", color.YellowString("%t", group.RequestAccessEnabled))
		fmt.Fprintf(w, "  FullName              %s\n", color.YellowString(group.FullName))
		fmt.Fprintf(w, "  FullPath              %s\n", color.YellowString(group.FullPath))
		fmt.Fprintf(w, "  ParentId              %s\n", color.YellowString("%d", group.ParentId))

		fmt.Fprintln(w, "  Projects")
		if len(group.Projects) > 0 {
			for _, project := range group.Projects {
				fmt.Fprintf(w, "    %-32s %s\n", color.YellowString(project.Name), fmt.Sprintf("%d", project.Id))
			}
		} else {
			fmt.Fprintln(w, "    No project")
		}

		fmt.Fprintln(w, "  SharedProjects")
		if len(group.SharedProjects) > 0 {
			for _, project := range group.SharedProjects {
				fmt.Fprintf(w, "    %-32s %s\n", color.YellowString(project.Name), fmt.Sprintf("%d", project.Id))
			}
		} else {
			fmt.Fprintln(w, "    No shared project")
		}

		fmt.Fprintln(w, "")
	}
}
