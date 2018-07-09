package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/fatih/color"
	"github.com/plouc/go-gitlab-client/integration/utils"
)

var update = flag.Bool("update", true, "update golden files")

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
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if !testCase.config.Exists() {
				testCase.config.Write()
			}

			args := append(testCase.args, "-c", testCase.config.File)

			fmt.Printf("%s\n", args)

			c := exec.Command(binaryPath, args...)
			output, err := c.CombinedOutput()
			if (err != nil) != testCase.wantErr {
				t.Fatalf("%s\nexpected (err != nil) to be %v, but got %v. err: %v", output, testCase.wantErr, err != nil, err)
			}
			actual := string(output)

			snapshot := utils.NewSnapshotFile(t, testCase.snapshot, snapshotsDir)
			if *update {
				snapshot.Write(actual)
			}
			expected := snapshot.Load()

			if !reflect.DeepEqual(expected, actual) {
				t.Fatalf("diff:\n%v", utils.StringsDiff(expected, actual))
			}
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
