package cmd

import (
	"fmt"

	"github.com/fatih/color"
	out "github.com/plouc/go-gitlab-client/cli/output"
	"github.com/spf13/cobra"
)

func init() {
	listCmd.AddCommand(listProjectHooksCmd)
}

func fetchHooks(projectId string) {
	color.Yellow("Fetching project's hooks (id: %s)…", projectId)

	loader.Start()
	collection, meta, err := client.ProjectHooks(projectId)
	loader.Stop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(collection.Items) == 0 {
		color.Red("  No hook found for project %s", projectId)
	} else {
		out.Hooks(output, outputFormat, collection)
	}

	printMeta(meta, true)

	handlePaginatedResult(meta, func() {
		fetchHooks(projectId)
	})
}

var listProjectHooksCmd = &cobra.Command{
	Use:     resourceCmd("project-hooks", "project"),
	Aliases: []string{"ph"},
	Short:   "List project's hooks",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project", args)
		if err != nil {
			return err
		}

		fetchHooks(ids["project_id"])

		return nil
	},
}
