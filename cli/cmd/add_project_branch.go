package cmd

import (
	"github.com/fatih/color"
	out "github.com/plouc/go-gitlab-client/cli/output"
	"github.com/spf13/cobra"
)

var branchName, ref string

func init() {
	addCmd.AddCommand(addProjectBranchCmd)

	addProjectBranchCmd.Flags().StringVarP(&branchName, "branch", "b", "", "Name of the branch")
	addProjectBranchCmd.Flags().StringVarP(&ref, "ref", "r", "", "	The branch name or commit SHA to create branch from")
}

var addProjectBranchCmd = &cobra.Command{
	Use:     resourceCmd("project-branch", "project"),
	Aliases: []string{"pb"},
	Short:   "Create project branch",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project", args)
		if err != nil {
			return err
		}

		color.Yellow("Creating project's branch %s from %s (project id: %s)…", branchName, ref, ids["project_id"])

		loader.Start()
		createdBranch, meta, err := client.AddProjectBranch(ids["project_id"], branchName, ref)
		loader.Stop()
		if err != nil {
			return err
		}

		out.Branch(output, outputFormat, createdBranch)

		out.Meta(meta, false)

		return nil
	},
}
