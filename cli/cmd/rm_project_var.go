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
	Use:     "project-var [project id] [var key]",
	Aliases: []string{"pv"},
	Short:   "Remove a project's variable",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return fmt.Errorf("you must specify a project id and variable key")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		projectId := args[0]
		varKey := args[1]

		color.Yellow("Removing project variable (id: %s, key: %s)…", projectId, varKey)

		confirmed := confirmAction(
			fmt.Sprintf("Are you sure you want to remove project %s variable %s?", projectId, varKey),
			"aborted project variable removal",
			autoConfirmRemoval,
		)
		if !confirmed {
			return
		}

		loader.Start()
		meta, err := client.RemoveProjectVariable(projectId, varKey)
		loader.Stop()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		color.Green("✔ Successfully removed variable: %s", varKey)

		metaOutput(meta, false)
	},
}
