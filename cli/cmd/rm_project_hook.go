package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rmCmd.AddCommand(rmProjectHookCmd)
}

var rmProjectHookCmd = &cobra.Command{
	Use:     "project-hook [project id] [hook id]",
	Aliases: []string{"ph"},
	Short:   "Remove project hook",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return fmt.Errorf("you must specify a project id and a hook id")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		projectId := args[0]
		hookId := args[1]

		color.Yellow("Removing project hook (project id: %s, hook id: %s)…", projectId, hookId)

		confirmed := confirmAction(
			fmt.Sprintf("Are you sure you want to remove project %s hook %s?", projectId, hookId),
			"aborted project hook removal",
			autoConfirmRemoval,
		)
		if !confirmed {
			return
		}

		loader.Start()
		meta, err := client.RemoveProjectHook(projectId, hookId)
		loader.Stop()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		metaOutput(meta, false)
	},
}
