package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"strconv"
)

func init() {
	getCmd.AddCommand(getProjectJobCmd)
}

var getProjectJobCmd = &cobra.Command{
	Use:     resourceCmd("project-job", "project-job"),
	Aliases: []string{"pj"},
	Short:   "Get project job info",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project-job", args)
		if err != nil {
			return err
		}

		jobId, err := strconv.Atoi(ids["job_id"])
		if err != nil {
			return err
		}

		color.Yellow("Fetching project's job (project id: %s, job id: %d)…", ids["project_id"], jobId)

		loader.Start()
		job, meta, err := client.ProjectJob(ids["project_id"], jobId)
		loader.Stop()
		if err != nil {
			return err
		}

		jobOutput(job)

		metaOutput(meta, false)

		fmt.Println("")
		color.Yellow("Related")
		fmt.Printf(
			"  To get job trace, run: %s\n",
			color.YellowString("glc get project-job-trace %s %d", ids["project_id"], jobId),
		)
		fmt.Println("")

		return nil
	},
}
