package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
)

func init() {
	lsCmd.AddCommand(lsProjectPipelinesCmd)
}

func fetchProjectPipelines(projectId string) {
	color.Yellow("Fetching project's pipelines (project id: %s)…", projectId)

	o := &gitlab.PipelinesOptions{}
	o.Page = page
	o.PerPage = perPage

	loader.Start()
	pipelines, meta, err := client.ProjectPipelines(projectId, o)
	loader.Stop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("")
	if len(pipelines) == 0 {
		color.Red("No pipeline found for project %s", projectId)
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{
			"Id",
			"Ref",
			"Sha",
			"Status",
		})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		for _, pipeline := range pipelines {
			table.Append([]string{
				fmt.Sprintf("%d", pipeline.Id),
				pipeline.Ref,
				pipeline.Sha,
				pipeline.Status,
			})
		}
		table.Render()
	}
	fmt.Println("")

	metaOutput(meta, true)

	handlePaginatedResult(meta, func() {
		fetchProjectPipelines(projectId)
	})
}

var lsProjectPipelinesCmd = &cobra.Command{
	Use:     "project-pipelines [project id]",
	Aliases: []string{"pp"},
	Short:   "List project pipelines",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("you must specify a project id")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fetchProjectPipelines(args[0])
	},
}
