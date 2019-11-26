package cmd

import (
	"testing"

	"github.com/edison-moreland/go-gitlab-client/test"
)

func TestListUsersCmd(t *testing.T) {
	test.RunCommandTestCases(t, "users", []*test.CommandTestCase{
		{
			[]string{"list", "users", "--help"},
			nil,
			//configs["default"],
			"list_users_help",
			false,
			nil,
		},
		{
			[]string{"list", "users"},
			nil,
			//configs["default"],
			"list_users",
			false,
			nil,
		},
		/*
			{
				[]string{"list", "users", "--verbose"},
				nil,
				//configs["default"],
				"ls_users_verbose",
				false,
				nil,
			},
		*/
		{
			[]string{"list", "users", "-f", "json"},
			nil,
			//configs["default"],
			"list_users_json",
			false,
			nil,
		},
		{
			[]string{"list", "users", "-f", "yaml"},
			nil,
			//configs["default"],
			"list_users_yaml",
			false,
			nil,
		},
	})
}
