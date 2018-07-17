package cmd

import (
	"fmt"

	"github.com/fatih/color"
	out "github.com/plouc/go-gitlab-client/cli/output"
	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
)

func init() {
	listCmd.AddCommand(listMergeRequestsCmd)
}

func fetchMergeRequests() {
	color.Yellow("Fetching merge requests…")

	o := &gitlab.MergeRequestsOptions{}
	o.Page = page
	o.PerPage = perPage

	loader.Start()
	collection, meta, err := client.MergeRequests(o)
	loader.Stop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(collection.Items) == 0 {
		color.Red("No merge request found")
	} else {
		out.MergeRequests(output, outputFormat, collection)
	}

	printMeta(meta, true)

	handlePaginatedResult(meta, fetchMergeRequests)
}

var listMergeRequestsCmd = &cobra.Command{
	Use:     "merge-requests",
	Aliases: []string{"mr"},
	Short:   "List merge requests",
	Run: func(cmd *cobra.Command, args []string) {
		fetchMergeRequests()
	},
}
