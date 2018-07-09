package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	addCmd.AddCommand(addProjectStarCmd)
}

var addProjectStarCmd = &cobra.Command{
	Use:     "project-star [project id]",
	Aliases: []string{"ps"},
	Short:   "Stars a given project",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("you must specify a project id")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		projectId := args[0]

		color.Yellow("Staring project (project id: %s)…", projectId)

		loader.Start()
		project, meta, err := client.StarProject(projectId)
		loader.Stop()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if meta.StatusCode == 304 {
			color.Red("\n  You already stared project %s!", projectId)
		}

		if project != nil {
			projectOutput(project, false)
		}

		metaOutput(meta, false)
	},
}
