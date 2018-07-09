package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
)

var namespacesSearch string

func init() {
	lsCmd.AddCommand(lsNamespacesCmd)

	lsNamespacesCmd.Flags().StringVarP(&namespacesSearch, "search", "s", "", "Returns a list of namespaces the user is authorized to see based on the search criteria")
}

func fetchNamespaces() {
	color.Yellow("Fetching namespaces…")

	o := &gitlab.NamespacesOptions{}
	o.Page = page
	o.PerPage = perPage
	if namespacesSearch != "" {
		o.Search = namespacesSearch
	}

	loader.Start()
	namespaces, meta, err := client.Namespaces(o)
	loader.Stop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(namespaces) == 0 {
		color.Red("No namespace found")
	} else {
		namespacesOutput(namespaces)
	}

	metaOutput(meta, true)

	handlePaginatedResult(meta, fetchNamespaces)
}

var lsNamespacesCmd = &cobra.Command{
	Use:     "namespaces",
	Aliases: []string{"ns"},
	Short:   "List namespaces",
	Run: func(cmd *cobra.Command, args []string) {
		fetchNamespaces()
	},
}
