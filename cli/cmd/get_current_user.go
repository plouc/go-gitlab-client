package cmd

import (
	"fmt"

	"github.com/fatih/color"
	out "github.com/plouc/go-gitlab-client/cli/output"
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

		out.User(output, outputFormat, user)

		printMeta(meta, false)
	},
}
