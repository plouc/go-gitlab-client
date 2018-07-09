package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var groupWithCustomAttributes bool

func init() {
	getCmd.AddCommand(getGroupCmd)

	getGroupCmd.Flags().BoolVarP(&groupWithCustomAttributes, "with-custom-attributes", "x", false, "Include custom attributes (admins only)")
}

var getGroupCmd = &cobra.Command{
	Use:     "group [group id]",
	Aliases: []string{"g"},
	Short:   "Get all details of a group",
	Args: func(cmd *cobra.Command, args []string) error {
		if currentAlias == "" && len(args) < 1 {
			return fmt.Errorf("you must specify a group id")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var groupId string
		if currentAlias != "" {
			_, a := config.findAliasE(currentAlias, "group")
			groupId = a.ResourceIds["id"]

		} else {
			groupId = args[0]
		}

		color.Yellow("Fetching group (id: %s)…", groupId)

		loader.Start()
		group, meta, err := client.Group(groupId, groupWithCustomAttributes)
		loader.Stop()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		groupOutput(group)

		metaOutput(meta, false)
	},
}
