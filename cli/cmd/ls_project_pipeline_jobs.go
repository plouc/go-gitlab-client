package cmd

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
	out "github.com/plouc/go-gitlab-client/cli/output"
	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
)

var projectPipelineJobsScope string
var projectPipelineJobsPrettyOutput bool

func init() {
	lsCmd.AddCommand(lsProjectPipelineJobsCmd)

	lsProjectPipelineJobsCmd.Flags().StringVarP(&projectPipelineJobsScope, "scope", "s", "", "Scope")
	lsProjectPipelineJobsCmd.Flags().BoolVar(&projectPipelineJobsPrettyOutput, "pretty", false, "Use custom output formatting")
}

func fetchProjectPipelineJobs(projectId string, pipelineId int) {
	color.Yellow("Fetching project's pipeline jobs (project id: %s, pipelined id: %d)…", projectId, pipelineId)

	o := &gitlab.JobsOptions{}
	o.Page = page
	o.PerPage = perPage
	if projectJobsScope != "" {
		o.Scope = []string{projectPipelineJobsScope}
	}

	loader.Start()
	collection, meta, err := client.ProjectPipelineJobs(projectId, pipelineId, o)
	loader.Stop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(collection.Items) == 0 {
		color.Red("No job found for project %s pipeline %d", projectId, pipelineId)
	} else {
		out.Jobs(output, outputFormat, collection, projectPipelineJobsPrettyOutput)
	}

	printMeta(meta, true)

	handlePaginatedResult(meta, func() {
		fetchProjectPipelineJobs(projectId, pipelineId)
	})
}

var lsProjectPipelineJobsCmd = &cobra.Command{
	Use:     resourceCmd("project-pipeline-jobs", "project-pipeline"),
	Aliases: []string{"ppj"},
	Short:   "List project pipeline jobs",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project-pipeline", args)
		if err != nil {
			return err
		}

		pipelineId, err := strconv.Atoi(ids["pipeline_id"])
		if err != nil {
			return err
		}

		fetchProjectPipelineJobs(ids["project_id"], pipelineId)

		return nil
	},
}
