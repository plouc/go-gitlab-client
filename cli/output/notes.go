package output

import (
	"fmt"

	"github.com/fatih/color"
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
		fmt.Fprintf(w, "  Id              %s\n", color.YellowString("%d", note.Id))
		fmt.Fprintf(w, "  Body            %s\n", color.YellowString(note.Body))
		fmt.Fprintf(w, "  Attachment      %s\n", color.YellowString(note.Attachment))
		fmt.Fprintf(w, "  CreatedAt       %s\n", color.YellowString(note.CreatedAtRaw))
		fmt.Fprintf(w, "  AuthorId        %s\n", color.YellowString("%d", note.Author.Id))
		fmt.Fprintf(w, "  AuthorUsername  %s\n", color.YellowString(note.Author.Username))
		fmt.Fprintf(w, "  AuthorName      %s\n", color.YellowString(note.Author.Name))
		fmt.Fprintln(w, "")
	}
}
