package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	getCmd.AddCommand(getNamespaceCmd)
}

var getNamespaceCmd = &cobra.Command{
	Use:     "namespace [group id]",
	Aliases: []string{"ns"},
	Short:   "Get a single namespace",
	Args: func(cmd *cobra.Command, args []string) error {
		if currentAlias == "" && len(args) < 1 {
			return fmt.Errorf("you must specify a namespace id")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var namespaceId string
		if currentAlias != "" {
			_, a := config.findAliasE(currentAlias, "namespace")
			namespaceId = a.ResourceIds["id"]

		} else {
			namespaceId = args[0]
		}

		color.Yellow("Fetching namespace (id: %s)…", namespaceId)

		loader.Start()
		namespace, meta, err := client.Namespace(namespaceId)
		loader.Stop()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		namespaceOutput(namespace)

		metaOutput(meta, false)
	},
}
