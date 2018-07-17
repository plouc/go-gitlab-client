package cmd

import (
	"testing"

	"github.com/plouc/go-gitlab-client/test"
)

func TestRmCmd(t *testing.T) {
	test.RunCommandTestCases(t, "users", []*test.CommandTestCase{
		{
			[]string{"rm"},
			nil,
			//configs["default"],
			"rm_help",
			false,
			nil,
		},
		{
			[]string{"rm", "--help"},
			nil,
			//configs["default"],
			"rm_help",
			false,
			nil,
		},
	})
}
