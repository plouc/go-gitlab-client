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
	addCmd.AddCommand(addGroupEpicNoteCmd)
}

var addGroupEpicNoteCmd = &cobra.Command{
	Use:     resourceCmd("group-epic-note", "group-epic"),
	Aliases: []string{"epic-note"},
	Short:   "Add group epic note",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "group-epic", args)
		if err != nil {
			return err
		}

		epicId, err := strconv.Atoi(ids["epic_id"])
		if err != nil {
			return err
		}

		color.Yellow("Creating note for group epic (group id: %s, epic id: %d)…", ids["project_id"], epicId)

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
		createdNote, meta, err := client.AddGroupEpicNote(ids["project_id"], epicId, note)
		loader.Stop()
		if err != nil {
			return err
		}

		out.Note(output, outputFormat, createdNote)

		printMeta(meta, false)

		return nil
	},
}
