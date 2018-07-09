package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(lsCmd)

	lsCmd.PersistentFlags().IntVarP(&page, "page", "p", 1, "Page")
	lsCmd.PersistentFlags().IntVarP(&perPage, "per-page", "l", 10, "Items per page")
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List resource",
}
