package test

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/plouc/gosnap"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

const cliPath = "cli"
const binaryName = "glc"

type CommandTestCase struct {
	Args []string
	Env  map[string]string
	//config   *CliConfig
	Snapshot string
	WantErr  bool
	Manual   func(t *testing.T, output string)
}

func BuildCli() {
	d := baseDir(nil)

	err := os.Chdir(filepath.Join(d, ".."))
	if err != nil {
		fmt.Printf("could not change dir: %v", err)
		os.Exit(1)
	}

	cmd := exec.Command("make", "cli_build", "--no-print-directory")
	cmd.Start()
	if err := cmd.Wait(); err != nil {
		output, _ := cmd.CombinedOutput()
		fmt.Printf("could not make '%s' binary:\n%v\n%s", binaryName, err, output)
		os.Exit(1)
	}

	err = os.Chdir(filepath.Join(d, "..", cliPath))
	if err != nil {
		fmt.Printf("could not change dir: %v", err)
		os.Exit(1)
	}

	fmt.Println(filepath.Join(d, "..", cliPath))

	/*



		snapshotsDir = filepath.Join(filepath.Dir(filename), "snapshots")

		configDir, err = filepath.Abs(cliPath)
		if err != nil {
			fmt.Printf("could not determine config dir: %v", err)
			os.Exit(1)
		}

		os.Exit(m.Run())
	*/
}

func RunCommandTestCase(t *testing.T, ts *httptest.Server, ctx *gosnap.Context, tc *CommandTestCase) {
	t.Run(strcase.ToSnake(strings.Replace(strings.Join(tc.Args, "_"), "-", "_", -1)), func(t *testing.T) {
		var s *gosnap.Snapshot
		if !ctx.Has(tc.Snapshot) {
			s = ctx.NewSnapshot(tc.Snapshot)
		} else {
			s = ctx.Get(tc.Snapshot)
		}

		a := append(tc.Args, "--host", ts.URL)
		c := exec.Command(fmt.Sprintf("./%s", binaryName), a...)
		env := os.Environ()
		for k, v := range tc.Env {
			env = append(env, fmt.Sprintf("%s=%s", k, v))
		}
		c.Env = env

		output, err := c.CombinedOutput()
		if tc.WantErr {
			assert.Error(t, err, "command '%s' was expected to throw an error", strings.Join(tc.Args, " "))
		} else {
			assert.NoError(t, err, "command '%s' wasn't expected to throw an error", strings.Join(tc.Args, " "))
		}

		if !t.Failed() && tc.Manual != nil {
			tc.Manual(t, string(output))
		}

		if !t.Failed() {
			s.AssertString(string(output))
		}
	})
}

func RunCommandTestCases(t *testing.T, mappingDir string, testCases []*CommandTestCase) {
	ts := CreateMockServerFromDir(t, mappingDir)
	defer ts.Close()

	ctx := gosnap.NewContext(t, filepath.Join(baseDir(t), "..", "snapshots"))

	for _, testCase := range testCases {
		RunCommandTestCase(t, ts, ctx, testCase)
	}
}
