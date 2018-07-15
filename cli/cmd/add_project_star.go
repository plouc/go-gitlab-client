package cmd

import (
	"github.com/fatih/color"
	out "github.com/plouc/go-gitlab-client/cli/output"
	"github.com/spf13/cobra"
)

func init() {
	addCmd.AddCommand(addProjectStarCmd)
}

var addProjectStarCmd = &cobra.Command{
	Use:     resourceCmd("project-star", "project"),
	Aliases: []string{"ps"},
	Short:   "Stars a given project",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project", args)
		if err != nil {
			return err
		}

		color.Yellow("Staring project (project id: %s)…", ids["project_id"])

		loader.Start()
		project, meta, err := client.StarProject(ids["project_id"])
		loader.Stop()
		if err != nil {
			return err
		}

		if meta.StatusCode == 304 {
			color.Red("\n  You already stared project %s!", ids["project_id"])
		}

		if project != nil {
			out.Project(output, outputFormat, project, false)
		}

		printMeta(meta, false)

		return nil
	},
}
