package cmd

import (
	"fmt"

	"github.com/fatih/color"
	out "github.com/plouc/go-gitlab-client/cli/output"
	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
)

func init() {
	lsCmd.AddCommand(lsProjectMergeRequestsCmd)
}

func fetchProjectMergeRequests(projectId string) {
	color.Yellow("Fetching project %s merge requests…", projectId)

	o := &gitlab.MergeRequestsOptions{}
	o.Page = page
	o.PerPage = perPage

	loader.Start()
	mergeRequests, meta, err := client.ProjectMergeRequests(projectId, o)
	loader.Stop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(mergeRequests) == 0 {
		color.Red("No merge request found for project %s", projectId)
	} else {
		out.MergeRequests(output, outputFormat, mergeRequests)
	}

	metaOutput(meta, true)

	handlePaginatedResult(meta, func() {
		fetchProjectMergeRequests(projectId)
	})
}

var lsProjectMergeRequestsCmd = &cobra.Command{
	Use:     resourceCmd("project-merge-requests", "project"),
	Aliases: []string{"pmr"},
	Short:   "List project merge requests",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project", args)
		if err != nil {
			return err
		}

		fetchProjectMergeRequests(ids["project_id"])

		return nil
	},
}
