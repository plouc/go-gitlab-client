package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	addCmd.AddCommand(addProjectVarCmd)
}

var addProjectVarCmd = &cobra.Command{
	Use:     resourceCmd("project-var", "project"),
	Aliases: []string{"pv"},
	Short:   "Create a new project variable",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project", args)
		if err != nil {
			return err
		}

		color.Yellow("Creating variable for project (project id: %s)…", ids["project_id"])

		variable, err := promptVariable()
		if err != nil {
			return err
		}

		loader.Start()
		createdVariable, meta, err := client.AddProjectVariable(ids["project_id"], variable)
		loader.Stop()
		if err != nil {
			return err
		}

		varOutput(createdVariable)

		metaOutput(meta, false)

		return nil
	},
}
