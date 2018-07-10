package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
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
	members, meta, err := client.ProjectMembers(projectId, o)
	loader.Stop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("")
	if len(members) == 0 {
		color.Red("No member found in project")
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Id", "Username", "Name", "State", "Expires at", "Access level", "Web url"})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		for _, member := range members {
			table.Append([]string{
				fmt.Sprintf("%d", member.Id),
				member.Username,
				member.Name,
				member.State,
				member.ExpiresAt,
				fmt.Sprintf("%d", member.AccessLevel),
				member.WebUrl,
			})
		}
		table.Render()
	}
	fmt.Println("")

	metaOutput(meta, true)

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
