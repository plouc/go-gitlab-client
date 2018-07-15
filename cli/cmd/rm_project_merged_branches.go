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
	Use:     resourceCmd("project-merged-branches", "project"),
	Aliases: []string{"pmb"},
	Short:   "Remove project merged branches",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project", args)
		if err != nil {
			return err
		}

		color.Yellow("Removing project merged branch (project id: %s)…", ids["project_id"])

		confirmed := confirmAction(
			fmt.Sprintf("Are you sure you want to remove project %s merged branches?", ids["project_id"]),
			"aborted merged branches removal",
			autoConfirmRemoval,
		)
		if !confirmed {
			return nil
		}

		loader.Start()
		message, meta, err := client.RemoveProjectMergedBranches(ids["project_id"])
		loader.Stop()
		if err != nil {
			return err
		}

		color.Green("✔ Project merged branches were successfully removed: %s", message)

		printMeta(meta, false)

		return nil
	},
}
