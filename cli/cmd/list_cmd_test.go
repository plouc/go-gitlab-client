package cmd

import (
	"testing"

	"github.com/edison-moreland/go-gitlab-client/test"
)

func TestListCmd(t *testing.T) {
	test.RunCommandTestCases(t, "users", []*test.CommandTestCase{
		{
			[]string{"list"},
			nil,
			//configs["default"],
			"list_help",
			false,
			nil,
		},
		{
			[]string{"ls"},
			nil,
			//configs["default"],
			"list_help",
			false,
			nil,
		},
		{
			[]string{"list", "--help"},
			nil,
			//configs["default"],
			"list_help",
			false,
			nil,
		},
	})
}
