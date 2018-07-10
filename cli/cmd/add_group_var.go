package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	addCmd.AddCommand(addGroupVarCmd)
}

var addGroupVarCmd = &cobra.Command{
	Use:     resourceCmd("group-var", "group"),
	Aliases: []string{"gv"},
	Short:   "Create a new group variable",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "group", args)
		if err != nil {
			return err
		}

		color.Yellow("Creating variable for group (group id: %s)…", ids["group_id"])

		variable, err := promptVariable()
		if err != nil {
			return err
		}

		loader.Start()
		createdVariable, meta, err := client.AddGroupVariable(ids["group_id"], variable)
		loader.Stop()
		if err != nil {
			return err
		}

		varOutput(createdVariable)

		metaOutput(meta, false)

		return nil
	},
}
