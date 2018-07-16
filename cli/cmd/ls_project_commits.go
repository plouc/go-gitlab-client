package cmd

import (
	"fmt"

	"github.com/fatih/color"
	out "github.com/plouc/go-gitlab-client/cli/output"
	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
)

func init() {
	lsCmd.AddCommand(lsProjectCommitsCmd)
}

func fetchProjectCommits(projectId string) {
	color.Yellow("Fetching project commits (project id: %s)…", projectId)

	o := &gitlab.CommitsOptions{}
	o.Page = page
	o.PerPage = perPage

	loader.Start()
	collection, meta, err := client.ProjectCommits(projectId, o)
	loader.Stop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(collection.Items) == 0 {
		color.Red("No commit found for project %s", projectId)
	} else {
		out.Commits(output, outputFormat, collection)
	}

	printMeta(meta, true)

	handlePaginatedResult(meta, func() {
		fetchProjectCommits(projectId)
	})
}

var lsProjectCommitsCmd = &cobra.Command{
	Use:     resourceCmd("project-commits", "project"),
	Aliases: []string{"pc"},
	Short:   "List project repository commits",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project", args)
		if err != nil {
			return err
		}

		fetchProjectCommits(ids["project_id"])

		return nil
	},
}
