package cmd

import (
	"testing"

	"github.com/edison-moreland/go-gitlab-client/test"
)

func TestListGroupsCmd(t *testing.T) {
	test.RunCommandTestCases(t, "groups", []*test.CommandTestCase{
		{
			[]string{"list", "groups", "--help"},
			nil,
			//configs["default"],
			"list_groups_help",
			false,
			nil,
		},
		{
			[]string{"list", "groups"},
			nil,
			//configs["default"],
			"list_groups",
			false,
			nil,
		},
		/*
			{
				[]string{"list", "groups", "--verbose"},
				nil,
				//configs["default"],
				"list_groups_verbose",
				false,
				nil,
			},
		*/
		{
			[]string{"list", "groups", "-f", "json"},
			nil,
			//configs["default"],
			"list_groups_json",
			false,
			nil,
		},
		{
			[]string{"list", "groups", "-f", "yaml"},
			nil,
			//configs["default"],
			"list_groups_yaml",
			false,
			nil,
		},
	})
}
