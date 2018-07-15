package cmd

import (
	"fmt"

	"github.com/fatih/color"
	out "github.com/plouc/go-gitlab-client/cli/output"
	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
)

func init() {
	lsCmd.AddCommand(lsProjectEnvironmentsCmd)
}

func fetchProjectEnvironments(projectId string) {
	color.Yellow("Fetching project environments (project id: %s)…", projectId)

	o := &gitlab.PaginationOptions{}
	o.Page = page
	o.PerPage = perPage

	loader.Start()
	collection, meta, err := client.ProjectEnvironments(projectId, o)
	loader.Stop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	i := len(collection.Items)
	if i == 0 {
		color.Red("No environment found for project %s", projectId)
	} else {
		out.Environments(output, outputFormat, collection)
	}

	printMeta(meta, true)

	handlePaginatedResult(meta, func() {
		fetchProjectEnvironments(projectId)
	})
}

var lsProjectEnvironmentsCmd = &cobra.Command{
	Use:     resourceCmd("project-environments", "project"),
	Aliases: []string{"project-envs", "pe"},
	Short:   "List project environments",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project", args)
		if err != nil {
			return err
		}

		fetchProjectEnvironments(ids["project_id"])

		return nil
	},
}
