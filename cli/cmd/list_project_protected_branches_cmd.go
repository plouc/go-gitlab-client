package cmd

import (
	"fmt"

	"github.com/fatih/color"
	out "github.com/plouc/go-gitlab-client/cli/output"
	"github.com/spf13/cobra"
)

func init() {
	listCmd.AddCommand(listProtectedBranchesCmd)
}

var listProtectedBranchesCmd = &cobra.Command{
	Use:     resourceCmd("project-protected-branches", "project"),
	Aliases: []string{"ppb"},
	Short:   "List project protected branches",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project", args)
		if err != nil {
			return err
		}

		color.Yellow("Fetching project protected branches (id: %s)…", ids["project_id"])

		loader.Start()
		protectedBranches, meta, err := client.ProtectedBranches(ids["project_id"], nil)
		loader.Stop()
		if err != nil {
			return err
		}

		fmt.Println("")
		if len(protectedBranches) == 0 {
			color.Red("No protected branch found for project %s", ids["project_id"])
		} else {
			out.ProtectedBranches(output, outputFormat, protectedBranches)
		}

		printMeta(meta, true)

		return nil
	},
}
