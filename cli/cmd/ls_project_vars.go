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
	Use:     "project-vars [project id]",
	Aliases: []string{"pv"},
	Short:   "Get list of a project's variables",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("you must specify a project id")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		projectId := args[0]

		color.Yellow("Fetching project variables (id: %s)…", projectId)

		o := &gitlab.PaginationOptions{
			Page:    page,
			PerPage: perPage,
		}

		loader.Start()
		variables, meta, err := client.ProjectVariables(projectId, o)
		loader.Stop()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("")
		if len(variables) == 0 {
			color.Red("  No variable found for project %s", projectId)
		} else {
			varsOutput(variables)
		}

		metaOutput(meta, true)
	},
}
