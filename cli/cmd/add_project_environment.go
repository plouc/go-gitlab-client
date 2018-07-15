package cmd

import (
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	out "github.com/plouc/go-gitlab-client/cli/output"
	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
)

func init() {
	addCmd.AddCommand(addProjectEnvironmentCmd)
}

var addProjectEnvironmentCmd = &cobra.Command{
	Use:     resourceCmd("project-environment", "project"),
	Aliases: []string{"project-env", "pe"},
	Short:   "Create project environment",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project", args)
		if err != nil {
			return err
		}

		color.Yellow("Creating project environment (project id: %s)…", ids["project_id"])

		environment := new(gitlab.EnvironmentAddPayload)

		prompt := promptui.Prompt{
			Label: "Name",
		}
		name, err := prompt.Run()
		if err != nil {
			return err
		}
		environment.Name = name

		prompt = promptui.Prompt{
			Label: "ExternalUrl",
		}
		externalUrl, err := prompt.Run()
		if err != nil {
			return err
		}
		environment.ExternalUrl = externalUrl

		loader.Start()
		createdEnvironment, meta, err := client.AddProjectEnvironment(ids["project_id"], environment)
		loader.Stop()
		if err != nil {
			return err
		}

		out.Environment(output, outputFormat, createdEnvironment)

		out.Meta(meta, false)

		return nil
	},
}
