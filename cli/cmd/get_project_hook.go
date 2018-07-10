package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	getCmd.AddCommand(getProjectHookCmd)
}

var getProjectHookCmd = &cobra.Command{
	Use:     resourceCmd("project-hook", "project-hook"),
	Aliases: []string{"ph"},
	Short:   "Get project hook info",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project-hook", args)
		if err != nil {
			return err
		}

		color.Yellow("Fetching project's hook (project id: %s, hook id: %s)…", ids["project_id"], ids["hook_id"])

		loader.Start()
		hook, meta, err := client.ProjectHook(ids["project_id"], ids["hook_id"])
		loader.Stop()
		if err != nil {
			return err
		}

		hookOutput(hook)

		metaOutput(meta, false)

		return nil
	},
}
