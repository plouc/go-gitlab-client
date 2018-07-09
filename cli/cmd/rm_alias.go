package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rmCmd.AddCommand(rmAliasCmd)
}

var rmAliasCmd = &cobra.Command{
	Use:     "alias [alias] [resource type]",
	Aliases: []string{"a"},
	Short:   "Remove resource alias",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return fmt.Errorf("you must specify an alias and a resource type")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		alias := args[0]
		resourceType := args[1]

		color.Yellow("Removing %s alias %s…", resourceType, alias)

		_, a := config.findAlias(alias, resourceType)
		if a == nil {
			color.Red("✘ No %s alias found for %s", resourceType, alias)
			return
		}

		config.Write(configFile)

		color.Green("✔ Alias %s was successfully removed", alias)
	},
}
