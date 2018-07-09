package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	lsCmd.AddCommand(lsAliasesCmd)
}

var lsAliasesCmd = &cobra.Command{
	Use:     "aliases",
	Aliases: []string{"a"},
	Short:   "List resource aliases",
	Run: func(cmd *cobra.Command, args []string) {
		color.Yellow("Listing registered aliases…")

		if len(config.Aliases) == 0 {
			color.Red("  No alias found, create one with `glc add alias`")
		} else {
			aliasesOutput(config.Aliases)
		}
	},
}
