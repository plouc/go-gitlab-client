package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/plouc/go-gitlab-client/gogitlab"
)

func usersOutput(users []*gogitlab.User) {
	if outputFormat == "json" {
		jsonOutput(users)
	} else if outputFormat == "yaml" {
		yamlOutput(users)
	} else {
		fmt.Fprintln(output, "")
		table := tablewriter.NewWriter(output)
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
		for _, user := range users {
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
		fmt.Fprintln(output, "")
	}
}

func userOutput(user *gogitlab.User) {
	if outputFormat == "json" {
		jsonOutput(user)
	} else if outputFormat == "yaml" {
		yamlOutput(user)
	} else {
		fmt.Fprintln(output, "")

		fmt.Fprintln(output, "Basic info")
		fmt.Fprintf(output, "  Id        %s\n", color.YellowString("%d", user.Id))
		fmt.Fprintf(output, "  Username  %s\n", color.YellowString(user.Username))
		fmt.Fprintf(output, "  Email     %s\n", color.YellowString(user.Email))
		fmt.Fprintf(output, "  Name      %s\n", color.YellowString(user.Name))
		fmt.Fprintf(output, "  State     %s\n", color.YellowString(user.State))

		fmt.Fprintln(output, "")
		fmt.Fprintln(output, "Rights")
		fmt.Fprintf(output, "  IsAdmin           %s\n", color.YellowString("%t", user.IsAdmin))
		fmt.Fprintf(output, "  CanCreateGroup    %s\n", color.YellowString("%t", user.CanCreateGroup))
		fmt.Fprintf(output, "  CanCreateProject  %s\n", color.YellowString("%t", user.CanCreateProject))
		fmt.Fprintf(output, "  ProjectsLimit     %s\n", color.YellowString("%d", user.ProjectsLimit))

		fmt.Fprintln(output, "")
		fmt.Fprintln(output, "Security")
		fmt.Fprintf(output, "  External           %s\n", color.YellowString("%t", user.External))
		fmt.Fprintf(output, "  TwoFactorEnabled    %s\n", color.YellowString("%t", user.TwoFactorEnabled))
		fmt.Fprintln(output, "  Identities")
		for _, identity := range user.Identities {
			fmt.Fprintf(output, "  - Provider   %s\n", color.YellowString(identity.Provider))
			fmt.Fprintf(output, "    ExternUid  %s\n", color.YellowString(identity.ExternUid))
		}

		fmt.Fprintln(output, "")
	}
}
