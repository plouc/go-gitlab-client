package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"strconv"
)

func init() {
	getProjectJobCmd.AddCommand(getProjectJobCancelCmd)
}

var getProjectJobCancelCmd = &cobra.Command{
	Use:   resourceCmd("cancel", "project-job"),
	Short: "Cancel project job",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project-job", args)
		if err != nil {
			return err
		}

		jobId, err := strconv.Atoi(ids["job_id"])
		if err != nil {
			return err
		}

		color.Yellow("Cancelling project's job (project id: %s, job id: %d)…", ids["project_id"], jobId)

		loader.Start()
		job, meta, err := client.CancelProjectJob(ids["project_id"], jobId)
		loader.Stop()
		if err != nil {
			return err
		}

		jobOutput(job)

		metaOutput(meta, false)

		return nil
	},
}
