package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	getCmd.AddCommand(getCurrentCmd)
}

var getCurrentCmd = &cobra.Command{
	Use:     "current-user",
	Aliases: []string{"cu"},
	Short:   "Get current user",
	Run: func(cmd *cobra.Command, args []string) {
		color.Yellow("Fetching current user…")

		loader.Start()
		user, meta, err := client.CurrentUser()
		loader.Stop()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		userOutput(user)

		metaOutput(meta, false)
	},
}
