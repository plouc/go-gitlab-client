package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/plouc/go-gitlab-client/gitlab"
)

func branchOutput(branch *gitlab.Branch) {
	if outputFormat == "json" {
		jsonOutput(branch)
	} else {
		fmt.Fprintln(output, "")

		fmt.Fprintf(output, "  Name                %s\n", color.YellowString(branch.Name))
		fmt.Fprintf(output, "  Protected           %s\n", color.YellowString("%t", branch.Protected))
		fmt.Fprintf(output, "  Merged              %s\n", color.YellowString("%t", branch.Merged))
		fmt.Fprintf(output, "  DevelopersCanPush   %s\n", color.YellowString("%t", branch.DevelopersCanPush))
		fmt.Fprintf(output, "  DevelopersCanMerge  %s\n", color.YellowString("%t", branch.DevelopersCanMerge))

		fmt.Fprintln(output, "  Commit")
		fmt.Fprintf(output, "    Id                %s\n", color.YellowString(branch.Commit.Id))
		fmt.Fprintf(output, "    Message           %s\n", color.YellowString(branch.Commit.Message))
		fmt.Fprintf(output, "    Tree              %s\n", color.YellowString(branch.Commit.Tree))
		fmt.Fprintf(output, "    AuthoredDateRaw   %s\n", color.YellowString(branch.Commit.AuthoredDateRaw))
		fmt.Fprintf(output, "    CommittedDateRaw  %s\n", color.YellowString(branch.Commit.CommittedDateRaw))
		// @todo Author
		// @todo Committer

		fmt.Fprintln(output, "")
	}
}

func printAccessLevelInfo(accessLevel *gitlab.AccessLevelInfo) {
	fmt.Printf("    Access level              %d\n", accessLevel.AccessLevel)
	fmt.Printf("    Access level description  %s\n", accessLevel.AccessLevelDescription)
	if accessLevel.GroupId != 0 {
		fmt.Printf("    Group id                  %d\n", accessLevel.GroupId)
	}
	if accessLevel.UserId != 0 {
		fmt.Printf("    User id                   %d\n", accessLevel.UserId)
	}
}

func printProtectedBranch(protectedBranch *gitlab.ProtectedBranch) {
	color.Blue(protectedBranch.Name)

	fmt.Println("  Push access levels:")
	for _, accessLevel := range protectedBranch.PushAccessLevels {
		printAccessLevelInfo(accessLevel)
	}

	fmt.Println("\n  Merge access levels:")
	for _, accessLevel := range protectedBranch.MergeAccessLevels {
		printAccessLevelInfo(accessLevel)
	}
}
