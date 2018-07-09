package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/plouc/go-gitlab-client/gogitlab"
	"github.com/spf13/cobra"
)

var usersSearch, usersUsername string
var usersActive, usersBlocked bool

func init() {
	lsCmd.AddCommand(lsUsersCmd)

	lsUsersCmd.Flags().StringVarP(&usersSearch, "search", "s", "", "Search users by email or username")
	lsUsersCmd.Flags().StringVarP(&usersUsername, "username", "u", "", "Search users by username")
	lsUsersCmd.Flags().BoolVar(&usersActive, "active", false, "Limit to active users")
	lsUsersCmd.Flags().BoolVar(&usersBlocked, "blocked", false, "Limit to blocked users")
}

func fetchUsers() {
	color.Yellow("Fetching users…")

	o := &gogitlab.UsersOptions{}
	o.Page = page
	o.PerPage = perPage
	if usersSearch != "" {
		o.Search = usersSearch
	}
	if usersUsername != "" {
		o.Username = usersUsername
	}
	if usersActive {
		o.Active = true
	}
	if usersBlocked {
		o.Blocked = true
	}

	loader.Start()
	users, meta, err := client.Users(o)
	loader.Stop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(users) == 0 {
		color.Red("  No user found")
	} else {
		usersOutput(users)
	}

	metaOutput(meta, true)

	handlePaginatedResult(meta, fetchUsers)
}

var lsUsersCmd = &cobra.Command{
	Use:     "users",
	Aliases: []string{"u"},
	Short:   "List users",
	Run: func(cmd *cobra.Command, args []string) {
		fetchUsers()
	},
}
