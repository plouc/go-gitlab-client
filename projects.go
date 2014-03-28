package gogitlab

import (
	"encoding/json"
)

const (
	projects_url         = "/projects"                         // Get a list of projects owned by the authenticated user
	projects_search_url  = "/projects/search/:query"           // Search for projects by name
	project_url          = "/projects/:id"                     // Get a specific project, identified by project ID or NAME
	project_url_events   = "/projects/:id/events"              // Get project events
	project_url_branches = "/projects/:id/repository/branches" // Lists all branches of a project
	project_url_members  = "/projects/:id/members"             // List project team members
	project_url_member   = "/projects/:id/members/:user_id"    // Get project team member
)

// A gitlab project
type Project struct {
	Id                   int        `json:"id,omitempty"`
	Name                 string     `json:"name,omitempty"`
	Description          string     `json:"description,omitempty"`
	DefaultBranch        string     `json:"default_branch,omitempty"`
	Owner                *Member    `json:"owner,omitempty"`
	Public               bool       `json:"public,omitempty"`
	Path                 string     `json:"path,omitempty"`
	PathWithNamespace    string     `json:"path_with_namespace,omitempty"`
	IssuesEnabled        bool       `json:"issues_enabled,omitempty"`
	MergeRequestsEnabled bool       `json:"merge_requests_enabled,omitempty"`
	WallEnabled          bool       `json:"wall_enabled,omitempty"`
	WikiEnabled          bool       `json:"wiki_enabled,omitempty"`
	CreatedAtRaw         string     `json:"created_at,omitempty"`
	Namespace            *Namespace `json:"namespace,omitempty"`
	SshRepoUrl           string     `json:"ssh_url_to_repo"`
	HttpRepoUrl          string     `json:"http_url_to_repo"`
}

/*
Get a list of projects owned by the authenticated user.
*/
func (g *Gitlab) Projects() ([]*Project, error) {

	url := g.ResourceUrl(projects_url, nil)

	var projects []*Project

	contents, err := g.buildAndExecRequest("GET", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &projects)
	}

	return projects, err
}

/*
Get a specific project, identified by project ID or NAME,
which is owned by the authentication user.
Currently namespaced projects cannot be retrieved by name.
*/
func (g *Gitlab) Project(id string) (*Project, error) {

	url := g.ResourceUrl(project_url, map[string]string{":id": id})

	var project *Project

	contents, err := g.buildAndExecRequest("GET", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &project)
	}

	return project, err
}

/*
Lists all branches of a project.
*/
func (g *Gitlab) ProjectBranches(id string) ([]*Branch, error) {

	url := g.ResourceUrl(project_url_branches, map[string]string{":id": id})

	var branches []*Branch

	contents, err := g.buildAndExecRequest("GET", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &branches)
	}

	return branches, err
}

func (g *Gitlab) ProjectMembers(id string) ([]*Member, error) {
	url := g.ResourceUrl(project_url_members, map[string]string{":id": id})

	var members []*Member

	contents, err := g.buildAndExecRequest("GET", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &members)
	}

	return members, err
}
