package cmd

import (
	"testing"

	"github.com/plouc/go-gitlab-client/test"
)

func TestListProjectSnippetNotesCmd(t *testing.T) {
	test.RunCommandTestCases(t, "notes", []*test.CommandTestCase{
		{
			[]string{"list", "project-snippet-notes", "--help"},
			nil,
			//configs["default"],
			"list_project_snippet_notes_help",
			false,
			nil,
		},
		{
			[]string{"list", "project-snippet-notes"},
			nil,
			//configs["default"],
			"list_project_snippet_notes_no_project_id",
			true,
			nil,
		},
		{
			[]string{"list", "project-snippet-notes", "1"},
			nil,
			//configs["default"],
			"list_project_snippet_notes_no_snippet_id",
			true,
			nil,
		},
		{
			[]string{"list", "project-snippet-notes", "1", "7"},
			nil,
			//configs["default"],
			"list_project_snippet_notes",
			false,
			nil,
		},
		{
			[]string{"list", "project-snippet-notes", "1", "7", "-f", "json"},
			nil,
			//configs["default"],
			"list_project_snippet_notes_json",
			false,
			nil,
		},
		{
			[]string{"list", "project-snippet-notes", "1", "7", "-f", "yaml"},
			nil,
			//configs["default"],
			"list_project_snippet_notes_yaml",
			false,
			nil,
		},
	})
}
