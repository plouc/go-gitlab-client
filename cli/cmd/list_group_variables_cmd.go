package cmd

import (
	"github.com/fatih/color"
	out "github.com/plouc/go-gitlab-client/cli/output"
	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
)

func init() {
	listCmd.AddCommand(listGroupVariablesCmd)
}

var listGroupVariablesCmd = &cobra.Command{
	Use:     resourceCmd("group-variables", "group"),
	Aliases: []string{"group-vars", "gv"},
	Short:   "Get list of a group's variables",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "group", args)
		if err != nil {
			return err
		}

		color.Yellow("Fetching group variables (id: %s)…", ids["group_id"])

		o := &gitlab.PaginationOptions{
			Page:    page,
			PerPage: perPage,
		}

		loader.Start()
		collection, meta, err := client.GroupVariables(ids["group_id"], o)
		loader.Stop()
		if err != nil {
			return err
		}

		if len(collection.Items) == 0 {
			color.Red("  No variable found for group %s", ids["group_id"])
		} else {
			out.Variables(output, outputFormat, collection)
		}

		printMeta(meta, true)

		return nil
	},
}
