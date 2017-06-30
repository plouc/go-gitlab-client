package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/plouc/go-gitlab-client"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Host    string `json:"host"`
	ApiPath string `json:"api_path"`
	Token   string `json:"token"`
}

func buildProject(name string, namespaceId string, desc string, builds string, mergeRequests string) gogitlab.Project {
	project := gogitlab.Project{
		Description: desc,
	}

	if name != "" {
		project.Name = name
	}
	if namespaceId != "" {
		project.NamespaceId, _ = strconv.Atoi(namespaceId)
	}
	if builds != "" {
		project.BuildsEnabled, _ = strconv.ParseBool(builds)
	}
	if mergeRequests != "" {
		project.MergeRequestsEnabled, _ = strconv.ParseBool(mergeRequests)
	}

	return project
}

func printProject(project gogitlab.Project) {
	format := "> %-23s: %s\n"

	fmt.Printf("%s\n", project.Name)
	fmt.Printf(format, "id", strconv.Itoa(project.Id))
	fmt.Printf(format, "name", project.Name)
	fmt.Printf(format, "description", project.Description)
	fmt.Printf(format, "default branch", project.DefaultBranch)
	if project.Owner != nil {
		fmt.Printf(format, "owner.name", project.Owner.Username)
	}
	fmt.Printf(format, "public", strconv.FormatBool(project.Public))
	fmt.Printf(format, "path", project.Path)
	fmt.Printf(format, "path with namespace", project.PathWithNamespace)
	fmt.Printf(format, "issues enabled", strconv.FormatBool(project.IssuesEnabled))
	fmt.Printf(format, "merge requests enabled", strconv.FormatBool(project.MergeRequestsEnabled))
	fmt.Printf(format, "builds enabled", strconv.FormatBool(project.BuildsEnabled))
	fmt.Printf(format, "wall enabled", strconv.FormatBool(project.WallEnabled))
	fmt.Printf(format, "wiki enabled", strconv.FormatBool(project.WikiEnabled))
	fmt.Printf(format, "shared runners enabled", strconv.FormatBool(project.SharedRunners))
	fmt.Printf(format, "created at", project.CreatedAtRaw)
	//fmt.Printf(format, "namespace",           project.Namespace)
}

func main() {
	help := flag.Bool("help", false, "Show usage")

	file, e := ioutil.ReadFile("../config.json")
	if e != nil {
		fmt.Printf("Config file error: %v\n", e)
		os.Exit(1)
	}

	var config Config
	json.Unmarshal(file, &config)
	fmt.Printf("Results: %+v\n", config)

	gitlab := gogitlab.NewGitlab(config.Host, config.ApiPath, config.Token)

	var method string
	flag.StringVar(&method, "m", "", "Specify method to retrieve projects infos, available methods:\n"+
		"  > -m projects\n"+
		"  > -m project        -id PROJECT_ID [-o edit -desc PROJECT_DESCRIPTION -builds -merge_requests]\n"+
		"  > -m project        -o create -name PROJECT_NAME [-namespace_id PROJECT_NAMESPACE -desc PROJECT_DESCRIPTION -builds -merge_requests]\n"+
		"  > -m hooks          -id PROJECT_ID\n"+
		"  > -m branches       -id PROJECT_ID\n"+
		"  > -m team           -id PROJECT_ID\n"+
		"  > -m merge_requests -id PROJECT_ID [-state <all|merged|opened|closed>] [-order <created_at|updated_at>] [-sort <asc|desc>]")

	var id string
	flag.StringVar(&id, "id", "", "Specify repository id")

	var state string
	flag.StringVar(&state, "state", "", "Specify merge request state")

	var order string
	flag.StringVar(&order, "order", "", "Specify merge request order")

	var sort string
	flag.StringVar(&sort, "sort", "", "Specify merge request sort")

	var operation string
	flag.StringVar(&operation, "o", "", "Specify operation")

	var name string
	flag.StringVar(&name, "name", "", "Specify name")

	var namespaceId string
	flag.StringVar(&namespaceId, "namespace_id", "", "Specify namepace id")

	var desc string
	flag.StringVar(&desc, "desc", "", "Specify description")

	var builds string
	flag.StringVar(&builds, "builds", "", "Specify whether builds are enabled (default no change)")

	var mergeRequests string
	flag.StringVar(&mergeRequests, "merge_requests", "", "Specify whether merge requests are enabled (default no change)")

	flag.Usage = func() {
		fmt.Printf("Usage:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *help == true || method == "" {
		flag.Usage()
		return
	}

	startedAt := time.Now()
	defer func() {
		fmt.Printf("processed in %v\n", time.Now().Sub(startedAt))
	}()

	switch method {
	case "projects":
		fmt.Println("Fetching projects…")

		projects, err := gitlab.Projects()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		for _, project := range projects {
			fmt.Printf("> %6d | %s\n", project.Id, project.Name)
		}

	case "project":

		if operation == "" {
			fmt.Println("Fetching project…")

			if id == "" {
				flag.Usage()
				return
			}

			project, err := gitlab.Project(id)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			printProject(*project)
		}

		switch operation {
		case "create":
			fmt.Println("Create project...")

			if name == "" {
				flag.Usage()
				return
			}

			project := buildProject(name, namespaceId, desc, builds, mergeRequests)

			result, err := gitlab.CreateProject(&project)

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			printProject(*result)

		case "edit":
			fmt.Println("Edit project...")

			if id == "" {
				flag.Usage()
				return
			}

			project := buildProject(name, namespaceId, desc, builds, mergeRequests)

			result, err := gitlab.UpdateProject(id, &project)

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			printProject(*result)

		}

	case "branches":
		fmt.Println("Fetching project branches…")

		if id == "" {
			flag.Usage()
			return
		}

		branches, err := gitlab.ProjectBranches(id)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		for _, branch := range branches {
			fmt.Printf("> %s\n", branch.Name)
		}

	case "hooks":
		fmt.Println("Fetching project hooks…")

		if id == "" {
			flag.Usage()
			return
		}

		hooks, err := gitlab.ProjectHooks(id)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		for _, hook := range hooks {
			fmt.Printf("> [%d] %s, created on %s\n", hook.Id, hook.Url, hook.CreatedAtRaw)
		}

	case "team":
		fmt.Println("Fetching project team members…")

		if id == "" {
			flag.Usage()
			return
		}

		members, err := gitlab.ProjectMembers(id)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		for _, member := range members {
			fmt.Printf("> [%d] %s (%s) since %s\n", member.Id, member.Username, member.Name, member.CreatedAt)
		}

	case "merge_requests":
		fmt.Println("Fetching project merge requests...")

		if id == "" {
			flag.Usage()
			return
		}

		var params map[string]string = make(map[string]string)
		if state != "" {
			params["state"] = state
		}
		if order != "" {
			params["order_by"] = order
		}
		if sort != "" {
			params["sort"] = sort
		}

		mrs, err := gitlab.ProjectMergeRequests(id, params)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		for _, mr := range mrs {
			fmt.Printf("> [%d] %s (+%d) by %s on %s.\n", mr.Id, mr.Title, mr.Upvotes, mr.Author.Name, mr.CreatedAt)
		}

	}
}
