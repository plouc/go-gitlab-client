package output

import (
	"fmt"
	"io"

	"github.com/olekukonko/tablewriter"
	"github.com/plouc/go-gitlab-client/gitlab"
)

func Commits(w io.Writer, format string, collection *gitlab.CommitCollection) {
	if format == "json" {
		collection.RenderJson(w)
	} else if format == "yaml" {
		collection.RenderYaml(w)
	} else {
		fmt.Fprintln(w, "")
		table := tablewriter.NewWriter(w)
		table.SetHeader([]string{
			"Short Id",
			"Author name",
			"Title",
		})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		for _, commit := range collection.Items {
			table.Append([]string{
				commit.ShortId,
				commit.AuthorName,
				commit.Title,
			})
		}
		table.Render()
		fmt.Fprintln(w, "")
	}
}

func MinimalCommits(w io.Writer, format string, collection *gitlab.MinimalCommitCollection) {
	if format == "json" {
		collection.RenderJson(w)
	} else if format == "yaml" {
		collection.RenderYaml(w)
	} else {
		fmt.Fprintln(w, "")
		table := tablewriter.NewWriter(w)
		table.SetHeader([]string{
			"Short Id",
			"Author name",
			"Title",
		})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		for _, commit := range collection.Items {
			table.Append([]string{
				commit.ShortId,
				commit.AuthorName,
				commit.Title,
			})
		}
		table.Render()
		fmt.Fprintln(w, "")
	}
}
