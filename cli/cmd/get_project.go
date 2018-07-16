package cmd

import (
	"github.com/fatih/color"
	out "github.com/plouc/go-gitlab-client/cli/output"
	"github.com/spf13/cobra"
)

var projectStatistics bool

func init() {
	getCmd.AddCommand(getProjectCmd)

	getProjectCmd.Flags().BoolVarP(&projectStatistics, "statistics", "s", false, "Include project statistics")
}

var getProjectCmd = &cobra.Command{
	Use:     resourceCmd("project", "project"),
	Aliases: []string{"p"},
	Short:   "Get a specific project",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project", args)
		if err != nil {
			return err
		}

		color.Yellow("Fetching project (project id: %s)…", ids["project_id"])

		loader.Start()
		project, meta, err := client.Project(ids["project_id"], projectStatistics)
		loader.Stop()
		if err != nil {
			return err
		}

		out.Project(output, outputFormat, project, projectStatistics)

		printMeta(meta, false)

		relatedCommands([]*relatedCommand{
			newRelatedCommand(lsProjectBranchesCmd, map[string]string{
				"project_id": ids["project_id"],
			}),
			newRelatedCommand(addProjectBranchCmd, map[string]string{
				"project_id": ids["project_id"],
			}),
			newRelatedCommand(lsProjectHooksCmd, map[string]string{
				"project_id": ids["project_id"],
			}),
			newRelatedCommand(addProjectHookCmd, map[string]string{
				"project_id": ids["project_id"],
			}),
			newRelatedCommand(lsProjectMembersCmd, map[string]string{
				"project_id": ids["project_id"],
			}),
			newRelatedCommand(lsProjectPipelinesCmd, map[string]string{
				"project_id": ids["project_id"],
			}),
			newRelatedCommand(lsProjectJobsCmd, map[string]string{
				"project_id": ids["project_id"],
			}),
			newRelatedCommand(lsProjectMergeRequestsCmd, map[string]string{
				"project_id": ids["project_id"],
			}),
			newRelatedCommand(lsProjectCommitsCmd, map[string]string{
				"project_id": ids["project_id"],
			}),
			newRelatedCommand(lsProjectEnvironmentsCmd, map[string]string{
				"project_id": ids["project_id"],
			}),
			newRelatedCommand(addProjectEnvironmentCmd, map[string]string{
				"project_id": ids["project_id"],
			}),
			newRelatedCommand(lsProjectVariablesCmd, map[string]string{
				"project_id": ids["project_id"],
			}),
			newRelatedCommand(addProjectBadgeCmd, map[string]string{
				"project_id": ids["project_id"],
			}),
		})

		return nil
	},
}
