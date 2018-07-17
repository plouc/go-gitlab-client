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
	testCases := []*TestCase{}

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
