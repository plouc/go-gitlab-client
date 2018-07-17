package cmd

import (
	"testing"

	"github.com/plouc/go-gitlab-client/test"
)

func TestGetProjectMergeRequestCmd(t *testing.T) {
	test.RunCommandTestCases(t, "merge_requests", []*test.CommandTestCase{
		/*
			{
				[]string{"get", "project-merge-request", "-h"},
				nil,
				//configs["default"],
				"get_project_merge_request_help",
				false,
				nil,
			},
			{
				[]string{"get", "project-mr"},
				nil,
				//configs["default"],
				"get_project_merge_request_no_arg",
				true,
				nil,
			},
			{
				[]string{"get", "project-mr", "1"},
				nil,
				//configs["default"],
				"get_project_merge_request_no_merge_request_iid",
				true,
				nil,
			},
		*/
		{
			[]string{"get", "project-mr", "1", "1"},
			nil,
			//configs["default"],
			"get_project_merge_request",
			false,
			nil,
		},
		/*
			{
				[]string{"get", "project-mr", "1", "1", "-v"},
				nil,
				//configs["default"],
				"get_project_merge_request_verbose",
				false,
				nil,
			},
		*/
		{
			[]string{"get", "project-mr", "1", "1", "-f", "json"},
			nil,
			//configs["default"],
			"get_project_merge_request_json",
			false,
			nil,
		},
		{
			[]string{"get", "project-mr", "1", "1", "-f", "yaml"},
			nil,
			//configs["default"],
			"get_project_merge_request_yaml",
			false,
			nil,
		},
	})
}
