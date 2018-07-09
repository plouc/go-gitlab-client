package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rmCmd.AddCommand(rmProjectBranchCmd)
}

var rmProjectBranchCmd = &cobra.Command{
	Use:     "project-branch [group id] [branch name]",
	Aliases: []string{"pb"},
	Short:   "Remove project branch",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return fmt.Errorf("you must specify a project id and a branch name")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		projectId := args[0]
		branchName := args[1]

		color.Yellow("Removing project branch (project id: %s, branch name: %s)…", projectId, branchName)

		confirmed := confirmAction(
			fmt.Sprintf("Are you sure you want to remove project %s branch: %s?", projectId, branchName),
			"aborted branch removal",
			autoConfirmRemoval,
		)
		if !confirmed {
			return
		}

		loader.Start()
		meta, err := client.RemoveProjectBranch(projectId, branchName)
		loader.Stop()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		color.Green("✔ Branch %s was successfully removed", branchName)

		metaOutput(meta, false)
	},
}
