package cmd

import (
	"strconv"

	"github.com/fatih/color"
	out "github.com/plouc/go-gitlab-client/cli/output"
	"github.com/spf13/cobra"
)

func init() {
	getCmd.AddCommand(getProjectMergeRequestNoteCmd)
}

var getProjectMergeRequestNoteCmd = &cobra.Command{
	Use:     resourceCmd("project-merge-request-note", "project-merge-request-note"),
	Aliases: []string{"project-mr-note", "merge-request-note", "mr-note"},
	Short:   "Get project merge request note",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project-merge-request-note", args)
		if err != nil {
			return err
		}

		mergeRequestIid, err := strconv.Atoi(ids["merge_request_iid"])
		if err != nil {
			return err
		}

		noteId, err := strconv.Atoi(ids["note_id"])
		if err != nil {
			return err
		}

		color.Yellow(
			"Fetching project's merge request note (project id: %s, merge request iid: %d, note id: %d)…",
			ids["project_id"], mergeRequestIid, noteId,
		)

		loader.Start()
		note, meta, err := client.ProjectMergeRequestNote(ids["project_id"], mergeRequestIid, noteId)
		loader.Stop()
		if err != nil {
			return err
		}

		out.Note(output, outputFormat, note)

		printMeta(meta, false)

		return nil
	},
}
