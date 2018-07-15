package cmd

import (
	"fmt"

	"github.com/fatih/color"
	out "github.com/plouc/go-gitlab-client/cli/output"
	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
)

func init() {
	lsCmd.AddCommand(lsProjectBadgesCmd)
}

func fetchProjectBadges(projectId string) {
	color.Yellow("Fetching project badges (id: %s)…", projectId)

	o := &gitlab.PaginationOptions{}
	o.Page = page
	o.PerPage = perPage

	loader.Start()
	badges, meta, err := client.ProjectBadges(projectId, o)
	loader.Stop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(badges) == 0 {
		color.Red("No badge found for project %s", projectId)
	} else {
		out.Badges(output, outputFormat, badges)
	}

	out.Meta(meta, true)

	handlePaginatedResult(meta, func() {
		fetchProjectBadges(projectId)
	})
}

var lsProjectBadgesCmd = &cobra.Command{
	Use:     resourceCmd("project-badges", "project"),
	Aliases: []string{"pbdg"},
	Short:   "List project badges",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project", args)
		if err != nil {
			return err
		}

		fetchProjectBadges(ids["project_id"])

		return nil
	},
}
