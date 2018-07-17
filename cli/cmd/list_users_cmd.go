package cmd

import (
	"fmt"

	"github.com/fatih/color"
	out "github.com/plouc/go-gitlab-client/cli/output"
	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
)

var usersSearch, usersUsername string
var usersActive, usersBlocked bool

func init() {
	listCmd.AddCommand(listUsersCmd)

	listUsersCmd.Flags().StringVarP(&usersSearch, "search", "s", "", "Search users by email or username")
	listUsersCmd.Flags().StringVarP(&usersUsername, "username", "u", "", "Search users by username")
	listUsersCmd.Flags().BoolVar(&usersActive, "active", false, "Limit to active users")
	listUsersCmd.Flags().BoolVar(&usersBlocked, "blocked", false, "Limit to blocked users")
}

func fetchUsers() {
	color.Yellow("Fetching users…")

	o := &gitlab.UsersOptions{}
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
	collection, meta, err := client.Users(o)
	loader.Stop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(collection.Items) == 0 {
		color.Red("  No user found")
	} else {
		out.Users(output, outputFormat, collection)
	}

	printMeta(meta, true)

	handlePaginatedResult(meta, fetchUsers)
}

var listUsersCmd = &cobra.Command{
	Use:     "users",
	Aliases: []string{"u"},
	Short:   "List users",
	Run: func(cmd *cobra.Command, args []string) {
		fetchUsers()
	},
}
