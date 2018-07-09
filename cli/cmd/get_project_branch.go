package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	getCmd.AddCommand(getProjectBranchCmd)
}

var getProjectBranchCmd = &cobra.Command{
	Use:     "project-branch [project id] [branch name]",
	Aliases: []string{"pb"},
	Short:   "Get project branch info",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return fmt.Errorf("you must specify a project id and a branch name")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		projectId := args[0]
		branchName := args[1]

		color.Yellow("Fetching project's branch (project id: %s, branch name: %s)…", projectId, branchName)

		loader.Start()
		branch, meta, err := client.ProjectBranch(projectId, branchName)
		loader.Stop()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		branchOutput(branch)

		metaOutput(meta, false)
	},
}
