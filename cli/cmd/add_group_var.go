package cmd

import (
	"github.com/fatih/color"
	out "github.com/plouc/go-gitlab-client/cli/output"
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

		out.Variable(output, outputFormat, createdVariable)

		printMeta(meta, false)

		return nil
	},
}
