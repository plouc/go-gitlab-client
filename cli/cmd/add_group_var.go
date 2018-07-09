package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	addCmd.AddCommand(addGroupVarCmd)
}

var addGroupVarCmd = &cobra.Command{
	Use:     "group-var [group id]",
	Aliases: []string{"gv"},
	Short:   "Create a new group variable",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("you must specify a group id")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		groupId := args[0]

		color.Yellow("Creating variable for group (group id: %s)…", groupId)

		variable, err := promptVariable()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		loader.Start()
		createdVariable, meta, err := client.AddGroupVariable(groupId, variable)
		loader.Stop()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		varOutput(createdVariable)

		metaOutput(meta, false)
	},
}
