package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rmCmd.AddCommand(rmGroupVarCmd)
}

var rmGroupVarCmd = &cobra.Command{
	Use:     "group-var [group id] [var key]",
	Aliases: []string{"gv"},
	Short:   "Remove a group's variable",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return fmt.Errorf("you must specify a group id and variable key")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		groupId := args[0]
		varKey := args[1]

		color.Yellow("Removing group variable (group id: %s, key: %s)…", groupId, varKey)

		confirmed := confirmAction(
			fmt.Sprintf("Are you sure you want to remove group %s variable %s?", groupId, varKey),
			"aborted group variable removal",
			autoConfirmRemoval,
		)
		if !confirmed {
			return
		}

		loader.Start()
		meta, err := client.RemoveGroupVariable(groupId, varKey)
		loader.Stop()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		color.Green("✔ Successfully removed variable: %s", varKey)

		metaOutput(meta, false)
	},
}
