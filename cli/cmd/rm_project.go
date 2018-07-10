package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rmCmd.AddCommand(rmProjectCmd)
}

var rmProjectCmd = &cobra.Command{
	Use:     resourceCmd("project", "project"),
	Aliases: []string{"p"},
	Short:   "Remove project",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project", args)
		if err != nil {
			return err
		}

		color.Yellow("Removing project (id: %s)…", ids["project_id"])

		confirmed := confirmAction(
			fmt.Sprintf("Are you sure you want to remove project %s?", ids["project_id"]),
			"aborted project removal",
			autoConfirmRemoval,
		)
		if !confirmed {
			return nil
		}

		loader.Start()
		message, meta, err := client.RemoveProject(ids["project_id"])
		loader.Stop()
		if err != nil {
			return err
		}

		color.Green("✔ Project was successfully removed: %s", message)

		metaOutput(meta, false)

		return nil
	},
}
