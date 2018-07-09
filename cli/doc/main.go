package main

import (
	"github.com/plouc/go-gitlab-client/cli/cmd"
	"github.com/spf13/cobra/doc"
	"log"
)

func main() {
	err := doc.GenMarkdownTree(cmd.RootCmd, "./")
	if err != nil {
		log.Fatal(err)
	}
}
