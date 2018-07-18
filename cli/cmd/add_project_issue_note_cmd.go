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
	addCmd.AddCommand(addProjectIssueNoteCmd)
}

var addProjectIssueNoteCmd = &cobra.Command{
	Use:     resourceCmd("project-issue-note", "project-issue"),
	Aliases: []string{"issue-note"},
	Short:   "Add project issue note",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project-issue", args)
		if err != nil {
			return err
		}

		issueIid, err := strconv.Atoi(ids["issue_iid"])
		if err != nil {
			return err
		}

		color.Yellow("Creating note for project issue (project id: %s, issue iid: %d)…", ids["project_id"], issueIid)

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
		createdNote, meta, err := client.AddProjectIssueNote(ids["project_id"], issueIid, note)
		loader.Stop()
		if err != nil {
			return err
		}

		out.Note(output, outputFormat, createdNote)

		printMeta(meta, false)

		return nil
	},
}
