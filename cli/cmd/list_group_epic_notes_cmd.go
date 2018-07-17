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
	listCmd.AddCommand(listGroupEpicNotesCmd)
}

func fetchGroupEpicNotes(groupId string, epicId int) {
	color.Yellow("Fetching group epic notes (project id: %s, epic id: %d)…", groupId, epicId)

	o := &gitlab.NotesOptions{}
	o.Page = page
	o.PerPage = perPage

	loader.Start()
	collection, meta, err := client.GroupEpicNotes(groupId, epicId, o)
	loader.Stop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(collection.Items) == 0 {
		color.Red("No note found for group %s epic %d", groupId, epicId)
	} else {
		out.Notes(output, outputFormat, collection)
	}

	printMeta(meta, true)

	handlePaginatedResult(meta, func() {
		fetchGroupEpicNotes(groupId, epicId)
	})
}

var listGroupEpicNotesCmd = &cobra.Command{
	Use:     resourceCmd("group-epic-notes", "group-epic"),
	Aliases: []string{"epic-notes"},
	Short:   "List group epic notes",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "group-epic", args)
		if err != nil {
			return err
		}

		epicId, err := strconv.Atoi(ids["epic_id"])
		if err != nil {
			return err
		}

		fetchGroupEpicNotes(ids["group_id"], epicId)

		return nil
	},
}
