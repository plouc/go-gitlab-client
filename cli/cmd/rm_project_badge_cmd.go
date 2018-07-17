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
	Use:     resourceCmd("project-badge", "project-badge"),
	Aliases: []string{"pbdg"},
	Short:   "Remove project badge",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project-badge", args)
		if err != nil {
			return err
		}

		confirmed := confirmAction(
			fmt.Sprintf("Are you sure you want to remove project %s badge %s?", ids["project_id"], ids["badge_id"]),
			"aborted project badge removal",
			autoConfirmRemoval,
		)
		if !confirmed {
			return nil
		}

		color.Yellow("Removing project badge (project id: %s, badge id: %s)…", ids["project_id"], ids["badge_id"])

		loader.Start()
		meta, err := client.RemoveProjectBadge(ids["project_id"], ids["badge_id"])
		loader.Stop()
		if err != nil {
			return err
		}

		color.Green("✔ Project badge was successfully removed")

		printMeta(meta, false)

		return nil
	},
}
