package cmd

import (
	"github.com/spf13/cobra"
)

var autoConfirmRemoval bool

func init() {
	RootCmd.AddCommand(rmCmd)

	rmCmd.PersistentFlags().BoolVarP(&autoConfirmRemoval, "yes", "y", false, "Do not ask for confirmation")
}

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove resource",
}
