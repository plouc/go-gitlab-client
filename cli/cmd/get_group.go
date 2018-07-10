package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var groupWithCustomAttributes bool

func init() {
	getCmd.AddCommand(getGroupCmd)

	getGroupCmd.Flags().BoolVarP(&groupWithCustomAttributes, "with-custom-attributes", "x", false, "Include custom attributes (admins only)")
}

var getGroupCmd = &cobra.Command{
	Use:     resourceCmd("group", "group"),
	Aliases: []string{"g"},
	Short:   "Get all details of a group",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "group", args)
		if err != nil {
			return err
		}

		color.Yellow("Fetching group (id: %s)…", ids["group_id"])

		loader.Start()
		group, meta, err := client.Group(ids["group_id"], groupWithCustomAttributes)
		loader.Stop()
		if err != nil {
			return err
		}

		groupOutput(group)

		metaOutput(meta, false)

		return nil
	},
}
