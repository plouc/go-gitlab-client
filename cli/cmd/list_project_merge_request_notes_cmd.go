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
	listCmd.AddCommand(listProjectMergeRequestNotesCmd)
}

func fetchProjectMergeRequestNotes(projectId string, mergeRequestIid int) {
	color.Yellow("Fetching project merge request notes (project id: %s, merge request iid: %d)…", projectId, mergeRequestIid)

	o := &gitlab.NotesOptions{}
	o.Page = page
	o.PerPage = perPage

	loader.Start()
	collection, meta, err := client.ProjectMergeRequestNotes(projectId, mergeRequestIid, o)
	loader.Stop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(collection.Items) == 0 {
		color.Red("No note found for project %s merge request %d", projectId, mergeRequestIid)
	} else {
		out.Notes(output, outputFormat, collection)
	}

	printMeta(meta, true)

	handlePaginatedResult(meta, func() {
		fetchProjectMergeRequestNotes(projectId, mergeRequestIid)
	})
}

var listProjectMergeRequestNotesCmd = &cobra.Command{
	Use:     resourceCmd("project-merge-request-notes", "project-merge-request"),
	Aliases: []string{"project-mr-notes", "merge-request-notes", "mr-notes"},
	Short:   "List project merge request notes",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project-merge-request", args)
		if err != nil {
			return err
		}

		mergeRequestIid, err := strconv.Atoi(ids["merge_request_iid"])
		if err != nil {
			return err
		}

		fetchProjectMergeRequestNotes(ids["project_id"], mergeRequestIid)

		return nil
	},
}
