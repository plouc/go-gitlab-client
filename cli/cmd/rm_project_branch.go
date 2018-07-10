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
	Use:     resourceCmd("project-branch", "project-branch"),
	Aliases: []string{"pb"},
	Short:   "Remove project branch",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project-branch", args)
		if err != nil {
			return err
		}

		color.Yellow("Removing project branch (project id: %s, branch name: %s)…", ids["project_id"], ids["branch_name"])

		confirmed := confirmAction(
			fmt.Sprintf("Are you sure you want to remove project %s branch: %s?", ids["project_id"], ids["branch_name"]),
			"aborted branch removal",
			autoConfirmRemoval,
		)
		if !confirmed {
			return nil
		}

		loader.Start()
		meta, err := client.RemoveProjectBranch(ids["project_id"], ids["branch_name"])
		loader.Stop()
		if err != nil {
			return err
		}

		color.Green("✔ Branch %s was successfully removed", ids["branch_name"])

		metaOutput(meta, false)

		return nil
	},
}
