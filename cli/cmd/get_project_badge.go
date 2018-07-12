package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"strconv"
)

func init() {
	getCmd.AddCommand(getProjectBadgeCmd)
}

var getProjectBadgeCmd = &cobra.Command{
	Use:     resourceCmd("project-badge", "project-badge"),
	Aliases: []string{"pbdg"},
	Short:   "Get project badge info",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project-badge", args)
		if err != nil {
			return err
		}

		badgeId, err := strconv.Atoi(ids["badge_id"])
		if err != nil {
			return err
		}

		color.Yellow("Fetching project's badge (project id: %s, badge id: %s)…", ids["project_id"], badgeId)

		loader.Start()
		badge, meta, err := client.ProjectBadge(ids["project_id"], badgeId)
		loader.Stop()
		if err != nil {
			return err
		}

		badgeOutput(badge)

		metaOutput(meta, false)

		return nil
	},
}
