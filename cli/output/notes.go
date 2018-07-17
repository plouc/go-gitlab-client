package output

import (
	"fmt"

	"github.com/olekukonko/tablewriter"
	"github.com/plouc/go-gitlab-client/gitlab"
	"io"
)

func Notes(w io.Writer, format string, collection *gitlab.NoteCollection) {
	if format == "json" {
		collection.RenderJson(w)
	} else if format == "yaml" {
		collection.RenderYaml(w)
	} else {
		fmt.Fprintln(w, "")
		table := tablewriter.NewWriter(w)
		table.SetHeader([]string{
			"Id",
			"Body",
			"Author",
			"Resolvable",
		})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		for _, note := range collection.Items {
			table.Append([]string{
				fmt.Sprintf("%d", note.Id),
				note.Body,
				note.Author.Name,
				fmt.Sprintf("%t", note.Resolvable),
			})
		}
		table.Render()
		fmt.Fprintln(w, "")
	}
}

func Note(w io.Writer, format string, note *gitlab.Note) {
	if format == "json" {
		note.RenderJson(w)
	} else if format == "yaml" {
		note.RenderYaml(w)
	} else {
		fmt.Fprintln(w, "")
		/*
			fmt.Fprintf(w, "  Id            %s\n", color.YellowString("%d", environment.Id))
			fmt.Fprintf(w, "  Name          %s\n", color.YellowString(environment.Name))
			fmt.Fprintf(w, "  Slug          %s\n", color.YellowString(environment.Slug))
			fmt.Fprintf(w, "  External URL  %s\n", color.YellowString(environment.ExternalUrl))
		*/
		fmt.Fprintln(w, "")
	}
}
