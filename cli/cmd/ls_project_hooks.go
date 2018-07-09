package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	lsCmd.AddCommand(lsProjectHooksCmd)
}

func fetchHooks(projectId string) {
	color.Yellow("Fetching project's hooks (id: %s)…", projectId)

	loader.Start()
	hooks, meta, err := client.ProjectHooks(projectId)
	loader.Stop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(hooks) == 0 {
		color.Red("  No hook found for project %s", projectId)
	} else {
		hooksOutput(hooks)
	}

	metaOutput(meta, true)

	handlePaginatedResult(meta, func() {
		fetchHooks(projectId)
	})
}

var lsProjectHooksCmd = &cobra.Command{
	Use:     "project-hooks [project id]",
	Aliases: []string{"ph"},
	Short:   "List project's hooks",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("you must specify a project id")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fetchHooks(args[0])
	},
}
