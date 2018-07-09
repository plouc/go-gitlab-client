package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	getCmd.AddCommand(getProjectHookCmd)
}

var getProjectHookCmd = &cobra.Command{
	Use:     "project-hook [project id] [hook id]",
	Aliases: []string{"ph"},
	Short:   "Get project hook info",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return fmt.Errorf("you must specify a project id and a hook id")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		projectId := args[0]
		hookId := args[1]

		color.Yellow("Fetching project's hook (project id: %s, hook id: %s)…", projectId, hookId)

		loader.Start()
		hook, meta, err := client.ProjectHook(projectId, hookId)
		loader.Stop()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		hookOutput(hook)

		metaOutput(meta, false)
	},
}
