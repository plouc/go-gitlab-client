package cmd

import (
	"fmt"

	"github.com/olekukonko/tablewriter"
	"os"
)

func aliasesOutput(aliases []*Alias) {
	fmt.Println("")
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"Alias",
		"Resource type",
		"Resource id",
	})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	for _, alias := range aliases {
		table.Append([]string{
			alias.Alias,
			alias.ResourceType,
			alias.IdsString(),
		})
	}
	table.Render()
	fmt.Println("")
}
