package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rmCmd.AddCommand(rmProjectCmd)
}

var rmProjectCmd = &cobra.Command{
	Use:     "project [project id]",
	Aliases: []string{"p"},
	Short:   "Remove project",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("you must specify a project id")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		projectId := args[0]

		color.Yellow("Removing project (id: %s)…", projectId)

		confirmed := confirmAction(
			fmt.Sprintf("Are you sure you want to remove project %s?", projectId),
			"aborted project removal",
			autoConfirmRemoval,
		)
		if !confirmed {
			return
		}

		loader.Start()
		message, meta, err := client.RemoveProject(projectId)
		loader.Stop()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		color.Green("✔ Project was successfully removed: %s", message)

		metaOutput(meta, false)
	},
}
