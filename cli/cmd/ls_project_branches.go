package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
)

var branchesSearch string

func init() {
	lsCmd.AddCommand(lsProjectBranchesCmd)

	lsProjectBranchesCmd.Flags().StringVarP(&branchesSearch, "search", "s", "", "Search term")
}

func fetchProjectBranches(projectId string) {
	color.Yellow("Fetching project's branches (id: %s)…", projectId)

	o := &gitlab.BranchesOptions{}
	o.Page = page
	o.PerPage = perPage
	if branchesSearch != "" {
		o.Search = branchesSearch
	}

	loader.Start()
	branches, meta, err := client.ProjectBranches(projectId, o)
	loader.Stop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("")
	if len(branches) == 0 {
		color.Red("No branch found for project %s", projectId)
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{
			"Name",
			"Protected",
			"Merged",
			"Developers Can Push",
			"Developers Can Merge",
		})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		for _, branch := range branches {
			table.Append([]string{
				branch.Name,
				fmt.Sprintf("%t", branch.Protected),
				fmt.Sprintf("%t", branch.Merged),
				fmt.Sprintf("%t", branch.DevelopersCanPush),
				fmt.Sprintf("%t", branch.DevelopersCanMerge),
			})
		}
		table.Render()
	}
	fmt.Println("")

	metaOutput(meta, true)

	handlePaginatedResult(meta, func() {
		fetchProjectBranches(projectId)
	})
}

var lsProjectBranchesCmd = &cobra.Command{
	Use:     "project-branches [project id]",
	Aliases: []string{"pb"},
	Short:   "List project branches",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("you must specify a project id")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fetchProjectBranches(args[0])
	},
}
