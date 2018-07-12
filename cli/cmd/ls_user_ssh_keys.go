package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
	"strconv"
)

func init() {
	lsCmd.AddCommand(lsUserSshKeysCmd)
}

func fetchUserSshKeys(userId int) {
	color.Yellow("Fetching user %d ssh keys…", userId)

	o := &gitlab.PaginationOptions{}
	o.Page = page
	o.PerPage = perPage

	loader.Start()
	keys, meta, err := client.CurrentUserSshKeys(o)
	loader.Stop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(keys) == 0 {
		color.Red("No ssh key found for user: %d", userId)
	} else {
		sshKeysOutput(keys)
	}

	metaOutput(meta, true)

	handlePaginatedResult(meta, func() {
		fetchUserSshKeys(userId)
	})
}

var lsUserSshKeysCmd = &cobra.Command{
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
