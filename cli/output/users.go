package output

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/plouc/go-gitlab-client/gitlab"
	"io"
)

func Users(w io.Writer, format string, collection *gitlab.UserCollection) {
	if format == "json" {
		collection.RenderJson(w)
	} else if format == "yaml" {
		collection.RenderYaml(w)
	} else {
		fmt.Fprintln(w, "")
		table := tablewriter.NewWriter(w)
		table.SetHeader([]string{
			"Id",
			"Username",
			"Email",
			"Name",
			"State",
			"Is admin",
			"External",
		})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		for _, user := range collection.Items {
			table.Append([]string{
				fmt.Sprintf("%d", user.Id),
				user.Username,
				user.Email,
				user.Name,
				user.State,
				fmt.Sprintf("%t", user.IsAdmin),
				fmt.Sprintf("%t", user.External),
			})
		}
		table.Render()
		fmt.Fprintln(w, "")
	}
}

func User(w io.Writer, format string, user *gitlab.User) {
	if format == "json" {
		user.RenderJson(w)
	} else if format == "yaml" {
		user.RenderYaml(w)
	} else {
		fmt.Fprintln(w, "")

		fmt.Fprintln(w, "Basic info")
		fmt.Fprintf(w, "  Id        %s\n", color.YellowString("%d", user.Id))
		fmt.Fprintf(w, "  Username  %s\n", color.YellowString(user.Username))
		fmt.Fprintf(w, "  Email     %s\n", color.YellowString(user.Email))
		fmt.Fprintf(w, "  Name      %s\n", color.YellowString(user.Name))
		fmt.Fprintf(w, "  State     %s\n", color.YellowString(user.State))

		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "Rights")
		fmt.Fprintf(w, "  IsAdmin           %s\n", color.YellowString("%t", user.IsAdmin))
		fmt.Fprintf(w, "  CanCreateGroup    %s\n", color.YellowString("%t", user.CanCreateGroup))
		fmt.Fprintf(w, "  CanCreateProject  %s\n", color.YellowString("%t", user.CanCreateProject))
		fmt.Fprintf(w, "  ProjectsLimit     %s\n", color.YellowString("%d", user.ProjectsLimit))

		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "Security")
		fmt.Fprintf(w, "  External           %s\n", color.YellowString("%t", user.External))
		fmt.Fprintf(w, "  TwoFactorEnabled    %s\n", color.YellowString("%t", user.TwoFactorEnabled))
		fmt.Fprintln(w, "  Identities")
		for _, identity := range user.Identities {
			fmt.Fprintf(w, "  - Provider   %s\n", color.YellowString(identity.Provider))
			fmt.Fprintf(w, "    ExternUid  %s\n", color.YellowString(identity.ExternUid))
		}

		fmt.Fprintln(w, "")
	}
}
