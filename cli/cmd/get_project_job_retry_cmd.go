package cmd

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
	out "github.com/plouc/go-gitlab-client/cli/output"
	"github.com/spf13/cobra"
)

func init() {
	getProjectJobCmd.AddCommand(getProjectJobRetryCmd)
}

var getProjectJobRetryCmd = &cobra.Command{
	Use:   resourceCmd("retry", "project-job"),
	Short: "Retry project job",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project-job", args)
		if err != nil {
			return err
		}

		jobId, err := strconv.Atoi(ids["job_id"])
		if err != nil {
			return err
		}

		color.Yellow("Retrying project's job (project id: %s, job id: %d)…", ids["project_id"], jobId)

		loader.Start()
		job, meta, err := client.RetryProjectJob(ids["project_id"], jobId)
		loader.Stop()
		if err != nil {
			return err
		}

		out.Job(output, outputFormat, job)

		printMeta(meta, false)

		fmt.Println("")
		color.Yellow("Related")
		fmt.Printf(
			"  To get job trace, run: %s\n",
			color.YellowString("glc get project-job-trace %s %d", ids["project_id"], job.Id),
		)
		fmt.Println("")

		return nil
	},
}
