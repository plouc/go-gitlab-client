package output

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/plouc/go-gitlab-client/gitlab"
	"io"
)

func Projects(w io.Writer, format string, collection *gitlab.ProjectCollection) {
	if format == "json" {
		collection.RenderJson(w)
	} else if format == "yaml" {
		collection.RenderYaml(w)
	} else {
		fmt.Fprintln(w, "")
		table := tablewriter.NewWriter(w)
		table.SetHeader([]string{
			"Id",
			"Name",
			"Public",
			"WebUrl",
		})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		for _, project := range collection.Items {
			table.Append([]string{
				strconv.Itoa(project.Id),
				project.Name,
				fmt.Sprintf("%t", project.Public),
				project.WebUrl,
			})
		}
		table.Render()
		fmt.Fprintln(w, "")
	}
}

func Project(w io.Writer, format string, project *gitlab.Project, withStatistics bool) {
	if format == "json" {
		project.RenderJson(w)
	} else if format == "yaml" {
		project.RenderYaml(w)
	} else {
		fmt.Fprintln(w, "")

		fmt.Fprintf(w, "  Id                                         %s\n", color.YellowString("%d", project.Id))
		fmt.Fprintf(w, "  Name                                       %s\n", color.YellowString(project.Name))
		fmt.Fprintf(w, "  NameWithNamespace                          %s\n", color.YellowString(project.NameWithNamespace))
		fmt.Fprintf(w, "  Description                                %s\n", color.YellowString(project.Description))
		fmt.Fprintf(w, "  DefaultBranch                              %s\n", color.YellowString(project.DefaultBranch))
		// Owner: (*gitlab.Member)(<nil>)
		fmt.Fprintf(w, "  Public                                     %s\n", color.YellowString("%t", project.Public))
		fmt.Fprintf(w, "  Path                                       %s\n", color.YellowString(project.Path))
		fmt.Fprintf(w, "  PathWithNamespace                          %s\n", color.YellowString(project.PathWithNamespace))
		fmt.Fprintf(w, "  Visibility                                 %s\n", color.YellowString("%s", project.Visibility))
		fmt.Fprintf(w, "  IssuesEnabled                              %s\n", color.YellowString("%t", project.IssuesEnabled))
		fmt.Fprintf(w, "  OpenIssuesCount                            %s\n", color.YellowString("%d", project.OpenIssuesCount))
		fmt.Fprintf(w, "  MergeRequestsEnabled                       %s\n", color.YellowString("%t", project.MergeRequestsEnabled))
		fmt.Fprintf(w, "  WallEnabled                                %s\n", color.YellowString("%t", project.WallEnabled))
		fmt.Fprintf(w, "  WikiEnabled                                %s\n", color.YellowString("%t", project.WikiEnabled))
		fmt.Fprintf(w, "  CreatedAtRaw                               %s\n", color.YellowString(project.CreatedAtRaw))
		fmt.Fprintf(w, "  SshRepoUrl                                 %s\n", color.YellowString(project.SshRepoUrl))
		fmt.Fprintf(w, "  HttpRepoUrl                                %s\n", color.YellowString(project.HttpRepoUrl))
		fmt.Fprintf(w, "  WebUrl                                     %s\n", color.YellowString(project.WebUrl))
		fmt.Fprintf(w, "  ReadmeUrl                                  %s\n", color.YellowString(project.ReadmeUrl))
		fmt.Fprintf(w, "  SharedRunnersEnabled                       %s\n", color.YellowString("%t", project.SharedRunnersEnabled))
		fmt.Fprintf(w, "  Archived                                   %s\n", color.YellowString("%t", project.Archived))
		fmt.Fprintf(w, "  OnlyAllowMergeIfPipelineSucceeds           %s\n", color.YellowString("%t", project.OnlyAllowMergeIfPipelineSucceeds))
		fmt.Fprintf(w, "  OnlyAllowMergeIfAllDiscussionsAreResolved  %s\n", color.YellowString("%t", project.OnlyAllowMergeIfAllDiscussionsAreResolved))
		fmt.Fprintf(w, "  MergeMethod                                %s\n", color.YellowString(project.MergeMethod))
		fmt.Fprintf(w, "  TagList                                    %s\n", color.YellowString(strings.Join(project.TagList, ", ")))
		fmt.Fprintf(w, "  ForksCount                                 %s\n", color.YellowString("%d", project.ForksCount))
		fmt.Fprintf(w, "  StarCount                                  %s\n", color.YellowString("%d", project.StarCount))

		fmt.Fprintln(w, "  Namespace")
		fmt.Fprintf(w, "    Id           %s\n", color.YellowString("%d", project.Namespace.Id))
		fmt.Fprintf(w, "    Name         %s\n", color.YellowString(project.Namespace.Name))
		fmt.Fprintf(w, "    Path         %s\n", color.YellowString(project.Namespace.Path))
		fmt.Fprintf(w, "    Kind         %s\n", color.YellowString(project.Namespace.Kind))
		fmt.Fprintf(w, "    Description  %s\n", color.YellowString(project.Namespace.Description))
		fmt.Fprintf(w, "    CreatedAt    %s\n", color.YellowString(project.Namespace.CreatedAt))
		fmt.Fprintf(w, "    UpdatedAt    %s\n", color.YellowString(project.Namespace.UpdatedAt))

		fmt.Fprintln(w, "  SharedWithGroups")
		if len(project.SharedWithGroups) > 0 {
			for _, group := range project.SharedWithGroups {
				fmt.Fprintf(w, "    - GroupId           %s\n", color.YellowString("%d", group.GroupId))
				fmt.Fprintf(w, "      GroupName         %s\n", color.YellowString(group.GroupName))
				fmt.Fprintf(w, "      GroupAccessLevel  %s\n", color.YellowString("%d", group.GroupAccessLevel))
			}
		} else {
			fmt.Fprintln(w, "    No shared group")
		}

		if withStatistics {
			fmt.Fprintln(w, "  Statistics")
			fmt.Fprintf(w, "    CommitCount       %s\n", color.YellowString("%d", project.Statistics.CommitCount))
			fmt.Fprintf(w, "    StorageSize       %s\n", color.YellowString("%d", project.Statistics.StorageSize))
			fmt.Fprintf(w, "    RepositorySize    %s\n", color.YellowString("%d", project.Statistics.RepositorySize))
			fmt.Fprintf(w, "    LfsObjectsSize    %s\n", color.YellowString("%d", project.Statistics.LfsObjectsSize))
			fmt.Fprintf(w, "    JobArtifactsSize  %s\n", color.YellowString("%d", project.Statistics.JobArtifactsSize))

		}

		fmt.Fprintln(w, "")
	}
}
