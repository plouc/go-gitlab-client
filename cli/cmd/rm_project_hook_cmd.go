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
	Use:     resourceCmd("project-hook", "project-hook"),
	Aliases: []string{"ph"},
	Short:   "Remove project hook",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project-hook", args)
		if err != nil {
			return err
		}

		color.Yellow("Removing project hook (project id: %s, hook id: %s)…", ids["project_id"], ids["hook_id"])

		confirmed := confirmAction(
			fmt.Sprintf("Are you sure you want to remove project %s hook %s?", ids["project_id"], ids["hook_id"]),
			"aborted project hook removal",
			autoConfirmRemoval,
		)
		if !confirmed {
			return nil
		}

		loader.Start()
		meta, err := client.RemoveProjectHook(ids["project_id"], ids["hook_id"])
		loader.Stop()
		if err != nil {
			return err
		}

		printMeta(meta, false)

		return nil
	},
}
