package cmd

import (
	"github.com/fatih/color"
	out "github.com/plouc/go-gitlab-client/cli/output"
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

		out.Hook(output, outputFormat, hook)

		out.Meta(meta, false)

		return nil
	},
}
