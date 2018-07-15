package output

import (
	"fmt"
	"io"

	"github.com/fatih/color"
	"github.com/plouc/go-gitlab-client/gitlab"
)

func Branch(w io.Writer, format string, branch *gitlab.Branch) {
	if format == "json" {
		Json(w, branch)
	} else if format == "yaml" {
		Yaml(w, branch)
	} else {
		fmt.Fprintln(w, "")

		fmt.Fprintf(w, "  Name                %s\n", color.YellowString(branch.Name))
		fmt.Fprintf(w, "  Protected           %s\n", color.YellowString("%t", branch.Protected))
		fmt.Fprintf(w, "  Merged              %s\n", color.YellowString("%t", branch.Merged))
		fmt.Fprintf(w, "  DevelopersCanPush   %s\n", color.YellowString("%t", branch.DevelopersCanPush))
		fmt.Fprintf(w, "  DevelopersCanMerge  %s\n", color.YellowString("%t", branch.DevelopersCanMerge))

		fmt.Fprintln(w, "  Commit")
		fmt.Fprintf(w, "    Id                %s\n", color.YellowString(branch.Commit.Id))
		fmt.Fprintf(w, "    Message           %s\n", color.YellowString(branch.Commit.Message))
		fmt.Fprintf(w, "    Tree              %s\n", color.YellowString(branch.Commit.Tree))
		fmt.Fprintf(w, "    AuthoredDateRaw   %s\n", color.YellowString(branch.Commit.AuthoredDateRaw))
		fmt.Fprintf(w, "    CommittedDateRaw  %s\n", color.YellowString(branch.Commit.CommittedDateRaw))
		// @todo Author
		// @todo Committer

		fmt.Fprintln(w, "")
	}
}

func printAccessLevelInfo(w io.Writer, accessLevel *gitlab.AccessLevelInfo) {
	fmt.Fprintf(w, "    Access level              %d\n", accessLevel.AccessLevel)
	fmt.Fprintf(w, "    Access level description  %s\n", accessLevel.AccessLevelDescription)
	if accessLevel.GroupId != 0 {
		fmt.Fprintf(w, "    Group id                  %d\n", accessLevel.GroupId)
	}
	if accessLevel.UserId != 0 {
		fmt.Fprintf(w, "    User id                   %d\n", accessLevel.UserId)
	}
}

func ProtectedBranches(w io.Writer, format string, protectedBranches []*gitlab.ProtectedBranch) {
	if format == "json" {
		Json(w, protectedBranches)
	} else if format == "yaml" {
		Yaml(w, protectedBranches)
	} else {
		for _, protectedBranch := range protectedBranches {
			fmt.Fprintln(w, protectedBranch.Name)

			fmt.Println("  Push access levels:")
			for _, accessLevel := range protectedBranch.PushAccessLevels {
				printAccessLevelInfo(w, accessLevel)
			}

			fmt.Println("\n  Merge access levels:")
			for _, accessLevel := range protectedBranch.MergeAccessLevels {
				printAccessLevelInfo(w, accessLevel)
			}
		}
	}
}
