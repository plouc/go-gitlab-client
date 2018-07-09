package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	getCmd.AddCommand(getUserCmd)
}

var getUserCmd = &cobra.Command{
	Use:     "user [user id]",
	Aliases: []string{"u"},
	Short:   "Get a single user",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("you must specify a user id")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		userId := args[0]

		color.Yellow("Fetching user (id: %s)…", userId)

		loader.Start()
		user, meta, err := client.User(userId)
		loader.Stop()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		userOutput(user)

		metaOutput(meta, false)
	},
}
