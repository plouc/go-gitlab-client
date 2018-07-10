package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rmCmd.AddCommand(rmProjectVarCmd)
}

var rmProjectVarCmd = &cobra.Command{
	Use:     resourceCmd("project-var", "project-var"),
	Aliases: []string{"pv"},
	Short:   "Remove a project's variable",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project-var", args)
		if err != nil {
			return err
		}

		color.Yellow("Removing project variable (id: %s, key: %s)…", ids["project_id"], ids["var_key"])

		confirmed := confirmAction(
			fmt.Sprintf("Are you sure you want to remove project %s variable %s?", ids["project_id"], ids["var_key"]),
			"aborted project variable removal",
			autoConfirmRemoval,
		)
		if !confirmed {
			return nil
		}

		loader.Start()
		meta, err := client.RemoveProjectVariable(ids["project_id"], ids["var_key"])
		loader.Stop()
		if err != nil {
			return err
		}

		color.Green("✔ Successfully removed variable: %s", ids["var_key"])

		metaOutput(meta, false)

		return nil
	},
}
