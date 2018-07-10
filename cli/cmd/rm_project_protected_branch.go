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
	Use:     resourceCmd("project-protected-branch", "project-branch"),
	Aliases: []string{"ppb"},
	Short:   "Unprotect project branch",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project-branch", args)
		if err != nil {
			return err
		}

		color.Yellow("Unprotecting project branch (project id: %s, branch name: %s)…", ids["project_id"], ids["branch_name"])

		confirmed := confirmAction(
			fmt.Sprintf("Are you sure you want to remove project %s protected branch %s?", ids["project_id"], ids["branch_name"]),
			"aborted project protected branch removal",
			autoConfirmRemoval,
		)
		if !confirmed {
			return nil
		}

		loader.Start()
		meta, err := client.UnprotectBranch(ids["project_id"], ids["branch_name"])
		loader.Stop()
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		metaOutput(meta, false)

		return nil
	},
}
