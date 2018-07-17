package cmd

import (
	"fmt"

	"github.com/fatih/color"
	out "github.com/plouc/go-gitlab-client/cli/output"
	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
)

var runnersScope string
var runnersAll bool

func init() {
	listCmd.AddCommand(listRunnersCmd)

	listRunnersCmd.Flags().StringVarP(&runnersScope, "scope", "s", "", "The scope of runners to show, one of: specific, shared, active, paused, online; showing all runners if none provided")
	listRunnersCmd.Flags().BoolVar(&runnersAll, "all", false, "Get a list of all runners in the GitLab instance (specific and shared). Access is restricted to users with admin privileges")
}

func fetchRunners() {
	color.Yellow("Fetching runners…")

	o := &gitlab.RunnersOptions{}
	o.Page = page
	o.PerPage = perPage
	if runnersScope != "" {
		o.Scope = gitlab.RunnerScope(runnersScope)
	}
	if runnersAll {
		o.All = true
	}

	loader.Start()
	collection, meta, err := client.Runners(o)
	loader.Stop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(collection.Items) == 0 {
		color.Red("No runner found")
	} else {
		out.Runners(output, outputFormat, collection)
	}

	printMeta(meta, true)

	handlePaginatedResult(meta, fetchProjects)
}

var listRunnersCmd = &cobra.Command{
	Use:     "runners",
	Aliases: []string{"r"},
	Short:   "List runners",
	Run: func(cmd *cobra.Command, args []string) {
		fetchRunners()
	},
}
