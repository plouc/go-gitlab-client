package cmd

import (
	"fmt"

	"github.com/fatih/color"
	out "github.com/plouc/go-gitlab-client/cli/output"
	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
)

var groupsSearch string
var groupsAllAvailable, groupsStatistics, groupsWithCustomAttributes, groupsOwned bool

func init() {
	listCmd.AddCommand(listGroupsCmd)

	listGroupsCmd.Flags().StringVarP(&groupsSearch, "search", "s", "", "Return the list of authorized groups matching the search criteria")
	listGroupsCmd.Flags().BoolVar(&groupsAllAvailable, "all", false, "Show all the groups you have access to (defaults to false for authenticated users, true for admin)")
	listGroupsCmd.Flags().BoolVar(&groupsStatistics, "statistics", false, "Include group statistics (admins only)")
	listGroupsCmd.Flags().BoolVarP(&groupsWithCustomAttributes, "with-custom-attributes", "x", false, "Include custom attributes in response (admins only)")
	listGroupsCmd.Flags().BoolVar(&groupsOwned, "owned", false, "Limit to groups owned by the current user")
}

func fetchGroups() {
	color.Yellow("Fetching groups…")

	o := &gitlab.GroupsOptions{}
	o.Page = page
	o.PerPage = perPage
	if groupsSearch != "" {
		o.Search = groupsSearch
	}
	if groupsAllAvailable {
		o.AllAvailable = true
	}
	if groupsStatistics {
		o.Statistics = true
	}
	if groupsWithCustomAttributes {
		o.WithCustomAttributes = true
	}
	if groupsOwned {
		o.Owned = true
	}

	loader.Start()
	collection, meta, err := client.Groups(o)
	loader.Stop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(collection.Items) == 0 {
		color.Red("No group found")
	} else {
		out.Groups(output, outputFormat, collection)
	}

	printMeta(meta, true)

	handlePaginatedResult(meta, fetchGroups)
}

var listGroupsCmd = &cobra.Command{
	Use:     "groups",
	Aliases: []string{"g"},
	Short:   "List groups",
	Run: func(cmd *cobra.Command, args []string) {
		fetchGroups()
	},
}
