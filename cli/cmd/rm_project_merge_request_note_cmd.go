package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"strconv"
)

func init() {
	rmCmd.AddCommand(rmProjectMergeRequestNoteCmd)
}

var rmProjectMergeRequestNoteCmd = &cobra.Command{
	Use:     resourceCmd("project-merge-request-note", "project-merge-request-note"),
	Aliases: []string{"merge-request-note", "mr-note"},
	Short:   "Remove project merge request note",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project-merge-request-note", args)
		if err != nil {
			return err
		}

		mergeRequestIid, err := strconv.Atoi(ids["merge_request_iid"])
		if err != nil {
			return err
		}

		noteId, err := strconv.Atoi(ids["note_id"])
		if err != nil {
			return err
		}

		color.Yellow("Removing project merge request note (project id: %s, merge request iid: %d, note id: %d)…", ids["project_id"], mergeRequestIid, noteId)

		confirmed := confirmAction(
			fmt.Sprintf("Are you sure you want to remove project %s merge request iid %d note %d?", ids["project_id"], mergeRequestIid, noteId),
			"aborted project merge request note removal",
			autoConfirmRemoval,
		)
		if !confirmed {
			return nil
		}

		loader.Start()
		meta, err := client.RemoveProjectMergeRequestNote(ids["project_id"], mergeRequestIid, noteId)
		loader.Stop()
		if err != nil {
			return err
		}

		color.Green("✔ Successfully removed note: %d", noteId)

		printMeta(meta, false)

		return nil
	},
}
