package cmd

import (
	"fmt"

	"github.com/fatih/color"
	out "github.com/plouc/go-gitlab-client/cli/output"
	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
	"strconv"
)

func init() {
	listCmd.AddCommand(listProjectMergeRequestCommitsCmd)
}

func fetchProjectMergeRequestCommits(projectId string, mergeRequestIid int) {
	color.Yellow("Fetching project merge request commits (project id: %s, merge request iid: %d)…", projectId, mergeRequestIid)

	o := &gitlab.PaginationOptions{}
	o.Page = page
	o.PerPage = perPage

	loader.Start()
	collection, meta, err := client.ProjectMergeRequestCommits(projectId, mergeRequestIid, o)
	loader.Stop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(collection.Items) == 0 {
		color.Red("No commit found for project %s merge request %d", projectId, mergeRequestIid)
	} else {
		out.MinimalCommits(output, outputFormat, collection)
	}

	printMeta(meta, true)

	handlePaginatedResult(meta, func() {
		fetchProjectMergeRequestCommits(projectId, mergeRequestIid)
	})
}

var listProjectMergeRequestCommitsCmd = &cobra.Command{
	Use:     resourceCmd("project-merge-request-commits", "project-merge-request"),
	Aliases: []string{"project-mr-commits"},
	Short:   "List project merge request commits",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project-merge-request", args)
		if err != nil {
			return err
		}

		mergeRequestIid, err := strconv.Atoi(ids["merge_request_iid"])
		if err != nil {
			return err
		}

		fetchProjectMergeRequestCommits(ids["project_id"], mergeRequestIid)

		return nil
	},
}
