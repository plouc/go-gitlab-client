package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"strconv"
)

func init() {
	rmCmd.AddCommand(rmGroupEpicNoteCmd)
}

var rmGroupEpicNoteCmd = &cobra.Command{
	Use:     resourceCmd("group-epic-note", "group-epic-note"),
	Aliases: []string{"epic-note"},
	Short:   "Remove group epic note",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "group-epic-note", args)
		if err != nil {
			return err
		}

		epicId, err := strconv.Atoi(ids["epic_id"])
		if err != nil {
			return err
		}

		noteId, err := strconv.Atoi(ids["note_id"])
		if err != nil {
			return err
		}

		color.Yellow("Removing group epic note (group id: %s, epic id: %d, note id: %d)…", ids["group_id"], epicId, noteId)

		confirmed := confirmAction(
			fmt.Sprintf("Are you sure you want to remove group %s epic %d note %d?", ids["group_id"], epicId, noteId),
			"aborted group epic note removal",
			autoConfirmRemoval,
		)
		if !confirmed {
			return nil
		}

		loader.Start()
		meta, err := client.RemoveGroupEpicNote(ids["group_id"], epicId, noteId)
		loader.Stop()
		if err != nil {
			return err
		}

		color.Green("✔ Successfully removed note: %d", noteId)

		printMeta(meta, false)

		return nil
	},
}
