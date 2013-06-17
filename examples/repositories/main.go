package main

import (
	"flag"
	"fmt"
	"time"
	"github.com/plouc/go-gitlab-client"
)

func main() {
	help := flag.Bool("help", false, "Show usage")

	var method string
	flag.StringVar(&method, "m", "", "Specify method to retrieve repositories, available methods:\n" +
									   "  > branches\n" +
									   "  > branch\n" +
									   "  > tags\n" +
									   "  > commits")

	var id string
	flag.StringVar(&id, "id", "", "Specify repository id")

	var token string
	flag.StringVar(&token, "t", "", "Specify gitlab token")

	flag.Usage = func() {
		fmt.Printf("Usage:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *help == true || method == "" || id == "" {
		flag.Usage()
		return
	}

	gitlab := gogitlab.NewGitlab("https://gitlab.fullsix.com", "/api/v3", "/digiposte/socle-php/commits/develop-v7.atom", token)

	startedAt := time.Now()
	defer func() {
		fmt.Printf("processed in %v\n", time.Now().Sub(startedAt))
	}()

	switch method {
	case "branches":
		branches, err := gitlab.RepoBranches(id)
		if err != nil {
			fmt.Println(err.Error())
		}

		for _, branch := range branches {
			fmt.Printf("> %s\n", branch.Name)
		}
	case "branch":
	case "tags":
		tags, err := gitlab.RepoTags(id)
		if err != nil {
			fmt.Println(err.Error())
		}

		for _, tag := range tags {
			fmt.Printf("> %s\n", tag.Name)
		}
	case "commits":
		commits, err := gitlab.RepoCommits(id)
		if err != nil {
			fmt.Println(err.Error())
		}
	
		for _, commit := range commits {
			fmt.Printf("%s > [%s] %s\n", commit.CreatedAt.Format("Mon 02 Jan 15:04"), commit.Author_Name, commit.Title)
		}
	}
}