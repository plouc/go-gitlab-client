package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var branchName, ref string

func init() {
	addCmd.AddCommand(addProjectBranchCmd)

	addProjectBranchCmd.Flags().StringVarP(&branchName, "branch", "b", "", "Name of the branch")
	addProjectBranchCmd.Flags().StringVarP(&ref, "ref", "r", "", "	The branch name or commit SHA to create branch from")
}

var addProjectBranchCmd = &cobra.Command{
	Use:     "project-branch [project id]",
	Aliases: []string{"pb"},
	Short:   "Create project branch",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("you must specify a project id")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		projectId := args[0]

		color.Yellow("Creating project's branch %s from %s (project id: %s)…", branchName, ref, projectId)

		loader.Start()
		createdBranch, meta, err := client.AddProjectBranch(projectId, branchName, ref)
		loader.Stop()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		branchOutput(createdBranch)

		metaOutput(meta, false)
	},
}
