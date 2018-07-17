package cmd

import (
	"testing"

	"github.com/plouc/go-gitlab-client/test"
)

func TestListGroupMergeRequestsCmd(t *testing.T) {
	test.RunCommandTestCases(t, "merge_requests", []*test.CommandTestCase{
		{
			[]string{"list", "group-merge-requests", "-h"},
			nil,
			//configs["default"],
			"list_group_merge_requests_help",
			false,
			nil,
		},
		{
			[]string{"list", "group-merge-requests", "1"},
			nil,
			//configs["default"],
			"list_group_merge_requests",
			false,
			nil,
		},
		/*
			{
				[]string{"list", "group-merge-requests", "1", "-v"},
				nil,
				//configs["default"],
				"list_group_merge_requests_verbose",
				false,
				nil,
			},
		*/
		{
			[]string{"list", "group-merge-requests", "1", "-f", "json"},
			nil,
			//configs["default"],
			"list_group_merge_requests_json",
			false,
			nil,
		},
		{
			[]string{"list", "group-merge-requests", "1", "-f", "yaml"},
			nil,
			//configs["default"],
			"list_group_merge_requests_yaml",
			false,
			nil,
		},
	})
}
