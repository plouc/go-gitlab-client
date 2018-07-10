package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	addCmd.AddCommand(addProjectProtectedBranchCmd)
}

var addProjectProtectedBranchCmd = &cobra.Command{
	Use:     resourceCmd("project-protected-branch", "project-branch"),
	Aliases: []string{"ppb"},
	Short:   "Protect project branch",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project-branch", args)
		if err != nil {
			return err
		}

		color.Yellow("Protecting project branch (project id: %s, branch name: %s)…", ids["project_id"], ids["branch_name"])

		loader.Start()
		meta, err := client.ProtectBranch(ids["project_id"], ids["branch_name"])
		loader.Stop()
		if err != nil {
			return err
		}

		metaOutput(meta, false)

		return nil
	},
}
