package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/fatih/color"
	"github.com/plouc/go-gitlab-client/integration/utils"
	"github.com/plouc/gosnap"
)

const cliPath = "cli"
const binaryName = "glc"

var configDir string
var snapshotsDir string
var binaryPath string

var configs = map[string]*utils.Config{
	"default": utils.NewConfig(configDir, ".glc.test.yml"),
}

type TestCase struct {
	name     string
	args     []string
	config   *utils.Config
	snapshot string
	wantErr  bool
}

func TestCLI(t *testing.T) {
	testCases := []*TestCase{
		{
			"no arg",
			[]string{},
			configs["default"],
			"help",
			false,
		},
		{
			"help",
			[]string{},
			configs["default"],
			"help",
			false,
		},
		{
			"ls no arg",
			[]string{"ls"},
			configs["default"],
			"ls_help",
			false,
		},
		{
			"get no arg",
			[]string{"get"},
			configs["default"],
			"get_help",
			false,
		},
		{
			"add no arg",
			[]string{"add"},
			configs["default"],
			"add_help",
			false,
		},
		{
			"rm no arg",
			[]string{"rm"},
			configs["default"],
			"rm_help",
			false,
		},
		{
			"ls groups help",
			[]string{"ls", "groups", "-h"},
			configs["default"],
			"ls_groups_help",
			false,
		},
		{
			"ls groups",
			[]string{"ls", "groups"},
			configs["default"],
			"ls_groups",
			false,
		},
		{
			"ls groups verbose",
			[]string{"ls", "groups", "-v"},
			configs["default"],
			"ls_groups_verbose",
			false,
		},
		{
			"ls groups json",
			[]string{"ls", "groups", "-f", "json"},
			configs["default"],
			"ls_groups_json",
			false,
		},
		{
			"ls groups yaml",
			[]string{"ls", "groups", "-f", "yaml"},
			configs["default"],
			"ls_groups_yaml",
			false,
		},
		{
			"ls users help",
			[]string{"ls", "users", "-h"},
			configs["default"],
			"ls_users_help",
			false,
		},
		{
			"ls users",
			[]string{"ls", "users"},
			configs["default"],
			"ls_users",
			false,
		},
		{
			"ls users verbose",
			[]string{"ls", "users", "-v"},
			configs["default"],
			"ls_users_verbose",
			false,
		},
		{
			"ls users json",
			[]string{"ls", "users", "-f", "json"},
			configs["default"],
			"ls_users_json",
			false,
		},
		{
			"ls users yaml",
			[]string{"ls", "users", "-f", "yaml"},
			configs["default"],
			"ls_users_yaml",
			false,
		},
		{
			"ls runners help",
			[]string{"ls", "runners", "-h"},
			configs["default"],
			"ls_runners_help",
			false,
		},
		{
			"ls runners",
			[]string{"ls", "runners"},
			configs["default"],
			"ls_runners",
			false,
		},
		{
			"ls runners verbose",
			[]string{"ls", "runners", "-v"},
			configs["default"],
			"ls_runners_verbose",
			false,
		},
		{
			"ls runners json",
			[]string{"ls", "runners", "-f", "json"},
			configs["default"],
			"ls_runners_json",
			false,
		},
		{
			"ls runners yaml",
			[]string{"ls", "runners", "-f", "yaml"},
			configs["default"],
			"ls_runners_yaml",
			false,
		},
		{
			"ls projects help",
			[]string{"ls", "projects", "-h"},
			configs["default"],
			"ls_projects_help",
			false,
		},
		{
			"ls projects",
			[]string{"ls", "projects"},
			configs["default"],
			"ls_projects",
			false,
		},
		{
			"ls projects verbose",
			[]string{"ls", "projects", "-v"},
			configs["default"],
			"ls_projects_verbose",
			false,
		},
		{
			"ls projects json",
			[]string{"ls", "projects", "-f", "json"},
			configs["default"],
			"ls_projects_json",
			false,
		},
		{
			"ls projects yaml",
			[]string{"ls", "projects", "-f", "yaml"},
			configs["default"],
			"ls_projects_yaml",
			false,
		},
		{
			"ls namespaces help",
			[]string{"ls", "namespaces", "-h"},
			configs["default"],
			"ls_namespaces_help",
			false,
		},
		{
			"ls namespaces",
			[]string{"ls", "namespaces"},
			configs["default"],
			"ls_namespaces",
			false,
		},
		{
			"ls namespaces search",
			[]string{"ls", "namespaces", "-s", "twitter"},
			configs["default"],
			"ls_namespaces_search",
			false,
		},
		{
			"ls namespaces verbose",
			[]string{"ls", "namespaces", "-v"},
			configs["default"],
			"ls_namespaces_verbose",
			false,
		},
		{
			"ls namespaces json",
			[]string{"ls", "namespaces", "-f", "json"},
			configs["default"],
			"ls_namespaces_json",
			false,
		},
		{
			"ls namespaces yaml",
			[]string{"ls", "namespaces", "-f", "yaml"},
			configs["default"],
			"ls_namespaces_yaml",
			false,
		},
		{
			"get user no arg",
			[]string{"get", "user"},
			configs["default"],
			"get_user_no_arg",
			true,
		},
		{
			"get user help",
			[]string{"get", "user", "-h"},
			configs["default"],
			"get_user_help",
			false,
		},
		{
			"get user",
			[]string{"get", "user", "1"},
			configs["default"],
			"get_user",
			false,
		},
		{
			"get user verbose",
			[]string{"get", "user", "1", "-v"},
			configs["default"],
			"get_user_verbose",
			false,
		},
		{
			"get user json",
			[]string{"get", "user", "1", "-f", "json"},
			configs["default"],
			"get_user_json",
			false,
		},
		{
			"get user yaml",
			[]string{"get", "user", "1", "-f", "yaml"},
			configs["default"],
			"get_user_yaml",
			false,
		},

		{
			"ls ssh keys help",
			[]string{"ls", "ssh-keys", "-h"},
			configs["default"],
			"ls_ssh_keys_help",
			false,
		},
		{
			"ls ssh keys",
			[]string{"ls", "ssh-keys"},
			configs["default"],
			"ls_ssh_keys",
			false,
		},
		{
			"ls ssh keys verbose",
			[]string{"ls", "ssh-keys", "-v"},
			configs["default"],
			"ls_ssh_keys_verbose",
			false,
		},
		{
			"ls ssh keys json",
			[]string{"ls", "ssh-keys", "-f", "json"},
			configs["default"],
			"ls_ssh_keys_json",
			false,
		},
		{
			"ls ssh keys yaml",
			[]string{"ls", "ssh-keys", "-f", "yaml"},
			configs["default"],
			"ls_ssh_keys_yaml",
			false,
		},
		{
			"ls user ssh keys help",
			[]string{"ls", "user-ssh-keys", "-h"},
			configs["default"],
			"ls_user_ssh_keys_help",
			false,
		},
		{
			"ls user ssh keys",
			[]string{"ls", "user-ssh-keys", "1"},
			configs["default"],
			"ls_user_ssh_keys",
			false,
		},
		{
			"ls user ssh keys verbose",
			[]string{"ls", "user-ssh-keys", "1", "-v"},
			configs["default"],
			"ls_user_ssh_keys_verbose",
			false,
		},
		{
			"ls user ssh keys json",
			[]string{"ls", "user-ssh-keys", "1", "-f", "json"},
			configs["default"],
			"ls_user_ssh_keys_json",
			false,
		},
		{
			"ls user ssh keys yaml",
			[]string{"ls", "user-ssh-keys", "1", "-f", "yaml"},
			configs["default"],
			"ls_user_ssh_keys_yaml",
			false,
		},
		{
			"ls merge requests help",
			[]string{"ls", "merge-requests", "-h"},
			configs["default"],
			"ls_merge_requests_help",
			false,
		},
		{
			"ls merge requests",
			[]string{"ls", "merge-requests"},
			configs["default"],
			"ls_merge_requests",
			false,
		},
		{
			"ls merge requests verbose",
			[]string{"ls", "merge-requests", "-v"},
			configs["default"],
			"ls_merge_requests_verbose",
			false,
		},
		{
			"ls merge requests json",
			[]string{"ls", "merge-requests", "-f", "json"},
			configs["default"],
			"ls_merge_requests_json",
			false,
		},
		{
			"ls merge requests yaml",
			[]string{"ls", "merge-requests", "-f", "yaml"},
			configs["default"],
			"ls_merge_requests_yaml",
			false,
		},
		{
			"ls project merge requests help",
			[]string{"ls", "project-merge-requests", "-h"},
			configs["default"],
			"ls_project_merge_requests_help",
			false,
		},
		{
			"ls project merge requests",
			[]string{"ls", "project-merge-requests", "1"},
			configs["default"],
			"ls_project_merge_requests",
			false,
		},
		{
			"ls project merge requests verbose",
			[]string{"ls", "project-merge-requests", "1", "-v"},
			configs["default"],
			"ls_project_merge_requests_verbose",
			false,
		},
		{
			"ls project merge requests json",
			[]string{"ls", "project-merge-requests", "1", "-f", "json"},
			configs["default"],
			"ls_project_merge_requests_json",
			false,
		},
		{
			"ls project merge requests yaml",
			[]string{"ls", "project-merge-requests", "1", "-f", "yaml"},
			configs["default"],
			"ls_project_merge_requests_yaml",
			false,
		},
		{
			"ls group merge requests help",
			[]string{"ls", "group-merge-requests", "-h"},
			configs["default"],
			"ls_group_merge_requests_help",
			false,
		},
		{
			"ls group merge requests",
			[]string{"ls", "group-merge-requests", "1"},
			configs["default"],
			"ls_group_merge_requests",
			false,
		},
		{
			"ls group merge requests verbose",
			[]string{"ls", "group-merge-requests", "1", "-v"},
			configs["default"],
			"ls_group_merge_requests_verbose",
			false,
		},
		{
			"ls group merge requests json",
			[]string{"ls", "group-merge-requests", "1", "-f", "json"},
			configs["default"],
			"ls_group_merge_requests_json",
			false,
		},
		{
			"ls group merge requests yaml",
			[]string{"ls", "group-merge-requests", "1", "-f", "yaml"},
			configs["default"],
			"ls_group_merge_requests_yaml",
			false,
		},
		{
			"ls project environments help",
			[]string{"ls", "project-environments", "-h"},
			configs["default"],
			"ls_project_environments_help",
			false,
		},
		{
			"ls project environments",
			[]string{"ls", "project-environments", "1"},
			configs["default"],
			"ls_project_environments",
			false,
		},
		{
			"ls project environments verbose",
			[]string{"ls", "project-environments", "1", "-v"},
			configs["default"],
			"ls_project_environments_verbose",
			false,
		},
		{
			"ls project environments json",
			[]string{"ls", "project-environments", "1", "-f", "json"},
			configs["default"],
			"ls_project_environments_json",
			false,
		},
		{
			"ls project environments yaml",
			[]string{"ls", "project-environments", "1", "-f", "yaml"},
			configs["default"],
			"ls_project_environments_yaml",
			false,
		},
		{
			"get project merge request help",
			[]string{"get", "project-merge-request", "-h"},
			configs["default"],
			"get_project_merge_request_help",
			false,
		},
		{
			"get project merge request no arg",
			[]string{"get", "project-mr"},
			configs["default"],
			"get_project_merge_request_no_arg",
			true,
		},
		{
			"get project merge request no merge request iid",
			[]string{"get", "project-mr", "1"},
			configs["default"],
			"get_project_merge_request_no_merge_request_iid",
			true,
		},
		{
			"get project merge request",
			[]string{"get", "project-mr", "1", "1"},
			configs["default"],
			"get_project_merge_request",
			false,
		},
		{
			"get project merge request verbose",
			[]string{"get", "project-mr", "1", "1", "-v"},
			configs["default"],
			"get_project_merge_request_verbose",
			false,
		},
		{
			"get project merge request json",
			[]string{"get", "project-mr", "1", "1", "-f", "json"},
			configs["default"],
			"get_project_merge_request_json",
			false,
		},
		{
			"get project merge request yaml",
			[]string{"get", "project-mr", "1", "1", "-f", "yaml"},
			configs["default"],
			"get_project_merge_request_yaml",
			false,
		},
	}

	ctx := gosnap.NewContext(t, snapshotsDir)

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if !testCase.config.Exists() {
				testCase.config.Write()
			}

			args := append(testCase.args, "-c", testCase.config.File)

			c := exec.Command(binaryPath, args...)
			output, err := c.CombinedOutput()
			if (err != nil) != testCase.wantErr {
				t.Fatalf("%s\nexpected (err != nil) to be %v, but got %v. err: %v", output, testCase.wantErr, err != nil, err)
			}
			actual := string(output)

			var snapshot *gosnap.Snapshot
			if !ctx.Has(testCase.snapshot) {
				snapshot = ctx.NewSnapshot(testCase.snapshot)
			} else {
				snapshot = ctx.Get(testCase.snapshot)
			}
			snapshot.AssertString(actual)
		})
	}
}

func TestMain(m *testing.M) {
	color.NoColor = false
	color.Yellow("Building CLI binary %s…", binaryName)

	err := os.Chdir("..")
	if err != nil {
		fmt.Printf("could not change dir: %v", err)
		os.Exit(1)
	}

	binaryPath, err = filepath.Abs(filepath.Join(cliPath, binaryName))
	if err != nil {
		fmt.Printf("could not get abs path for %s: %v", binaryName, err)
		os.Exit(1)
	}

	cmd := exec.Command("make", "_cli_build", "--no-print-directory")
	cmd.Start()
	if err := cmd.Wait(); err != nil {
		output, _ := cmd.CombinedOutput()
		fmt.Printf("could not make '%s' binary:\n%v\n%s", binaryName, err, output)
		os.Exit(1)
	}

	err = os.Chdir(cliPath)
	if err != nil {
		fmt.Printf("could not change dir: %v", err)
		os.Exit(1)
	}

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("problems recovering caller information")
		os.Exit(1)
	}
	snapshotsDir = filepath.Join(filepath.Dir(filename), "snapshots")

	configDir, err = filepath.Abs(cliPath)
	if err != nil {
		fmt.Printf("could not determine config dir: %v", err)
		os.Exit(1)
	}

	os.Exit(m.Run())
}
