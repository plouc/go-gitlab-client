package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
	"strconv"
)

func init() {
	lsCmd.AddCommand(lsGroupMergeRequestsCmd)
}

func fetchGroupMergeRequests(groupId int) {
	color.Yellow("Fetching group %d merge requests…", groupId)

	o := &gitlab.MergeRequestsOptions{}
	o.Page = page
	o.PerPage = perPage

	loader.Start()
	mergeRequests, meta, err := client.GroupMergeRequests(groupId, o)
	loader.Stop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(mergeRequests) == 0 {
		color.Red("No merge request found for group %d", groupId)
	} else {
		mergeRequestsOutput(mergeRequests)
	}

	metaOutput(meta, true)

	handlePaginatedResult(meta, func() {
		fetchGroupMergeRequests(groupId)
	})
}

var lsGroupMergeRequestsCmd = &cobra.Command{
	Use:     resourceCmd("group-merge-requests", "group"),
	Aliases: []string{"gmr"},
	Short:   "List group merge requests",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "group", args)
		if err != nil {
			return err
		}

		groupId, err := strconv.Atoi(ids["group_id"])
		if err != nil {
			return err
		}

		fetchGroupMergeRequests(groupId)

		return nil
	},
}
