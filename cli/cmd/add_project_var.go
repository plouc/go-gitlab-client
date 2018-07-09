package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	addCmd.AddCommand(addProjectVarCmd)
}

var addProjectVarCmd = &cobra.Command{
	Use:     "project-var [project id]",
	Aliases: []string{"pv"},
	Short:   "Create a new project variable",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("you must specify a project id")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		projectId := args[0]

		color.Yellow("Creating variable for project (project id: %s)…", projectId)

		variable, err := promptVariable()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		loader.Start()
		createdVariable, meta, err := client.AddProjectVariable(projectId, variable)
		loader.Stop()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		varOutput(createdVariable)

		metaOutput(meta, false)
	},
}
