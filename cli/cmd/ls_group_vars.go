package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/plouc/go-gitlab-client/gogitlab"
	"github.com/spf13/cobra"
)

func init() {
	lsCmd.AddCommand(lsGroupVarsCmd)
}

var lsGroupVarsCmd = &cobra.Command{
	Use:     "group-vars [group id]",
	Aliases: []string{"gv"},
	Short:   "Get list of a group's variables",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("you must specify a group id")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		groupId := args[0]

		color.Yellow("Fetching group variables (id: %s)…", groupId)

		o := &gogitlab.PaginationOptions{
			Page:    page,
			PerPage: perPage,
		}

		loader.Start()
		variables, meta, err := client.GroupVariables(groupId, o)
		loader.Stop()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if len(variables) == 0 {
			color.Red("  No variable found for group %s", groupId)
		} else {
			varsOutput(variables)
		}

		metaOutput(meta, true)
	},
}
