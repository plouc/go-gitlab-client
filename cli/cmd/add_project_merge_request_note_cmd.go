package cmd

import (
	"strconv"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	out "github.com/plouc/go-gitlab-client/cli/output"
	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
)

func init() {
	addCmd.AddCommand(addProjectMergeRequestNoteCmd)
}

var addProjectMergeRequestNoteCmd = &cobra.Command{
	Use:     resourceCmd("project-merge-request-note", "project-merge-request"),
	Aliases: []string{"merge-request-note", "mr-note"},
	Short:   "Add project issue note",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project-merge-request", args)
		if err != nil {
			return err
		}

		mergeRequestIid, err := strconv.Atoi(ids["merge_request_iid"])
		if err != nil {
			return err
		}

		color.Yellow("Creating note for project merge request (project id: %s, merge request iid: %d)…", ids["project_id"], mergeRequestIid)

		note := new(gitlab.NoteAddPayload)

		prompt := promptui.Prompt{
			Label: "Body",
		}
		body, err := prompt.Run()
		if err != nil {
			return err
		}
		note.Body = body

		loader.Start()
		createdNote, meta, err := client.AddProjectMergeRequestNote(ids["project_id"], mergeRequestIid, note)
		loader.Stop()
		if err != nil {
			return err
		}

		out.Note(output, outputFormat, createdNote)

		printMeta(meta, false)

		return nil
	},
}
