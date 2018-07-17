package cmd

import (
	"os"
	"testing"

	"github.com/plouc/go-gitlab-client/test"
)

func TestMain(m *testing.M) {
	test.BuildCli()

	os.Exit(m.Run())
}
