package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	out "github.com/plouc/go-gitlab-client/cli/output"
	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
)

func init() {
	addCmd.AddCommand(addProjectCmd)
}

var addProjectCmd = &cobra.Command{
	Use:     "project",
	Aliases: []string{"p"},
	Short:   "Create a new project",
	Run: func(cmd *cobra.Command, args []string) {
		color.Yellow("Creating project…")

		project := gitlab.ProjectAddPayload{}

		prompt := promptui.Prompt{
			Label: "Name",
		}
		name, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		project.Name = name

		prompt = promptui.Prompt{
			Label: "Path",
		}
		path, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		project.Path = path

		loader.Start()
		createdProject, meta, err := client.AddProject(&project)
		loader.Stop()

		out.Project(output, outputFormat, createdProject, false)

		printMeta(meta, false)
	},
}
