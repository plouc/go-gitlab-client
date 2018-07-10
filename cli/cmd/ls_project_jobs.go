package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
)

var projectJobsScope string

func init() {
	lsCmd.AddCommand(lsProjectJobsCmd)

	lsProjectJobsCmd.Flags().StringVarP(&projectJobsScope, "scope", "s", "", "Scope")
}

func fetchProjectJobs(projectId string) {
	color.Yellow("Fetching project's jobs (project id: %s)…", projectId)

	o := &gitlab.JobsOptions{}
	o.Page = page
	o.PerPage = perPage
	if projectJobsScope != "" {
		o.Scope = []string{projectJobsScope}
	}

	loader.Start()
	jobs, meta, err := client.ProjectJobs(projectId, o)
	loader.Stop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(jobs) == 0 {
		color.Red("No job found for project %s", projectId)
	} else {
		jobsOutput(jobs)
	}

	metaOutput(meta, true)

	handlePaginatedResult(meta, func() {
		fetchProjectJobs(projectId)
	})
}

var lsProjectJobsCmd = &cobra.Command{
	Use:     resourceCmd("project-jobs", "project"),
	Aliases: []string{"pj"},
	Short:   "List project jobs",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project", args)
		if err != nil {
			return err
		}

		fetchProjectJobs(ids["project_id"])

		return nil
	},
}
