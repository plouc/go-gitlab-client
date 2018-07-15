package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"strconv"
)

func init() {
	rmCmd.AddCommand(rmProjectEnvironmentCmd)
}

var rmProjectEnvironmentCmd = &cobra.Command{
	Use:     resourceCmd("project-environment", "project-environment"),
	Aliases: []string{"project-env", "pe"},
	Short:   "Remove project environment",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project-environment", args)
		if err != nil {
			return err
		}

		environmentId, err := strconv.Atoi(ids["environment_id"])
		if err != nil {
			return err
		}

		confirmed := confirmAction(
			fmt.Sprintf("Are you sure you want to remove project %s environment %d?", ids["project_id"], environmentId),
			"aborted project environment removal",
			autoConfirmRemoval,
		)
		if !confirmed {
			return nil
		}

		color.Yellow("Removing project environment (project id: %s, environment id: %d)…", ids["project_id"], environmentId)

		loader.Start()
		meta, err := client.RemoveProjectEnvironment(ids["project_id"], environmentId)
		loader.Stop()
		if err != nil {
			return err
		}

		color.Green("✔ Project environment was successfully removed")

		metaOutput(meta, false)

		return nil
	},
}
