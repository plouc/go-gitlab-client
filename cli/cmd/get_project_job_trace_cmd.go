package cmd

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	getCmd.AddCommand(getProjectJobTraceCmd)
}

var getProjectJobTraceCmd = &cobra.Command{
	Use:     resourceCmd("project-job-trace", "project-job"),
	Aliases: []string{"pjt"},
	Short:   "Get project job trace",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project-job", args)
		if err != nil {
			return err
		}

		jobId, err := strconv.Atoi(ids["job_id"])
		if err != nil {
			return err
		}

		color.Yellow("Fetching project's job trace (project id: %s, job id: %d)…", ids["project_id"], jobId)

		loader.Start()
		trace, meta, err := client.ProjectJobTrace(ids["project_id"], jobId)
		loader.Stop()
		if err != nil {
			return err
		}

		fmt.Fprintln(output, trace)

		printMeta(meta, false)

		return nil
	},
}
