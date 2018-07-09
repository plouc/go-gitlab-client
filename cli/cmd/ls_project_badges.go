package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/plouc/go-gitlab-client/gogitlab"
	"github.com/spf13/cobra"
)

func init() {
	lsCmd.AddCommand(lsProjectBadgesCmd)
}

func fetchProjectBadges(projectId string) {
	color.Yellow("Fetching project badges (id: %s)…", projectId)

	o := &gogitlab.PaginationOptions{}
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
		badgesOutput(badges)
	}

	metaOutput(meta, true)

	handlePaginatedResult(meta, func() {
		fetchProjectBadges(projectId)
	})
}

var lsProjectBadgesCmd = &cobra.Command{
	Use:     "project-badges",
	Aliases: []string{"pbdg"},
	Short:   "List project badges",
	Run: func(cmd *cobra.Command, args []string) {
		fetchProjectBadges(args[0])
	},
}
