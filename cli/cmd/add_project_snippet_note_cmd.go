package cmd

import (
	"strconv"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	out "github.com/plouc/go-gitlab-client/cli/output"
	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
)

func init() {
	addCmd.AddCommand(addProjectSnippetNoteCmd)
}

var addProjectSnippetNoteCmd = &cobra.Command{
	Use:     resourceCmd("project-snippet-note", "project-snippet"),
	Aliases: []string{"snippet-note"},
	Short:   "Add project snippet note",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project-snippet", args)
		if err != nil {
			return err
		}

		snippetId, err := strconv.Atoi(ids["snippet_id"])
		if err != nil {
			return err
		}

		color.Yellow("Creating note for project snippet (project id: %s, snippet id: %d)…", ids["project_id"], snippetId)

		note := new(gitlab.NoteAddPayload)

		prompt := promptui.Prompt{
			Label: "Body",
		}
		body, err := prompt.Run()
		if err != nil {
			return err
		}
		note.Body = body

		loader.Start()
		createdNote, meta, err := client.AddProjectSnippetNote(ids["project_id"], snippetId, note)
		loader.Stop()
		if err != nil {
			return err
		}

		out.Note(output, outputFormat, createdNote)

		printMeta(meta, false)

		return nil
	},
}
