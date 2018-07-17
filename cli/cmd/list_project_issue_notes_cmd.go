package cmd

import (
	"fmt"

	"github.com/fatih/color"
	out "github.com/plouc/go-gitlab-client/cli/output"
	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
	"strconv"
)

func init() {
	listCmd.AddCommand(listProjectIssueNotesCmd)
}

func fetchProjectIssueNotes(projectId string, issueIid int) {
	color.Yellow("Fetching project issue notes (project id: %s, issue iid: %d)…", projectId, issueIid)

	o := &gitlab.NotesOptions{}
	o.Page = page
	o.PerPage = perPage

	loader.Start()
	collection, meta, err := client.ProjectIssueNotes(projectId, issueIid, o)
	loader.Stop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(collection.Items) == 0 {
		color.Red("No note found for project %s issue %d", projectId, issueIid)
	} else {
		out.Notes(output, outputFormat, collection)
	}

	printMeta(meta, true)

	handlePaginatedResult(meta, func() {
		fetchProjectIssueNotes(projectId, issueIid)
	})
}

var listProjectIssueNotesCmd = &cobra.Command{
	Use:     resourceCmd("project-issue-notes", "project-issue"),
	Aliases: []string{"issue-notes"},
	Short:   "List project issue notes",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project-issue", args)
		if err != nil {
			return err
		}

		issueIid, err := strconv.Atoi(ids["issue_iid"])
		if err != nil {
			return err
		}

		fetchProjectIssueNotes(ids["project_id"], issueIid)

		return nil
	},
}
