package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"strings"
)

type relatedCommand struct {
	cmd  *cobra.Command
	args map[string]string
}

func newRelatedCommand(cmd *cobra.Command, args map[string]string) *relatedCommand {
	return &relatedCommand{cmd, args}
}

func getFullCommand(path string, cmd *cobra.Command) string {
	var p string
	if path == "" {
		p = fmt.Sprintf("%s %s", cmd.Use, path)
	} else {
		parts := strings.Split(cmd.Use, " ")
		p = fmt.Sprintf("%s %s", parts[0], path)
	}

	parent := cmd.Parent()
	if parent != nil {
		return getFullCommand(p, parent)
	}

	return p
}

func relatedCommands(cmds []*relatedCommand) {
	if !verbose {
		return
	}

	color.Yellow("Related commands:")
	fmt.Println("")
	for _, r := range cmds {
		full := getFullCommand("", r.cmd)
		for k, v := range r.args {
			full = strings.Replace(full, strings.ToUpper(k), v, -1)
		}

		fmt.Printf("- %-42s  %s\n", r.cmd.Short, color.YellowString(full))
	}
	fmt.Println("")
}
