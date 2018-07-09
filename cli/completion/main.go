package main

import (
	"github.com/plouc/go-gitlab-client/cli/cmd"
)

func main() {
	cmd.RootCmd.GenBashCompletionFile("completion.sh")
}
