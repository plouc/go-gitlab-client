package cmd

import (
	"testing"

	"github.com/edison-moreland/go-gitlab-client/test"
)

func TestListUserSshKeysCmd(t *testing.T) {
	test.RunCommandTestCases(t, "ssh_keys", []*test.CommandTestCase{
		{
			[]string{"list", "user-ssh-keys", "--help"},
			nil,
			//configs["default"],
			"list_user_ssh_keys_help",
			false,
			nil,
		},
		{
			[]string{"list", "user-ssh-keys", "1"},
			nil,
			//configs["default"],
			"list_user_ssh_keys",
			false,
			nil,
		},
		/*
			{
				[]string{"ls", "user-ssh-keys", "1", "-v"},
				nil,
				//configs["default"],
				"ls_user_ssh_keys_verbose",
				false,
				nil,
			},
		*/
		{
			[]string{"list", "user-ssh-keys", "1", "-f", "json"},
			nil,
			//configs["default"],
			"list_user_ssh_keys_json",
			false,
			nil,
		},
		{
			[]string{"list", "user-ssh-keys", "1", "-f", "yaml"},
			nil,
			//configs["default"],
			"list_user_ssh_keys_yaml",
			false,
			nil,
		},
	})
}
