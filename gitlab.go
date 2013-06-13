// Package github implements a simple client to consume gitlab API.
package gogitlab

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Gitlab struct {
	BaseUrl string
	ApiPath string
	Token   string
	Client  *http.Client
}

type Owner struct {
	Id         int
	Username   string
	Email      string
	Name       string
	State      string
	Created_At string
}

type Namespace struct {
	Id          int
	Name        string
	Path        string
	Description string
	Owner_Id    int
	Created_At  string
	Updated_At  string
}

type Project struct {
	Id                     int
	Name                   string
	description            string
	Default_Branch         string
	Owner                  *Owner
	Public                 bool
	Path                   string
	Path_With_Namespace    string
	Issues_enabled         bool
	Merge_Requests_Enabled bool
	Wall_Enabled           bool
	Wiki_Enabled           bool
	Created_At             string
	Namespace              *Namespace
}

func NewGitlab(baseUrl string, apiPath string, token string) *Gitlab {

	client := &http.Client{}

	return &Gitlab{
		BaseUrl: baseUrl,
		ApiPath: apiPath,
		Token:   token,
		Client:  client,
	}
}

func (g *Gitlab) buildAndExecRequest(method string, url string) []byte {
	
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic("Error while building gitlab request")
	}

	resp, err := g.Client.Do(req)
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}
	//fmt.Println(string(contents))

	return contents
}

func (g *Gitlab) Projects() []*Project {

	url := g.BaseUrl + g.ApiPath + "/projects?private_token=" + g.Token
	contents := g.buildAndExecRequest("GET", url)

	var projects []*Project
	err := json.Unmarshal(contents, &projects)
	if err != nil {
		fmt.Println("%s", err)
	}

	for _, project := range projects {
		fmt.Println("%+v", project)
	}

	return projects
}
