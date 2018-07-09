package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rmCmd.AddCommand(rmProjectMergedBranchesCmd)
}

var rmProjectMergedBranchesCmd = &cobra.Command{
	Use:     "project-merged-branches [project id]",
	Aliases: []string{"pmb"},
	Short:   "Remove project merged branches",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("you must specify a project id")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		projectId := args[0]

		color.Yellow("Removing project merged branch (project id: %s)…", projectId)

		confirmed := confirmAction(
			fmt.Sprintf("Are you sure you want to remove project %s merged branches?", projectId),
			"aborted merged branches removal",
			autoConfirmRemoval,
		)
		if !confirmed {
			return
		}

		loader.Start()
		message, meta, err := client.RemoveProjectMergedBranches(projectId)
		loader.Stop()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		color.Green("✔ Project merged branches were successfully removed: %s", message)

		metaOutput(meta, false)
	},
}
