package cmd

import (
	"testing"

	"github.com/edison-moreland/go-gitlab-client/test"
)

func TestListProjectsCmd(t *testing.T) {
	test.RunCommandTestCases(t, "projects", []*test.CommandTestCase{
		{
			[]string{"list", "projects", "--help"},
			nil,
			//configs["default"],
			"list_projects_help",
			false,
			nil,
		},
		{
			[]string{"list", "projects"},
			nil,
			//configs["default"],
			"list_projects",
			false,
			nil,
		},
		/*
			{
				[]string{"list", "projects", "--verbose"},
				nil,
				//configs["default"],
				"list_projects_verbose",
				false,
				nil,
			},
		*/
		{
			[]string{"list", "projects", "-f", "json"},
			nil,
			//configs["default"],
			"list_projects_json",
			false,
			nil,
		},
		{
			[]string{"list", "projects", "-f", "yaml"},
			nil,
			//configs["default"],
			"list_projects_yaml",
			false,
			nil,
		},
	})
}
