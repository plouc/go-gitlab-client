package cmd

import (
	"testing"

	"github.com/plouc/go-gitlab-client/test"
)

func TestListProjectMergeRequestsCmd(t *testing.T) {
	test.RunCommandTestCases(t, "merge_requests", []*test.CommandTestCase{
		{
			[]string{"list", "project-merge-requests", "--help"},
			nil,
			//configs["default"],
			"list_project_merge_requests_help",
			false,
			nil,
		},
		{
			[]string{"list", "project-merge-requests", "1"},
			nil,
			//configs["default"],
			"list_project_merge_requests",
			false,
			nil,
		},
		/*
			{
				[]string{"list", "project-merge-requests", "1", "--verbose"},
				nil,
				//configs["default"],
				"list_project_merge_requests_verbose",
				false,
				nil,
			},
		*/
		{
			[]string{"list", "project-merge-requests", "1", "-f", "json"},
			nil,
			//configs["default"],
			"list_project_merge_requests_json",
			false,
			nil,
		},
		{
			[]string{"list", "project-merge-requests", "1", "-f", "yaml"},
			nil,
			//configs["default"],
			"list_project_merge_requests_yaml",
			false,
			nil,
		},
	})
}
