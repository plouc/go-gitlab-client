package cmd

import (
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
)

func init() {
	addCmd.AddCommand(addProjectBadgeCmd)
}

var addProjectBadgeCmd = &cobra.Command{
	Use:     resourceCmd("project-badge", "project"),
	Aliases: []string{"pbdg"},
	Short:   "Create project badge",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project", args)
		if err != nil {
			return err
		}

		color.Yellow("Creating project's badge (project id: %s)…", ids["project_id"])

		badge := new(gitlab.Badge)

		prompt := promptui.Prompt{
			Label: "LinkUrl",
		}
		linkUrl, err := prompt.Run()
		if err != nil {
			return err
		}
		badge.LinkUrl = linkUrl

		prompt = promptui.Prompt{
			Label: "ImageUrl",
		}
		imageUrl, err := prompt.Run()
		if err != nil {
			return err
		}
		badge.ImageUrl = imageUrl

		loader.Start()
		createdBadge, meta, err := client.AddProjectBadge(ids["project_id"], badge)
		loader.Stop()
		if err != nil {
			return err
		}

		badgeOutput(createdBadge)

		metaOutput(meta, false)

		return nil
	},
}
