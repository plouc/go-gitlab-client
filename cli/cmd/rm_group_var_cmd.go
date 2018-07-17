package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rmCmd.AddCommand(rmGroupVarCmd)
}

var rmGroupVarCmd = &cobra.Command{
	Use:     resourceCmd("group-var", "group-var"),
	Aliases: []string{"gv"},
	Short:   "Remove a group's variable",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "group-var", args)
		if err != nil {
			return err
		}

		color.Yellow("Removing group variable (group id: %s, key: %s)…", ids["group_id"], ids["var_key"])

		confirmed := confirmAction(
			fmt.Sprintf("Are you sure you want to remove group %s variable %s?", ids["group_id"], ids["var_key"]),
			"aborted group variable removal",
			autoConfirmRemoval,
		)
		if !confirmed {
			return nil
		}

		loader.Start()
		meta, err := client.RemoveGroupVariable(ids["group_id"], ids["var_key"])
		loader.Stop()
		if err != nil {
			return err
		}

		color.Green("✔ Successfully removed variable: %s", ids["var_key"])

		printMeta(meta, false)

		return nil
	},
}
