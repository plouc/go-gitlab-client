package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"strconv"
)

func init() {
	rmCmd.AddCommand(rmProjectIssueNoteCmd)
}

var rmProjectIssueNoteCmd = &cobra.Command{
	Use:     resourceCmd("project-issue-note", "project-issue-note"),
	Aliases: []string{"issue-note"},
	Short:   "Remove project issue note",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project-issue-note", args)
		if err != nil {
			return err
		}

		issueIid, err := strconv.Atoi(ids["issue_iid"])
		if err != nil {
			return err
		}

		noteId, err := strconv.Atoi(ids["note_id"])
		if err != nil {
			return err
		}

		color.Yellow("Removing project issue note (project id: %s, issue iid: %d, note id: %d)…", ids["project_id"], issueIid, noteId)

		confirmed := confirmAction(
			fmt.Sprintf("Are you sure you want to remove project %s issue %d note %d?", ids["project_id"], issueIid, noteId),
			"aborted project issue note removal",
			autoConfirmRemoval,
		)
		if !confirmed {
			return nil
		}

		loader.Start()
		meta, err := client.RemoveProjectIssueNote(ids["project_id"], issueIid, noteId)
		loader.Stop()
		if err != nil {
			return err
		}

		color.Green("✔ Successfully removed note: %d", noteId)

		printMeta(meta, false)

		return nil
	},
}
