package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	getCmd.AddCommand(getProjectPipelineCmd)
}

var getProjectPipelineCmd = &cobra.Command{
	Use:     "project-pipeline [project id] [pipeline id]",
	Aliases: []string{"pp"},
	Short:   "Get project pipeline details",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return fmt.Errorf("you must specify a project id and a pipeline id")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		projectId := args[0]
		pipelineId := args[1]

		color.Yellow("Fetching project's pipeline (project id: %s, pipeline id: %s)…", projectId, pipelineId)

		loader.Start()
		pipeline, meta, err := client.ProjectPipeline(projectId, pipelineId)
		loader.Stop()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		pipelineOutput(pipeline)

		metaOutput(meta, false)
	},
}
