package output

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/plouc/go-gitlab-client/gitlab"
	"io"
)

func Namespaces(w io.Writer, format string, collection *gitlab.NamespaceCollection) {
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
			"Kind",
			"Full path",
		})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		for _, namespace := range collection.Items {
			table.Append([]string{
				fmt.Sprintf("%d", namespace.Id),
				namespace.Name,
				namespace.Path,
				namespace.Kind,
				namespace.FullPath,
			})
		}
		table.Render()
		fmt.Fprintln(w, "")
	}
}

func Namespace(w io.Writer, format string, namespace *gitlab.Namespace) {
	if format == "json" {
		namespace.RenderJson(w)
	} else if format == "yaml " {
		namespace.RenderYaml(w)
	} else {
		fmt.Fprintln(w, "")
		fmt.Fprintf(w, "  Id                           %s\n", color.YellowString("%d", namespace.Id))
		fmt.Fprintf(w, "  Name                         %s\n", color.YellowString(namespace.Name))
		fmt.Fprintf(w, "  Path                         %s\n", color.YellowString(namespace.Path))
		fmt.Fprintf(w, "  Kind                         %s\n", color.YellowString(namespace.Kind))
		fmt.Fprintf(w, "  FullPath                     %s\n", color.YellowString(namespace.FullPath))
		fmt.Fprintf(w, "  ParentId                     %s\n", color.YellowString("%d", namespace.ParentId))
		fmt.Fprintf(w, "  MembersCountWithDescendants  %s\n", color.YellowString("%d", namespace.MembersCountWithDescendants))
		fmt.Fprintf(w, "  Plan                         %s\n", color.YellowString(namespace.Plan))
		fmt.Fprintf(w, "  Description                  %s\n", color.YellowString(namespace.Description))
		fmt.Fprintf(w, "  OwnerId                      %s\n", color.YellowString("%d", namespace.OwnerId))
		fmt.Fprintf(w, "  CreatedAt                    %s\n", color.YellowString(namespace.CreatedAt))
		fmt.Fprintf(w, "  UpdatedAt                    %s\n", color.YellowString(namespace.UpdatedAt))
		fmt.Fprintln(w, "")
	}
}
