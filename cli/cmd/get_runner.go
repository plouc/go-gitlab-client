package cmd

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	getCmd.AddCommand(getRunnerCmd)
}

var getRunnerCmd = &cobra.Command{
	Use:     "runner [runner id]",
	Aliases: []string{"r"},
	Short:   "Get details of a runner",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("you must specify a runner id")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		runnerId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		color.Yellow("Fetching runner (id: %d)…", runnerId)

		loader.Start()
		runner, meta, err := client.Runner(runnerId)
		loader.Stop()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		runnerOutput(runner)

		metaOutput(meta, false)
	},
}
