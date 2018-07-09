package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rmCmd.AddCommand(rmProjectProtectedBranchCmd)
}

var rmProjectProtectedBranchCmd = &cobra.Command{
	Use:     "project-protected-branch [project id] [branch name]",
	Aliases: []string{"ppb"},
	Short:   "Unprotect project branch",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return fmt.Errorf("you must specify a project id and a branch name")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		projectId := args[0]
		branchName := args[1]

		color.Yellow("Unprotecting project branch (project id: %s, branch name: %s)…", projectId, branchName)

		confirmed := confirmAction(
			fmt.Sprintf("Are you sure you want to remove project %s protected branch %s?", projectId, branchName),
			"aborted project protected branch removal",
			autoConfirmRemoval,
		)
		if !confirmed {
			return
		}

		loader.Start()
		meta, err := client.UnprotectBranch(projectId, branchName)
		loader.Stop()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		metaOutput(meta, false)
	},
}
