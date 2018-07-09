package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var projectStatistics bool

func init() {
	getCmd.AddCommand(getProjectCmd)

	getProjectCmd.Flags().BoolVarP(&projectStatistics, "statistics", "s", false, "Include project statistics")
}

var getProjectCmd = &cobra.Command{
	Use:     "project [project id]",
	Aliases: []string{"p"},
	Short:   "Get a specific project",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("you must specify a project id")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		projectId := args[0]

		color.Yellow("Fetching project (project id: %s)…", projectId)

		loader.Start()
		project, meta, err := client.Project(projectId, projectStatistics)
		loader.Stop()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		projectOutput(project, projectStatistics)

		metaOutput(meta, false)
	},
}
