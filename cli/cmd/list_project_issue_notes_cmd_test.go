package cmd

import (
	"testing"

	"github.com/edison-moreland/go-gitlab-client/test"
)

func TestListProjectIssueNotesCmd(t *testing.T) {
	test.RunCommandTestCases(t, "notes", []*test.CommandTestCase{
		{
			[]string{"list", "project-issue-notes", "--help"},
			nil,
			//configs["default"],
			"list_project_issue_notes_help",
			false,
			nil,
		},
		{
			[]string{"list", "project-issue-notes"},
			nil,
			//configs["default"],
			"list_project_issue_notes_no_project_id",
			true,
			nil,
		},
		{
			[]string{"list", "project-issue-notes", "1"},
			nil,
			//configs["default"],
			"list_project_issue_notes_no_issue_iid",
			true,
			nil,
		},
		{
			[]string{"list", "project-issue-notes", "1", "5"},
			nil,
			//configs["default"],
			"list_project_issue_notes",
			false,
			nil,
		},
		{
			[]string{"list", "project-issue-notes", "1", "5", "-f", "json"},
			nil,
			//configs["default"],
			"list_project_issue_notes_json",
			false,
			nil,
		},
		{
			[]string{"list", "project-issue-notes", "1", "5", "-f", "yaml"},
			nil,
			//configs["default"],
			"list_project_issue_notes_yaml",
			false,
			nil,
		},
	})
}
