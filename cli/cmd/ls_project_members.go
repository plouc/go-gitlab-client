package cmd

import (
	"fmt"
	"github.com/fatih/color"
	out "github.com/plouc/go-gitlab-client/cli/output"
	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
)

var projectMembersQuery string

func init() {
	lsCmd.AddCommand(lsProjectMembersCmd)

	lsProjectMembersCmd.Flags().StringVarP(&projectMembersQuery, "query", "q", "", "Search term")
}

func fetchProjectMembers(projectId string) {
	color.Yellow("Fetching project members (id: %s)…", projectId)

	o := &gitlab.MembersOptions{}
	o.Page = page
	o.PerPage = perPage
	if projectMembersQuery != "" {
		o.Query = projectMembersQuery
	}

	loader.Start()
	collection, meta, err := client.ProjectMembers(projectId, o)
	loader.Stop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(collection.Items) == 0 {
		color.Red("No member found in project")
	} else {
		out.Members(output, outputFormat, collection)
	}

	printMeta(meta, true)

	handlePaginatedResult(meta, func() {
		fetchProjectMembers(projectId)
	})
}

var lsProjectMembersCmd = &cobra.Command{
	Use:     resourceCmd("project-members", "project"),
	Aliases: []string{"pm"},
	Short:   "List project members",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project", args)
		if err != nil {
			return err
		}

		fetchProjectMembers(ids["project_id"])

		return nil
	},
}
