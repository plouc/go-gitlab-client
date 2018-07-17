package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rmCmd.AddCommand(rmGroupCmd)
}

var rmGroupCmd = &cobra.Command{
	Use:     resourceCmd("group", "group"),
	Aliases: []string{"g"},
	Short:   "Remove group",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "group", args)
		if err != nil {
			return err
		}

		color.Yellow("Removing group (id: %s)…", ids["group_id"])

		confirmed := confirmAction(
			fmt.Sprintf("Are you sure you want to remove group %s?", ids["group_id"]),
			"aborted group removal",
			autoConfirmRemoval,
		)
		if !confirmed {
			return nil
		}

		loader.Start()
		message, meta, err := client.RemoveGroup(ids["group_id"])
		loader.Stop()
		if err != nil {
			return err
		}

		color.Green("✔ Group was successfully removed: %s", message)

		printMeta(meta, false)

		return nil
	},
}
