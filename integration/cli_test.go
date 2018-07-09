package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/fatih/color"
	"github.com/sergi/go-diff/diffmatchpatch"
	"gopkg.in/src-d/go-git.v4/utils/diff"
)

var update = flag.Bool("update", true, "update golden files")

const cliPath = "cli"
const binaryName = "glc"

var binaryPath string

type testFile struct {
	t    *testing.T
	name string
	dir  string
}

func newFixture(t *testing.T, name string) *testFile {
	return &testFile{t: t, name: name, dir: "fixtures"}
}

func newSnapshotFile(t *testing.T, name string) *testFile {
	return &testFile{t: t, name: name + ".snapshot", dir: "snapshots"}
}

func (tf *testFile) path() string {
	tf.t.Helper()
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		tf.t.Fatal("problems recovering caller information")
	}

	return filepath.Join(filepath.Dir(filename), tf.dir, tf.name)
}

func (tf *testFile) write(content string) {
	tf.t.Helper()
	err := ioutil.WriteFile(tf.path(), []byte(content), 0644)
	if err != nil {
		tf.t.Fatalf("could not write %s: %v", tf.name, err)
	}
}

func (tf *testFile) asFile() *os.File {
	tf.t.Helper()
	file, err := os.Open(tf.path())
	if err != nil {
		tf.t.Fatalf("could not open %s: %v", tf.name, err)
	}
	return file
}

func (tf *testFile) load() string {
	tf.t.Helper()

	content, err := ioutil.ReadFile(tf.path())
	if err != nil {
		tf.t.Fatalf("could not read file %s: %v", tf.name, err)
	}

	return string(content)
}

func TestCLI(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		snapshot string
		wantErr  bool
	}{
		{
			"no arg",
			[]string{},
			"help",
			false,
		},
		{
			"help",
			[]string{},
			"help",
			false,
		},
		{
			"ls no arg",
			[]string{"ls"},
			"ls_help",
			false,
		},
		{
			"get no arg",
			[]string{"get"},
			"get_help",
			false,
		},
		{
			"add no arg",
			[]string{"add"},
			"add_help",
			false,
		},
		{
			"rm no arg",
			[]string{"rm"},
			"rm_help",
			false,
		},
		{
			"ls groups",
			[]string{"ls", "groups"},
			"ls_groups",
			false,
		},
		{
			"ls groups verbose",
			[]string{"ls", "groups", "-v"},
			"ls_groups_verbose",
			false,
		},
		{
			"ls groups json",
			[]string{"ls", "groups", "-f", "json"},
			"ls_groups_json",
			false,
		},
		{
			"ls groups yaml",
			[]string{"ls", "groups", "-f", "yaml"},
			"ls_groups_yaml",
			false,
		},
		{
			"ls users",
			[]string{"ls", "users"},
			"ls_users",
			false,
		},
		{
			"ls users verbose",
			[]string{"ls", "users", "-v"},
			"ls_users_verbose",
			false,
		},
		{
			"ls users json",
			[]string{"ls", "users", "-f", "json"},
			"ls_users_json",
			false,
		},
		{
			"ls users yaml",
			[]string{"ls", "users", "-f", "yaml"},
			"ls_users_yaml",
			false,
		},
		{
			"ls runners",
			[]string{"ls", "runners"},
			"ls_runners",
			false,
		},
		{
			"ls runners verbose",
			[]string{"ls", "runners", "-v"},
			"ls_runners_verbose",
			false,
		},
		{
			"ls runners json",
			[]string{"ls", "runners", "-f", "json"},
			"ls_runners_json",
			false,
		},
		{
			"ls runners yaml",
			[]string{"ls", "runners", "-f", "yaml"},
			"ls_runners_yaml",
			false,
		},
		{
			"ls projects",
			[]string{"ls", "projects"},
			"ls_projects",
			false,
		},
		{
			"ls projects verbose",
			[]string{"ls", "projects", "-v"},
			"ls_projects_verbose",
			false,
		},
		{
			"ls projects json",
			[]string{"ls", "projects", "-f", "json"},
			"ls_projects_json",
			false,
		},
		{
			"ls projects yaml",
			[]string{"ls", "projects", "-f", "yaml"},
			"ls_projects_yaml",
			false,
		},
		{
			"ls namespaces",
			[]string{"ls", "namespaces"},
			"ls_namespaces",
			false,
		},
		{
			"ls namespaces search",
			[]string{"ls", "namespaces", "-s", "twitter"},
			"ls_namespaces_search",
			false,
		},
		{
			"ls namespaces verbose",
			[]string{"ls", "namespaces", "-v"},
			"ls_namespaces_verbose",
			false,
		},
		{
			"ls namespaces json",
			[]string{"ls", "namespaces", "-f", "json"},
			"ls_namespaces_json",
			false,
		},
		{
			"ls namespaces yaml",
			[]string{"ls", "namespaces", "-f", "yaml"},
			"ls_namespaces_yaml",
			false,
		},
		{
			"get user",
			[]string{"get", "user", "1"},
			"get_user",
			false,
		},
		{
			"get user verbose",
			[]string{"get", "user", "1", "-v"},
			"get_user_verbose",
			false,
		},
		{
			"get user json",
			[]string{"get", "user", "1", "-f", "json"},
			"get_user_json",
			false,
		},
		{
			"get user yaml",
			[]string{"get", "user", "1", "-f", "yaml"},
			"get_user_yaml",
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command(binaryPath, tt.args...)
			output, err := cmd.CombinedOutput()
			if (err != nil) != tt.wantErr {
				t.Fatalf("%s\nexpected (err != nil) to be %v, but got %v. err: %v", output, tt.wantErr, err != nil, err)
			}
			actual := string(output)

			snapshot := newSnapshotFile(t, tt.snapshot)
			if *update {
				snapshot.write(actual)
			}
			expected := snapshot.load()

			if !reflect.DeepEqual(expected, actual) {
				ds := []string{}
				diffs := diff.Do(expected, actual)
				for _, d := range diffs {
					if d.Type != diffmatchpatch.DiffEqual {
						if d.Type == diffmatchpatch.DiffDelete {
							ds = append(ds, color.RedString("- %s", d.Text))
						} else {
							ds = append(ds, color.RedString("+ %s", d.Text))
						}
					}
				}

				t.Fatalf("diff:\n%v", strings.Join(ds, ""))
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

	os.Exit(m.Run())
}
