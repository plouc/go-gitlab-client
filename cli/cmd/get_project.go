package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var projectStatistics bool

func init() {
	getCmd.AddCommand(getProjectCmd)

	getProjectCmd.Flags().BoolVarP(&projectStatistics, "statistics", "s", false, "Include project statistics")
}

var getProjectCmd = &cobra.Command{
	Use:     resourceCmd("project", "project"),
	Aliases: []string{"p"},
	Short:   "Get a specific project",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project", args)
		if err != nil {
			return err
		}

		color.Yellow("Fetching project (project id: %s)…", ids["project_id"])

		loader.Start()
		project, meta, err := client.Project(ids["project_id"], projectStatistics)
		loader.Stop()
		if err != nil {
			return err
		}

		projectOutput(project, projectStatistics)

		metaOutput(meta, false)

		return nil
	},
}
