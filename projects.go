package gogitlab

import (
	"fmt"
	"strings"
	"encoding/json"
)

const (
	projects_url         = "/projects"                         // Get a list of projects owned by the authenticated user
	project_url          = "/projects/:id"                     // Get a specific project, identified by project ID or NAME
	project_url_events   = "/projects/:id/events"              // Get project events
	project_url_branches = "/projects/:id/repository/branches" // Lists all branches of a project
	project_url_hooks    = "/projects/:id/hooks"               // Get list of project hooks
)

// A gitlab project
type Project struct {
	Id                   int        `json:"id,omitempty"`
	Name                 string     `json:"name,omitempty"`
	Description          string     `json:"description,omitempty"`
	DefaultBranch        string     `json:"default_branch,omitempty"`
	Owner                *Owner     `json:"owner,omitempty"`
	Public               bool       `json:"public,omitempty"`
	Path                 string     `json:"path,omitempty"`
	PathWithNamespace    string     `json:"path_with_namespace,omitempty"`
	IssuesEnabled        bool       `json:"issues_enabled,omitempty"`
	MergeRequestsEnabled bool       `json:"merge_requests_enabled,omitempty"`
	WallEnabled          bool       `json:"wall_enabled,omitempty"`
	WikiEnabled          bool       `json:"wiki_enabled,omitempty"`
	CreatedAtRaw         string     `json:"created_at,omitempty"`
	Namespace            *Namespace `json:"namespace,omitempty"`
}

/*
Get a list of projects owned by the authenticated user.
*/
func (g *Gitlab) Projects() ([]*Project, error) {
	
	url := g.BaseUrl + g.ApiPath + projects_url + "?private_token=" + g.Token

	contents, err := g.buildAndExecRequest("GET", url)
	if err != nil {
		fmt.Println("%s", err)
	}

	var projects []*Project
	err = json.Unmarshal(contents, &projects)
	if err != nil {
		fmt.Println("%s", err)
	}

	return projects, err
}

/*
Get a specific project, identified by project ID or NAME,
which is owned by the authentication user.
Currently namespaced projects cannot be retrieved by name.
*/
func (g *Gitlab) Project(id string) (*Project, error) {

	url := strings.Replace(project_url, ":id", id, -1)
	url = g.BaseUrl + g.ApiPath + url + "?private_token=" + g.Token

	var err error
	var project *Project

	contents, err := g.buildAndExecRequest("GET", url)
		
	if err == nil {
		err = json.Unmarshal(contents, &project)
	}

	return project, err
}

/*
Lists all branches of a project.
*/
func (g *Gitlab) ProjectBranches(id string) ([]*Branch, error) {

	url := strings.Replace(project_url_branches, ":id", id, -1)
	url = g.BaseUrl + g.ApiPath + url + "?private_token=" + g.Token

	var err error
	var branches []*Branch

	contents, err := g.buildAndExecRequest("GET", url)
	if err != nil {
		return branches, err
	}
		
	err = json.Unmarshal(contents, &branches)

	return branches, err
}

/*
Get list of project hooks.
*/
func (g *Gitlab) ProjectHooks(id string) ([]*Hook, error) {

	url := strings.Replace(project_url_hooks, ":id", id, -1)
	url = g.BaseUrl + g.ApiPath + url + "?private_token=" + g.Token

	var err error
	var hooks []*Hook

	contents, err := g.buildAndExecRequest("GET", url)
	if err != nil {
		return hooks, err
	}
		
	err = json.Unmarshal(contents, &hooks)

	return hooks, err
}
