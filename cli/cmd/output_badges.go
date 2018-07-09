package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/plouc/go-gitlab-client/gitlab"
)

func badgesOutput(badges []*gitlab.Badge) {
	if outputFormat == "json" {
		jsonOutput(badges)
	} else {
		fmt.Fprintln(output, "")
		table := tablewriter.NewWriter(output)
		table.SetHeader([]string{
			"Id",
			"Link Url",
			"Image Url",
		})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		for _, badge := range badges {
			table.Append([]string{
				fmt.Sprintf("%d", badge.Id),
				badge.LinkUrl,
				badge.ImageUrl,
			})
		}
		table.Render()
		fmt.Fprintln(output, "")
	}
}

func badgeOutput(badge *gitlab.Badge) {
	if outputFormat == "json" {
		jsonOutput(badge)
	} else {
		fmt.Fprintln(output, "")
		fmt.Fprintf(output, "  Id                %s\n", color.YellowString("%d", badge.Id))
		fmt.Fprintf(output, "  LinkUrl           %s\n", color.YellowString(badge.LinkUrl))
		fmt.Fprintf(output, "  ImageUrl          %s\n", color.YellowString(badge.ImageUrl))
		fmt.Fprintf(output, "  RenderedLinkUrl   %s\n", color.YellowString(badge.RenderedLinkUrl))
		fmt.Fprintf(output, "  RenderedImageUrl  %s\n", color.YellowString(badge.RenderedImageUrl))
		fmt.Fprintf(output, "  Kind              %s\n", color.YellowString(badge.Kind))
		fmt.Fprintln(output, "")
	}
}
