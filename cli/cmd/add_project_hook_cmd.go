package cmd

import (
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	out "github.com/plouc/go-gitlab-client/cli/output"
	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
	"strconv"
)

func init() {
	addCmd.AddCommand(addProjectHookCmd)
}

var addProjectHookCmd = &cobra.Command{
	Use:     resourceCmd("project-hook", "project"),
	Aliases: []string{"ph"},
	Short:   "Create a new hook for given project",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project", args)
		if err != nil {
			return err
		}

		color.Yellow("Creating hook for project (project id: %s)…", ids["project_id"])

		hook := gitlab.HookAddPayload{}

		prompt := promptui.Prompt{
			Label: "Url",
		}
		u, err := prompt.Run()
		if err != nil {
			return err
		}
		hook.Url = u

		selectPrompt := promptui.Select{
			Label: color.YellowString("Push events?"),
			Items: []string{"yes", "no"},
		}
		idx, _, err := selectPrompt.Run()
		if err != nil {
			return err
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
			return err
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
			return err
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
			return err
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
			return err
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
			return err
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
			return err
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
			return err
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
			return err
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
			return err
		}
		if idx == 0 {
			hook.EnableSslVerification = true
		}

		loader.Start()
		createdHook, meta, err := client.AddProjectHook(ids["project_id"], &hook)
		loader.Stop()
		if err != nil {
			return err
		}

		out.Hook(output, outputFormat, createdHook)

		printMeta(meta, false)

		relatedCommands([]*relatedCommand{
			newRelatedCommand(listProjectHooksCmd, map[string]string{
				"project_id": ids["project_id"],
			}),
			newRelatedCommand(getProjectHookCmd, map[string]string{
				"project_id": ids["project_id"],
				"hook_id":    strconv.Itoa(createdHook.Id),
			}),
		})

		return nil
	},
}
