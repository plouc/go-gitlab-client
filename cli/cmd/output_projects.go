package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/plouc/go-gitlab-client/gogitlab"
)

func projectsOutput(projects []*gogitlab.Project) {
	if outputFormat == "json" {
		jsonOutput(projects)
	} else if outputFormat == "yaml" {
		yamlOutput(projects)
	} else {
		fmt.Fprintln(output, "")
		table := tablewriter.NewWriter(output)
		table.SetHeader([]string{
			"Id",
			"Name",
			"Public",
			"WebUrl",
		})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		for _, project := range projects {
			table.Append([]string{
				strconv.Itoa(project.Id),
				project.Name,
				fmt.Sprintf("%t", project.Public),
				project.WebUrl,
			})
		}
		table.Render()
		fmt.Fprintln(output, "")
	}
}

func projectOutput(project *gogitlab.Project, withStatistics bool) {
	if outputFormat == "json" {
		jsonOutput(project)
	} else {
		fmt.Fprintln(output, "")

		fmt.Fprintf(output, "  Id                                         %s\n", color.YellowString("%d", project.Id))
		fmt.Fprintf(output, "  Name                                       %s\n", color.YellowString(project.Name))
		fmt.Fprintf(output, "  NameWithNamespace                          %s\n", color.YellowString(project.NameWithNamespace))
		fmt.Fprintf(output, "  Description                                %s\n", color.YellowString(project.Description))
		fmt.Fprintf(output, "  DefaultBranch                              %s\n", color.YellowString(project.DefaultBranch))
		// Owner: (*gogitlab.Member)(<nil>)
		fmt.Fprintf(output, "  Public                                     %s\n", color.YellowString("%t", project.Public))
		fmt.Fprintf(output, "  Path                                       %s\n", color.YellowString(project.Path))
		fmt.Fprintf(output, "  PathWithNamespace                          %s\n", color.YellowString(project.PathWithNamespace))
		fmt.Fprintf(output, "  Visibility                                 %s\n", color.YellowString("%s", project.Visibility))
		fmt.Fprintf(output, "  IssuesEnabled                              %s\n", color.YellowString("%t", project.IssuesEnabled))
		fmt.Fprintf(output, "  OpenIssuesCount                            %s\n", color.YellowString("%d", project.OpenIssuesCount))
		fmt.Fprintf(output, "  MergeRequestsEnabled                       %s\n", color.YellowString("%t", project.MergeRequestsEnabled))
		fmt.Fprintf(output, "  WallEnabled                                %s\n", color.YellowString("%t", project.WallEnabled))
		fmt.Fprintf(output, "  WikiEnabled                                %s\n", color.YellowString("%t", project.WikiEnabled))
		fmt.Fprintf(output, "  CreatedAtRaw                               %s\n", color.YellowString(project.CreatedAtRaw))
		fmt.Fprintf(output, "  SshRepoUrl                                 %s\n", color.YellowString(project.SshRepoUrl))
		fmt.Fprintf(output, "  HttpRepoUrl                                %s\n", color.YellowString(project.HttpRepoUrl))
		fmt.Fprintf(output, "  WebUrl                                     %s\n", color.YellowString(project.WebUrl))
		fmt.Fprintf(output, "  ReadmeUrl                                  %s\n", color.YellowString(project.ReadmeUrl))
		fmt.Fprintf(output, "  SharedRunnersEnabled                       %s\n", color.YellowString("%t", project.SharedRunnersEnabled))
		fmt.Fprintf(output, "  Archived                                   %s\n", color.YellowString("%t", project.Archived))
		fmt.Fprintf(output, "  OnlyAllowMergeIfPipelineSucceeds           %s\n", color.YellowString("%t", project.OnlyAllowMergeIfPipelineSucceeds))
		fmt.Fprintf(output, "  OnlyAllowMergeIfAllDiscussionsAreResolved  %s\n", color.YellowString("%t", project.OnlyAllowMergeIfAllDiscussionsAreResolved))
		fmt.Fprintf(output, "  MergeMethod                                %s\n", color.YellowString(project.MergeMethod))
		fmt.Fprintf(output, "  TagList                                    %s\n", color.YellowString(strings.Join(project.TagList, ", ")))
		fmt.Fprintf(output, "  ForksCount                                 %s\n", color.YellowString("%d", project.ForksCount))
		fmt.Fprintf(output, "  StarCount                                  %s\n", color.YellowString("%d", project.StarCount))

		fmt.Fprintln(output, "  Namespace")
		fmt.Fprintf(output, "    Id           %s\n", color.YellowString("%d", project.Namespace.Id))
		fmt.Fprintf(output, "    Name         %s\n", color.YellowString(project.Namespace.Name))
		fmt.Fprintf(output, "    Path         %s\n", color.YellowString(project.Namespace.Path))
		fmt.Fprintf(output, "    Kind         %s\n", color.YellowString(project.Namespace.Kind))
		fmt.Fprintf(output, "    Description  %s\n", color.YellowString(project.Namespace.Description))
		fmt.Fprintf(output, "    CreatedAt    %s\n", color.YellowString(project.Namespace.CreatedAt))
		fmt.Fprintf(output, "    UpdatedAt    %s\n", color.YellowString(project.Namespace.UpdatedAt))

		fmt.Fprintln(output, "  SharedWithGroups")
		if len(project.SharedWithGroups) > 0 {
			for _, group := range project.SharedWithGroups {
				fmt.Fprintf(output, "    - GroupId           %s\n", color.YellowString("%d", group.GroupId))
				fmt.Fprintf(output, "      GroupName         %s\n", color.YellowString(group.GroupName))
				fmt.Fprintf(output, "      GroupAccessLevel  %s\n", color.YellowString("%d", group.GroupAccessLevel))
			}
		} else {
			fmt.Fprintln(output, "    No shared group")
		}

		if withStatistics {
			fmt.Fprintln(output, "  Statistics")
			fmt.Fprintf(output, "    CommitCount       %s\n", color.YellowString("%d", project.Statistics.CommitCount))
			fmt.Fprintf(output, "    StorageSize       %s\n", color.YellowString("%d", project.Statistics.StorageSize))
			fmt.Fprintf(output, "    RepositorySize    %s\n", color.YellowString("%d", project.Statistics.RepositorySize))
			fmt.Fprintf(output, "    LfsObjectsSize    %s\n", color.YellowString("%d", project.Statistics.LfsObjectsSize))
			fmt.Fprintf(output, "    JobArtifactsSize  %s\n", color.YellowString("%d", project.Statistics.JobArtifactsSize))

		}

		fmt.Fprintln(output, "")
	}
}
