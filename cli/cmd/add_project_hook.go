package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
)

func init() {
	addCmd.AddCommand(addProjectHookCmd)
}

var addProjectHookCmd = &cobra.Command{
	Use:     "project-hook [project id]",
	Aliases: []string{"ph"},
	Short:   "Create a new hook for given project",
	Run: func(cmd *cobra.Command, args []string) {
		projectId := args[0]

		color.Yellow("Creating hook for project (project id: %s)…", projectId)

		hook := gitlab.HookAddPayload{}

		prompt := promptui.Prompt{
			Label: "Url",
		}
		u, err := prompt.Run()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		hook.Url = u

		selectPrompt := promptui.Select{
			Label: color.YellowString("Push events?"),
			Items: []string{"yes", "no"},
		}
		idx, _, err := selectPrompt.Run()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if idx == 0 {
			hook.PushEvents = true
		}

		selectPrompt = promptui.Select{
			Label: color.YellowString("Issue events?"),
			Items: []string{"yes", "no"},
		}
		idx, _, err = selectPrompt.Run()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if idx == 0 {
			hook.IssuesEvents = true
		}

		selectPrompt = promptui.Select{
			Label: color.YellowString("Confidential issues events?"),
			Items: []string{"yes", "no"},
		}
		idx, _, err = selectPrompt.Run()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if idx == 0 {
			hook.ConfidentialIssuesEvents = true
		}

		selectPrompt = promptui.Select{
			Label: color.YellowString("Merge requests events?"),
			Items: []string{"yes", "no"},
		}
		idx, _, err = selectPrompt.Run()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if idx == 0 {
			hook.MergeRequestsEvents = true
		}

		selectPrompt = promptui.Select{
			Label: color.YellowString("Tag push events?"),
			Items: []string{"yes", "no"},
		}
		idx, _, err = selectPrompt.Run()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if idx == 0 {
			hook.TagPushEvents = true
		}

		selectPrompt = promptui.Select{
			Label: color.YellowString("Note events?"),
			Items: []string{"yes", "no"},
		}
		idx, _, err = selectPrompt.Run()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if idx == 0 {
			hook.NoteEvents = true
		}

		selectPrompt = promptui.Select{
			Label: color.YellowString("Job events?"),
			Items: []string{"yes", "no"},
		}
		idx, _, err = selectPrompt.Run()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if idx == 0 {
			hook.JobEvents = true
		}

		selectPrompt = promptui.Select{
			Label: color.YellowString("Pipeline events?"),
			Items: []string{"yes", "no"},
		}
		idx, _, err = selectPrompt.Run()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if idx == 0 {
			hook.PipelineEvents = true
		}

		selectPrompt = promptui.Select{
			Label: color.YellowString("Wiki page events?"),
			Items: []string{"yes", "no"},
		}
		idx, _, err = selectPrompt.Run()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if idx == 0 {
			hook.WikiPageEvents = true
		}

		selectPrompt = promptui.Select{
			Label: color.YellowString("Enable ssl verification?"),
			Items: []string{"yes", "no"},
		}
		idx, _, err = selectPrompt.Run()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if idx == 0 {
			hook.EnableSslVerification = true
		}

		loader.Start()
		createdHook, meta, err := client.AddProjectHook(projectId, &hook)
		loader.Stop()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		hookOutput(createdHook)

		metaOutput(meta, false)
	},
}
