package cmd

import (
	"strconv"

	"github.com/fatih/color"
	out "github.com/plouc/go-gitlab-client/cli/output"
	"github.com/spf13/cobra"
)

func init() {
	getCmd.AddCommand(getRunnerCmd)
}

var getRunnerCmd = &cobra.Command{
	Use:     resourceCmd("runner", "runner"),
	Aliases: []string{"r"},
	Short:   "Get details of a runner",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "runner", args)
		if err != nil {
			return err
		}

		runnerId, err := strconv.Atoi(ids["runner_id"])
		if err != nil {
			return err
		}

		color.Yellow("Fetching runner (id: %d)…", runnerId)

		loader.Start()
		runner, meta, err := client.Runner(runnerId)
		loader.Stop()
		if err != nil {
			return err
		}

		out.Runner(output, outputFormat, runner)

		printMeta(meta, false)

		return nil
	},
}
