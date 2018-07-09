package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/plouc/go-gitlab-client/gogitlab"
	"github.com/spf13/cobra"
)

func init() {
	addCmd.AddCommand(addProjectBadgeCmd)
}

var addProjectBadgeCmd = &cobra.Command{
	Use:     "project-badge [project id]",
	Aliases: []string{"pbdg"},
	Short:   "Create project badge",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("you must specify a project id")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		projectId := args[0]

		color.Yellow("Creating project's badge (project id: %s)…", projectId)

		badge := new(gogitlab.Badge)

		prompt := promptui.Prompt{
			Label: "LinkUrl",
		}
		linkUrl, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		badge.LinkUrl = linkUrl

		prompt = promptui.Prompt{
			Label: "ImageUrl",
		}
		imageUrl, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		badge.ImageUrl = imageUrl

		loader.Start()
		createdBadge, meta, err := client.AddProjectBadge(projectId, badge)
		loader.Stop()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		badgeOutput(createdBadge)

		metaOutput(meta, false)
	},
}
