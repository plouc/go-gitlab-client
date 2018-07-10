package cmd

import (
	"github.com/fatih/color"
	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
)

func init() {
	lsCmd.AddCommand(lsGroupVarsCmd)
}

var lsGroupVarsCmd = &cobra.Command{
	Use:     resourceCmd("group-vars", "group"),
	Aliases: []string{"gv"},
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
		variables, meta, err := client.GroupVariables(ids["group_id"], o)
		loader.Stop()
		if err != nil {
			return err
		}

		if len(variables) == 0 {
			color.Red("  No variable found for group %s", ids["group_id"])
		} else {
			varsOutput(variables)
		}

		metaOutput(meta, true)

		return nil
	},
}
