package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	getCmd.AddCommand(getProjectVarCmd)
}

var getProjectVarCmd = &cobra.Command{
	Use:     "project-var [project id] [var key]",
	Aliases: []string{"pv"},
	Short:   "Get the details of a project's specific variable",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return fmt.Errorf("you must specify a project id and variable key")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		projectId := args[0]
		varKey := args[1]

		color.Yellow("Fetching project variable (id: %s, key: %s)…", projectId, varKey)

		loader.Start()
		variable, meta, err := client.ProjectVariable(projectId, varKey)
		loader.Stop()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		varOutput(variable)

		metaOutput(meta, false)
	},
}
