package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(listCmd)

	listCmd.PersistentFlags().IntVarP(&page, "page", "p", 1, "Page")
	listCmd.PersistentFlags().IntVarP(&perPage, "per-page", "l", 10, "Items per page")
}

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List resource",
}
