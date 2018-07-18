package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"strconv"
)

func init() {
	rmCmd.AddCommand(rmProjectSnippetNoteCmd)
}

var rmProjectSnippetNoteCmd = &cobra.Command{
	Use:     resourceCmd("project-snippet-note", "project-snippet-note"),
	Aliases: []string{"snippet-note"},
	Short:   "Remove project snippet note",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project-snippet-note", args)
		if err != nil {
			return err
		}

		snippetId, err := strconv.Atoi(ids["snippet_id"])
		if err != nil {
			return err
		}

		noteId, err := strconv.Atoi(ids["note_id"])
		if err != nil {
			return err
		}

		color.Yellow("Removing project snippet note (project id: %s, snippet id: %d, note id: %d)…", ids["project_id"], snippetId, noteId)

		confirmed := confirmAction(
			fmt.Sprintf("Are you sure you want to remove project %s snippet %d note %d?", ids["project_id"], snippetId, noteId),
			"aborted project snippet note removal",
			autoConfirmRemoval,
		)
		if !confirmed {
			return nil
		}

		loader.Start()
		meta, err := client.RemoveProjectSnippetNote(ids["project_id"], snippetId, noteId)
		loader.Stop()
		if err != nil {
			return err
		}

		color.Green("✔ Successfully removed note: %d", noteId)

		printMeta(meta, false)

		return nil
	},
}
