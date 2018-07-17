package cmd

import (
	"testing"

	"github.com/plouc/go-gitlab-client/test"
)

func TestListMergeRequestsCmd(t *testing.T) {
	test.RunCommandTestCases(t, "merge_requests", []*test.CommandTestCase{
		{
			[]string{"list", "merge-requests", "--help"},
			nil,
			//configs["default"],
			"list_merge_requests_help",
			false,
			nil,
		},
		{
			[]string{"list", "merge-requests"},
			nil,
			//configs["default"],
			"list_merge_requests",
			false,
			nil,
		},
		/*
			{
				[]string{"list", "merge-requests", "--verbose"},
				nil,
				//configs["default"],
				"list_merge_requests_verbose",
				false,
				nil,
			},
		*/
		{
			[]string{"list", "merge-requests", "-f", "json"},
			nil,
			//configs["default"],
			"list_merge_requests_json",
			false,
			nil,
		},
		{
			[]string{"list", "merge-requests", "-f", "yaml"},
			nil,
			//configs["default"],
			"list_merge_requests_yaml",
			false,
			nil,
		},
	})
}
