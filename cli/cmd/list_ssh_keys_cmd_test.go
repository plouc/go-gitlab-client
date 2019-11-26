package cmd

import (
	"testing"

	"github.com/edison-moreland/go-gitlab-client/test"
)

func TestListSshKeysCmd(t *testing.T) {
	test.RunCommandTestCases(t, "ssh_keys", []*test.CommandTestCase{
		{
			[]string{"list", "ssh-keys", "--help"},
			nil,
			//configs["default"],
			"list_ssh_keys_help",
			false,
			nil,
		},
		{
			[]string{"list", "ssh-keys"},
			nil,
			//configs["default"],
			"list_ssh_keys",
			false,
			nil,
		},
		/*
			{
				[]string{"list", "ssh-keys", "--verbose"},
				nil,
				//configs["default"],
				"list_ssh_keys_verbose",
				false,
				nil,
			},
		*/
		{
			[]string{"list", "ssh-keys", "-f", "json"},
			nil,
			//configs["default"],
			"list_ssh_keys_json",
			false,
			nil,
		},
		{
			[]string{"list", "ssh-keys", "-f", "yaml"},
			nil,
			//configs["default"],
			"list_ssh_keys_yaml",
			false,
			nil,
		},
	})
}
