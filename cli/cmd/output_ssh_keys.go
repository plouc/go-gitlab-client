package cmd

import (
	"fmt"

	"github.com/olekukonko/tablewriter"
	"github.com/plouc/go-gitlab-client/gitlab"
)

func sshKeysOutput(keys []*gitlab.SshKey) {
	if outputFormat == "json" {
		jsonOutput(keys)
	} else if outputFormat == "yaml" {
		yamlOutput(keys)
	} else {
		fmt.Fprintln(output, "")
		table := tablewriter.NewWriter(output)
		table.SetHeader([]string{
			"Id",
			"Title",
			"Key",
			"Created at",
		})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		for _, key := range keys {
			table.Append([]string{
				fmt.Sprintf("%d", key.Id),
				key.Title,
				key.Key[:16] + "…",
				key.CreatedAtRaw,
			})
		}
		table.Render()
		fmt.Fprintln(output, "")
	}
}
