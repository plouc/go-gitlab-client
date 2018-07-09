package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	lsCmd.AddCommand(lsProtectedBranchesCmd)
}

var lsProtectedBranchesCmd = &cobra.Command{
	Use:     "project-protected-branches [project id]",
	Aliases: []string{"ppb"},
	Short:   "List project protected branches",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("you must specify a project id")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		projectId := args[0]

		color.Yellow("Fetching project protected branches (id: %s)…", projectId)

		loader.Start()
		protectedBranches, meta, err := client.ProtectedBranches(projectId, nil)
		loader.Stop()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("")
		if len(protectedBranches) == 0 {
			color.Red("No protected branch found for project %s", projectId)
		} else {
			for _, protectedBranch := range protectedBranches {
				printProtectedBranch(protectedBranch)
				fmt.Println("")
			}
		}

		metaOutput(meta, true)
	},
}
