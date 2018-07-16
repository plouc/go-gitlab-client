package cmd

import (
	"strconv"

	"github.com/fatih/color"
	out "github.com/plouc/go-gitlab-client/cli/output"
	"github.com/spf13/cobra"
)

func init() {
	getCmd.AddCommand(getProjectMergeRequestCmd)
}

var getProjectMergeRequestCmd = &cobra.Command{
	Use:     resourceCmd("project-merge-request", "project-merge-request"),
	Aliases: []string{"project-mr", "pmr"},
	Short:   "Get project merge request info",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project-merge-request", args)
		if err != nil {
			return err
		}

		mergeRequestIid, err := strconv.Atoi(ids["merge_request_iid"])
		if err != nil {
			return err
		}

		color.Yellow("Fetching project merge request (project id: %s, merge request iid: %d)…", ids["project_id"], mergeRequestIid)

		loader.Start()
		mergeRequest, meta, err := client.ProjectMergeRequest(ids["project_id"], mergeRequestIid)
		loader.Stop()
		if err != nil {
			return err
		}

		out.MergeRequest(output, outputFormat, mergeRequest)

		printMeta(meta, false)

		return nil
	},
}
