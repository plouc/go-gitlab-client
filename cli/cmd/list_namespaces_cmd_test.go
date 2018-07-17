package cmd

import (
	"testing"

	"github.com/plouc/go-gitlab-client/test"
)

func TestListNamespacesCmd(t *testing.T) {
	test.RunCommandTestCases(t, "namespaces", []*test.CommandTestCase{
		{
			[]string{"list", "namespaces", "--help"},
			nil, //configs["default"],
			"list_namespaces_help",
			false,
			nil,
		},
		{
			[]string{"list", "namespaces"},
			nil, //configs["default"],
			"list_namespaces",
			false,
			nil,
		},
		{
			[]string{"list", "namespaces", "--search", "twitter"},
			nil, //configs["default"],
			"list_namespaces_search",
			false,
			nil,
		},
		{
			[]string{"list", "namespaces", "-s", "twitter"},
			nil, //configs["default"],
			"list_namespaces_search",
			false,
			nil,
		},
		/*
			{
				[]string{"list", "namespaces", "--verbose"},
				nil, //configs["default"],
				"list_namespaces_verbose",
				false,
				nil,
			},
		*/
		{
			[]string{"list", "namespaces", "-f", "json"},
			nil, //configs["default"],
			"list_namespaces_json",
			false,
			nil,
		},
		{
			[]string{"list", "namespaces", "-f", "yaml"},
			nil, //configs["default"],
			"list_namespaces_yaml",
			false,
			nil,
		},
	})
}
