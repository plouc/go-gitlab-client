package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
)

func init() {
	lsCmd.AddCommand(lsMergeRequestsCmd)
}

func fetchMergeRequests() {
	color.Yellow("Fetching merge requests…")

	o := &gitlab.MergeRequestsOptions{}
	o.Page = page
	o.PerPage = perPage

	loader.Start()
	mergeRequests, meta, err := client.MergeRequests(o)
	loader.Stop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(mergeRequests) == 0 {
		color.Red("No merge request found")
	} else {
		mergeRequestsOutput(mergeRequests)
	}

	metaOutput(meta, true)

	handlePaginatedResult(meta, fetchMergeRequests)
}

var lsMergeRequestsCmd = &cobra.Command{
	Use:     "merge-requests",
	Aliases: []string{"mr"},
	Short:   "List merge requests",
	Run: func(cmd *cobra.Command, args []string) {
		fetchMergeRequests()
	},
}
