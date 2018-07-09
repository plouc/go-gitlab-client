package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rmCmd.AddCommand(rmProjectBadgeCmd)
}

var rmProjectBadgeCmd = &cobra.Command{
	Use:     "project-badge [project id] [badge id]",
	Aliases: []string{"pbdg"},
	Short:   "Remove project badge",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return fmt.Errorf("you must specify a project id and a badge id")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		projectId := args[0]
		badgeId := args[1]

		color.Yellow("Removing project badge (project id: %s, badge id: %s)…", projectId, badgeId)

		confirmed := confirmAction(
			fmt.Sprintf("Are you sure you want to remove project %s badge %s?", projectId, badgeId),
			"aborted project badge removal",
			autoConfirmRemoval,
		)
		if !confirmed {
			return
		}

		loader.Start()
		meta, err := client.RemoveProjectBadge(projectId, badgeId)
		loader.Stop()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		color.Green("✔ Project badge was successfully removed")

		metaOutput(meta, false)
	},
}
