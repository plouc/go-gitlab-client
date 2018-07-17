package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Print the version number of glc",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("glc v0.1")
	},
}
