package cmd

import (
	"fmt"

	"github.com/fatih/color"
	out "github.com/plouc/go-gitlab-client/cli/output"
	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
)

var projectsSearch string
var projectsArchived, projectsOwned, projectsMembership, projectsStarred bool

func init() {
	listCmd.AddCommand(listProjectsCmd)

	listProjectsCmd.Flags().StringVarP(&projectsSearch, "search", "s", "", "Search term")
	listProjectsCmd.Flags().BoolVar(&projectsArchived, "archived", false, "Limit by archived status")
	listProjectsCmd.Flags().BoolVar(&projectsMembership, "membership", false, "Limit by projects that the current user is a member of")
	listProjectsCmd.Flags().BoolVar(&projectsOwned, "owned", false, "Limit by projects owned by the current user")
	listProjectsCmd.Flags().BoolVar(&projectsStarred, "starred", false, "Limit by projects starred by the current user")
}

func fetchProjects() {
	color.Yellow("Fetching projects…")

	o := &gitlab.ProjectsOptions{}
	o.Page = page
	o.PerPage = perPage
	if projectsSearch != "" {
		o.Search = projectsSearch
	}
	if projectsArchived {
		o.Archived = true
	}
	if projectsMembership {
		o.Membership = true
	}
	if projectsOwned {
		o.Owned = true
	}
	if projectsStarred {
		o.Starred = true
	}

	loader.Start()
	collection, meta, err := client.Projects(o)
	loader.Stop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(collection.Items) == 0 {
		color.Red("  No project found")
	} else {
		out.Projects(output, outputFormat, collection)
	}

	printMeta(meta, true)

	handlePaginatedResult(meta, fetchProjects)
}

var listProjectsCmd = &cobra.Command{
	Use:     "projects",
	Aliases: []string{"p"},
	Short:   "List projects",
	Run: func(cmd *cobra.Command, args []string) {
		fetchProjects()
	},
}
