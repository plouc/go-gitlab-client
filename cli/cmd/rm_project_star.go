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
	Use:     "project-star [project id]",
	Aliases: []string{"ps"},
	Short:   "Unstars a given project",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("you must specify a project id")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		projectId := args[0]

		color.Yellow("Unstaring project (project id: %s)…", projectId)

		confirmed := confirmAction(
			fmt.Sprintf("Are you sure you want to unstar project %s?", projectId),
			"aborted project star removal",
			autoConfirmRemoval,
		)
		if !confirmed {
			return
		}

		loader.Start()
		project, meta, err := client.UnstarProject(projectId)
		loader.Stop()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if meta.StatusCode == 304 {
			color.Red("\n  You didn't stared project %s!", projectId)
		}

		if project != nil {
			projectOutput(project, false)
		}

		metaOutput(meta, false)
	},
}
