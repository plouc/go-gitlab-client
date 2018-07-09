package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	getCmd.AddCommand(getGroupVarCmd)
}

var getGroupVarCmd = &cobra.Command{
	Use:     "group-var [group id] [var key]",
	Aliases: []string{"gv"},
	Short:   "Get the details of a group's specific variable",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return fmt.Errorf("you must specify a group id and variable key")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		groupId := args[0]
		varKey := args[1]

		color.Yellow("Fetching group variable (group id: %s, key: %s)…", groupId, varKey)

		loader.Start()
		variable, meta, err := client.GroupVariable(groupId, varKey)
		loader.Stop()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		varOutput(variable)

		metaOutput(meta, false)
	},
}
