package cmd

import (
	"os"
	"testing"

	"github.com/edison-moreland/go-gitlab-client/test"
)

func TestMain(m *testing.M) {
	test.BuildCli()

	os.Exit(m.Run())
}
