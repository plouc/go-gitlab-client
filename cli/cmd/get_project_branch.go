package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	getCmd.AddCommand(getProjectBranchCmd)
}

var getProjectBranchCmd = &cobra.Command{
	Use:     resourceCmd("project-branch", "project-branch"),
	Aliases: []string{"pb"},
	Short:   "Get project branch info",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project-branch", args)
		if err != nil {
			return err
		}

		color.Yellow("Fetching project's branch (project id: %s, branch name: %s)…", ids["project_id"], ids["branch_name"])

		loader.Start()
		branch, meta, err := client.ProjectBranch(ids["project_id"], ids["branch_name"])
		loader.Stop()
		if err != nil {
			return err
		}

		branchOutput(branch)

		metaOutput(meta, false)

		return nil
	},
}
