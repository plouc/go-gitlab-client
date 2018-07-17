package cmd

import (
	"fmt"

	"github.com/fatih/color"
	out "github.com/plouc/go-gitlab-client/cli/output"
	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
	"strconv"
)

func init() {
	listCmd.AddCommand(listProjectSnippetNotesCmd)
}

func fetchProjectSnippetNotes(projectId string, snippetId int) {
	color.Yellow("Fetching project snippet notes (project id: %s, snippet id: %d)…", projectId, snippetId)

	o := &gitlab.NotesOptions{}
	o.Page = page
	o.PerPage = perPage

	loader.Start()
	collection, meta, err := client.ProjectSnippetNotes(projectId, snippetId, o)
	loader.Stop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(collection.Items) == 0 {
		color.Red("No note found for project %s snippet %d", projectId, snippetId)
	} else {
		out.Notes(output, outputFormat, collection)
	}

	printMeta(meta, true)

	handlePaginatedResult(meta, func() {
		fetchProjectSnippetNotes(projectId, snippetId)
	})
}

var listProjectSnippetNotesCmd = &cobra.Command{
	Use:     resourceCmd("project-snippet-notes", "project-snippet"),
	Aliases: []string{"snippet-notes"},
	Short:   "List project snippet notes",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project-snippet", args)
		if err != nil {
			return err
		}

		snippetId, err := strconv.Atoi(ids["snippet_id"])
		if err != nil {
			return err
		}

		fetchProjectSnippetNotes(ids["project_id"], snippetId)

		return nil
	},
}
