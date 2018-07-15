package output

import (
	"fmt"
	"io"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/plouc/go-gitlab-client/gitlab"
)

func Badges(w io.Writer, format string, badges []*gitlab.Badge) {
	if format == "json" {
		Json(w, badges)
	} else if format == "yaml" {
		Yaml(w, badges)
	} else {
		fmt.Fprintln(w, "")
		table := tablewriter.NewWriter(w)
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
		fmt.Fprintln(w, "")
	}
}

func Badge(w io.Writer, format string, badge *gitlab.Badge) {
	if format == "json" {
		Json(w, badge)
	} else if format == "yaml" {
		Yaml(w, badge)
	} else {
		fmt.Fprintln(w, "")
		fmt.Fprintf(w, "  Id                %s\n", color.YellowString("%d", badge.Id))
		fmt.Fprintf(w, "  LinkUrl           %s\n", color.YellowString(badge.LinkUrl))
		fmt.Fprintf(w, "  ImageUrl          %s\n", color.YellowString(badge.ImageUrl))
		fmt.Fprintf(w, "  RenderedLinkUrl   %s\n", color.YellowString(badge.RenderedLinkUrl))
		fmt.Fprintf(w, "  RenderedImageUrl  %s\n", color.YellowString(badge.RenderedImageUrl))
		fmt.Fprintf(w, "  Kind              %s\n", color.YellowString(badge.Kind))
		fmt.Fprintln(w, "")
	}
}
