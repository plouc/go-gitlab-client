package cmd

import (
	"testing"

	"github.com/edison-moreland/go-gitlab-client/test"
)

func TestListProjectEnvironmentsCmd(t *testing.T) {
	test.RunCommandTestCases(t, "environments", []*test.CommandTestCase{
		{
			[]string{"list", "project-environments", "--help"},
			nil,
			//configs["default"],
			"list_project_environments_help",
			false,
			nil,
		},
		{
			[]string{"list", "project-environments", "1"},
			nil,
			//configs["default"],
			"list_project_environments",
			false,
			nil,
		},
		/*
			{
				[]string{"list", "project-environments", "1", "--verbose"},
				nil,
				//configs["default"],
				"list_project_environments_verbose",
				false,
				nil,
			},
		*/
		{
			[]string{"list", "project-environments", "1", "-f", "json"},
			nil,
			//configs["default"],
			"list_project_environments_json",
			false,
			nil,
		},
		{
			[]string{"list", "project-environments", "1", "-f", "yaml"},
			nil,
			//configs["default"],
			"list_project_environments_yaml",
			false,
			nil,
		},
	})
}
