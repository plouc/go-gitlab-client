package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	getCmd.AddCommand(getProjectBadgeCmd)
}

var getProjectBadgeCmd = &cobra.Command{
	Use:     "project-badge [project id] [badge id]",
	Aliases: []string{"pbdg"},
	Short:   "Get project badge info",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return fmt.Errorf("you must specify a project id and a badge id")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		projectId := args[0]
		badgeId := args[1]

		color.Yellow("Fetching project's badge (project id: %s, badge id: %s)…", projectId, badgeId)

		loader.Start()
		badge, meta, err := client.ProjectBadge(projectId, badgeId)
		loader.Stop()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		badgeOutput(badge)

		metaOutput(meta, false)
	},
}
