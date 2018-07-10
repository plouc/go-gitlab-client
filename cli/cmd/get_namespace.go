package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	getCmd.AddCommand(getNamespaceCmd)
}

var getNamespaceCmd = &cobra.Command{
	Use:     resourceCmd("namespace", "namespace"),
	Aliases: []string{"ns"},
	Short:   "Get a single namespace",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "namespace", args)
		if err != nil {
			return err
		}

		color.Yellow("Fetching namespace (id: %s)…", ids["namespace_id"])

		loader.Start()
		namespace, meta, err := client.Namespace(ids["namespace_id"])
		loader.Stop()
		if err != nil {
			return err
		}

		namespaceOutput(namespace)

		metaOutput(meta, false)

		return nil
	},
}
