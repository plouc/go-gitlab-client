package cmd

import (
	"testing"

	"github.com/plouc/go-gitlab-client/test"
)

func TestGetUserCmd(t *testing.T) {
	test.RunCommandTestCases(t, "users", []*test.CommandTestCase{
		/*
			{
				[]string{"get", "user"},
				nil,
				//configs["default"],
				"get_user_no_arg",
				true,
				nil,
			},
			{
				[]string{"get", "user", "-h"},
				nil,
				//configs["default"],
				"get_user_help",
				false,
				nil,
			},
		*/
		{
			[]string{"get", "user", "1"},
			nil,
			//configs["default"],
			"get_user",
			false,
			nil,
		},
		/*
			{
				[]string{"get", "user", "1", "-v"},
				nil,
				//configs["default"],
				"get_user_verbose",
				false,
				nil,
			},
		*/
		{
			[]string{"get", "user", "1", "-f", "json"},
			nil,
			//configs["default"],
			"get_user_json",
			false,
			nil,
		},
		{
			[]string{"get", "user", "1", "-f", "yaml"},
			nil,
			//configs["default"],
			"get_user_yaml",
			false,
			nil,
		},
	})
}
