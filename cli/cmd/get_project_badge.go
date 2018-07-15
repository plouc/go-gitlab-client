package cmd

import (
	"strconv"

	"github.com/fatih/color"
	out "github.com/plouc/go-gitlab-client/cli/output"
	"github.com/spf13/cobra"
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

		out.Badge(output, outputFormat, badge)

		printMeta(meta, false)

		return nil
	},
}
