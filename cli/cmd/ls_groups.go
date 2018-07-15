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
	lsCmd.AddCommand(lsGroupsCmd)

	lsGroupsCmd.Flags().StringVarP(&groupsSearch, "search", "s", "", "Return the list of authorized groups matching the search criteria")
	lsGroupsCmd.Flags().BoolVar(&groupsAllAvailable, "all", false, "Show all the groups you have access to (defaults to false for authenticated users, true for admin)")
	lsGroupsCmd.Flags().BoolVar(&groupsStatistics, "statistics", false, "Include group statistics (admins only)")
	lsGroupsCmd.Flags().BoolVarP(&groupsWithCustomAttributes, "with-custom-attributes", "x", false, "Include custom attributes in response (admins only)")
	lsGroupsCmd.Flags().BoolVar(&groupsOwned, "owned", false, "Limit to groups owned by the current user")
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
	groups, meta, err := client.Groups(o)
	loader.Stop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(groups) == 0 {
		color.Red("No group found")
	} else {
		out.Groups(output, outputFormat, groups)
	}

	out.Meta(meta, true)

	handlePaginatedResult(meta, fetchGroups)
}

var lsGroupsCmd = &cobra.Command{
	Use:     "groups",
	Aliases: []string{"g"},
	Short:   "List groups",
	Run: func(cmd *cobra.Command, args []string) {
		fetchGroups()
	},
}
