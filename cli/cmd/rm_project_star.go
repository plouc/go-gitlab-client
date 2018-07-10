package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rmCmd.AddCommand(rmProjectStarCmd)
}

var rmProjectStarCmd = &cobra.Command{
	Use:     resourceCmd("project-star", "project"),
	Aliases: []string{"ps"},
	Short:   "Unstars a given project",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project", args)
		if err != nil {
			return err
		}

		color.Yellow("Unstaring project (project id: %s)…", ids["project_id"])

		confirmed := confirmAction(
			fmt.Sprintf("Are you sure you want to unstar project %s?", ids["project_id"]),
			"aborted project star removal",
			autoConfirmRemoval,
		)
		if !confirmed {
			return nil
		}

		loader.Start()
		project, meta, err := client.UnstarProject(ids["project_id"])
		loader.Stop()
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		if meta.StatusCode == 304 {
			color.Red("\n  You didn't stared project %s!", ids["project_id"])
		}

		if project != nil {
			projectOutput(project, false)
		}

		metaOutput(meta, false)

		return nil
	},
}
