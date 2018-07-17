package cmd

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
	out "github.com/plouc/go-gitlab-client/cli/output"
	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
)

func init() {
	listCmd.AddCommand(listUserSshKeysCmd)
}

func fetchUserSshKeys(userId int) {
	color.Yellow("Fetching user %d ssh keys…", userId)

	o := &gitlab.PaginationOptions{}
	o.Page = page
	o.PerPage = perPage

	loader.Start()
	collection, meta, err := client.CurrentUserSshKeys(o)
	loader.Stop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(collection.Items) == 0 {
		color.Red("No ssh key found for user: %d", userId)
	} else {
		out.SshKeys(output, outputFormat, collection)
	}

	printMeta(meta, true)

	handlePaginatedResult(meta, func() {
		fetchUserSshKeys(userId)
	})
}

var listUserSshKeysCmd = &cobra.Command{
	Use:     resourceCmd("user-ssh-keys", "user"),
	Aliases: []string{"usk"},
	Short:   "List specific user ssh keys",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "user", args)
		if err != nil {
			return err
		}

		userId, err := strconv.Atoi(ids["user_id"])
		if err != nil {
			return err
		}

		fetchUserSshKeys(userId)

		return nil
	},
}
