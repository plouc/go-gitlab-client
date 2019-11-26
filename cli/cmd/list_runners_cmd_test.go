package cmd

import (
	"testing"

	"github.com/edison-moreland/go-gitlab-client/test"
)

func TestListRunnersCmd(t *testing.T) {
	test.RunCommandTestCases(t, "runners", []*test.CommandTestCase{
		{
			[]string{"list", "runners", "--help"},
			nil,
			//configs["default"],
			"list_runners_help",
			false,
			nil,
		},
		{
			[]string{"list", "runners"},
			nil,
			//configs["default"],
			"list_runners",
			false,
			nil,
		},
		/*
			{
				[]string{"list", "runners", "--verbose"},
				nil,
				//configs["default"],
				"list_runners_verbose",
				false,
				nil,
			},
		*/
		{
			[]string{"list", "runners", "-f", "json"},
			nil,
			//configs["default"],
			"list_runners_json",
			false,
			nil,
		},
		{
			[]string{"list", "runners", "-f", "yaml"},
			nil,
			//configs["default"],
			"list_runners_yaml",
			false,
			nil,
		},
	})
}
