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
		"  > -m project        -id PROJECT_ID [-o <edit> -desc PROJECT_DESCRIPTION]\n"+
		"  > -m hooks          -id PROJECT_ID\n"+
		"  > -m branches       -id PROJECT_ID\n"+
		"  > -m team           -id PROJECT_ID\n"+
		"  > -m merge_requests -id PROJECT_ID [-state <all|merged|opened|closed>] [-order <created_at|updated_at>] [-sort <asc|desc>]\n"+
		"  > -m variables      -id PROJECT_ID [-o <add|get|edit|rm>] [-key VARIABLE_KEY] [-value VARIABLE_VALUE]")

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

	var key string
	flag.StringVar(&key, "key", "", "Specify key")

	var value string
	flag.StringVar(&value, "value", "", "Specify value")

	var desc string
	flag.StringVar(&desc, "desc", "", "Specify description")

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

		if id == "" {
			flag.Usage()
			return
		}

		if operation == "" {
			fmt.Println("Fetching project…")

			project, err := gitlab.Project(id)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			format := "> %-23s: %s\n"

			fmt.Printf("%s\n", project.Name)
			fmt.Printf(format, "id", strconv.Itoa(project.Id))
			fmt.Printf(format, "name", project.Name)
			fmt.Printf(format, "description", project.Description)
			fmt.Printf(format, "default branch", project.DefaultBranch)
			fmt.Printf(format, "owner.name", project.Owner.Username)
			fmt.Printf(format, "public", strconv.FormatBool(project.Public))
			fmt.Printf(format, "path", project.Path)
			fmt.Printf(format, "path with namespace", project.PathWithNamespace)
			fmt.Printf(format, "issues enabled", strconv.FormatBool(project.IssuesEnabled))
			fmt.Printf(format, "merge requests enabled", strconv.FormatBool(project.MergeRequestsEnabled))
			fmt.Printf(format, "wall enabled", strconv.FormatBool(project.WallEnabled))
			fmt.Printf(format, "wiki enabled", strconv.FormatBool(project.WikiEnabled))
			fmt.Printf(format, "shared runners enabled", strconv.FormatBool(project.SharedRunners))
			fmt.Printf(format, "created at", project.CreatedAtRaw)
			//fmt.Printf(format, "namespace",           project.Namespace)
		}

		switch operation {
		case "edit":
			fmt.Println("Edit project...")

			project := gogitlab.Project{
				Description: desc,
			}

			_, err := gitlab.UpdateProject(id, &project)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
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

	case "variables":

		if id == "" {
			flag.Usage()
			return
		}

		if operation == "" {
			fmt.Println("Fetching project variables...")

			variables, err := gitlab.ProjectVariables(id)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			for _, variable := range variables {
				fmt.Printf("> %s -> %s.\n", variable.Key, variable.Value)
			}
			return
		}

		switch operation {
		case "get":
			fmt.Println("Fetching project variable...")
			if key == "" {
				flag.Usage()
				return
			}

			variable, err := gitlab.ProjectVariable(id, key)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			fmt.Printf("> %s -> %s.\n", variable.Key, variable.Value)

		case "add":
			fmt.Println("Add project variable...")
			if key == "" || value == "" {
				flag.Usage()
				return
			}

			req := gogitlab.Variable{
				Key:   key,
				Value: value,
			}

			variable, err := gitlab.AddProjectVariable(id, &req)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			fmt.Printf("> %s -> %s.\n", variable.Key, variable.Value)

		case "edit":
			fmt.Println("Edit project variable...")
			if key == "" || value == "" {
				flag.Usage()
				return
			}

			req := gogitlab.Variable{
				Key:   key,
				Value: value,
			}

			variable, err := gitlab.UpdateProjectVariable(id, &req)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			fmt.Printf("> %s -> %s.\n", variable.Key, variable.Value)

		case "rm":
			fmt.Println("Delete project variable...")
			if key == "" {
				flag.Usage()
				return
			}

			variable, err := gitlab.DeleteProjectVariable(id, key)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			fmt.Printf("> %s -> %s.\n", variable.Key, variable.Value)

		}
	}
}
