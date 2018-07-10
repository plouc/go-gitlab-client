package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
)

func init() {
	lsCmd.AddCommand(lsProjectVarsCmd)
}

var lsProjectVarsCmd = &cobra.Command{
	Use:     resourceCmd("project-vars", "project"),
	Aliases: []string{"pv"},
	Short:   "Get list of a project's variables",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project", args)
		if err != nil {
			return err
		}

		color.Yellow("Fetching project variables (id: %s)…", ids["project_id"])

		o := &gitlab.PaginationOptions{
			Page:    page,
			PerPage: perPage,
		}

		loader.Start()
		variables, meta, err := client.ProjectVariables(ids["project_id"], o)
		loader.Stop()
		if err != nil {
			return err
		}

		fmt.Println("")
		if len(variables) == 0 {
			color.Red("  No variable found for project %s", ids["project_id"])
		} else {
			varsOutput(variables)
		}

		metaOutput(meta, true)

		return nil
	},
}
