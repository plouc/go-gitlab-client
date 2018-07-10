package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	getCmd.AddCommand(getProjectVarCmd)
}

var getProjectVarCmd = &cobra.Command{
	Use:     resourceCmd("project-var", "project-var"),
	Aliases: []string{"pv"},
	Short:   "Get the details of a project's specific variable",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project-var", args)
		if err != nil {
			return err
		}

		color.Yellow("Fetching project variable (id: %s, key: %s)…", ids["project_id"], ids["var_key"])

		loader.Start()
		variable, meta, err := client.ProjectVariable(ids["project_id"], ids["var_key"])
		loader.Stop()
		if err != nil {
			return err
		}

		varOutput(variable)

		metaOutput(meta, false)

		return nil
	},
}
