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
	Use:     "group [group id]",
	Aliases: []string{"g"},
	Short:   "Remove group",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("you must specify a group id")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		groupId := args[0]

		color.Yellow("Removing group (id: %s)…", groupId)

		confirmed := confirmAction(
			fmt.Sprintf("Are you sure you want to remove group %s?", groupId),
			"aborted group removal",
			autoConfirmRemoval,
		)
		if !confirmed {
			return
		}

		loader.Start()
		message, meta, err := client.RemoveGroup(groupId)
		loader.Stop()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		color.Green("✔ Group was successfully removed: %s", message)

		metaOutput(meta, false)
	},
}
