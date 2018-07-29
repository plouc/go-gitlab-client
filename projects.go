package gogitlab

import (
	"encoding/json"
	"strconv"
)

const (
	projects_url         = "/projects?page=:page&per_page=:per_page"     // Get a list of projects owned by the authenticated user
	projects_all         = "/projects/all?page=:page&per_page=:per_page" // Get a list of all GitLab projects (admin only)
	projects_search_url  = "/projects/search/:query"                     // Search for projects by name
	project_url          = "/projects/:id"                               // Get a specific project, identified by project ID or NAME
	project_url_events   = "/projects/:id/events"                        // Get project events
	project_url_branches = "/projects/:id/repository/branches"           // Lists all branches of a project
	project_url_members  = "/projects/:id/members"                       // List project team members
	project_url_member   = "/projects/:id/members/:user_id"              // Get project team member
	max_page_size        = 20
)

type Member struct {
	Id        int
	Username  string
	Email     string
	Name      string
	State     string
	CreatedAt string `json:"created_at,omitempty"`
	// AccessLevel int
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

type Visibility string

const (
	// VisibilityPrivate indicates project access must be granted explicitly for each user.
	VisibilityPrivate = Visibility("private")

	// VisibilityInternal indicates the project can be cloned by any logged in user.
	VisibilityInternal = Visibility("internal")

	// VisibilityPublic indicates the project can be cloned without any authentication.
	VisibilityPublic = Visibility("public")
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
	Visibility           Visibility `json:"visibility,omitempty"`
	IssuesEnabled        bool       `json:"issues_enabled,omitempty"`
	MergeRequestsEnabled bool       `json:"merge_requests_enabled,omitempty"`
	WallEnabled          bool       `json:"wall_enabled,omitempty"`
	WikiEnabled          bool       `json:"wiki_enabled,omitempty"`
	CreatedAtRaw         string     `json:"created_at,omitempty"`
	Namespace            *Namespace `json:"namespace,omitempty"`
	NamespaceId          int        `json:"namespace_id,omitempty"` // Only used for create
	SshRepoUrl           string     `json:"ssh_url_to_repo"`
	HttpRepoUrl          string     `json:"http_url_to_repo"`
	WebUrl               string     `json:"web_url"`
	SharedRunners        bool       `json:"shared_runners_enabled"`
}

func append(slice []*Project, elements ...*Project) []*Project {
	n := len(slice)
	total := len(slice) + len(elements)
	if total > cap(slice) {
		// Reallocate. Grow to 1.5 times the new size, so we can still grow.
		newSize := total*3/2 + 1
		newSlice := make([]*Project, total, newSize)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[:total]
	copy(slice[n:], elements)
	return slice
}

func projects(u string, g *Gitlab) ([]*Project, error) {

	var projects []*Project
	var page int64 = 1
	for {
		url := g.ResourceUrl(
			u,
			map[string]string{
				":page":     strconv.FormatInt(page, 10),
				":per_page": strconv.FormatInt(max_page_size, 10)})

		var pageProjects []*Project

		contents, err := g.buildAndExecRequest("GET", url, nil)
		if err == nil {
			err = json.Unmarshal(contents, &pageProjects)
			if err == nil {
				projects = append(projects, pageProjects...)
				page++
				if len(pageProjects) < max_page_size {
					break
				}
				continue
			}
		}
		return projects, err
	}

	return projects, nil
}

/*
Get a list of projects owned by the authenticated user.
*/
func (g *Gitlab) Projects() ([]*Project, error) {
	return projects(projects_url, g)
}

/*
Get a list of all GitLab projects (admin only).
*/
func (g *Gitlab) AllProjects() ([]*Project, error) {
	return projects(projects_all, g)
}

/*
Creates a new project owned by the authenticated user.

One (or more) of the following fields are required:
	* Name
	* Path
*/
func (g *Gitlab) AddProject(project *Project) (*Project, error) {
	url := g.ResourceUrl(projects_url, nil)

	encodedRequest, err := json.Marshal(project)
	if err != nil {
		return nil, err
	}
	var result *Project
	contents, err := g.buildAndExecRequest("POST", url, encodedRequest)
	if err == nil {
		err = json.Unmarshal(contents, &result)
	}

	return result, err
}

/*
Remove a project.
*/
func (g *Gitlab) RemoveProject(id string) (bool, error) {

	url, opaque := g.ResourceUrlRaw(project_url, map[string]string{":id": id})
	result := false

	contents, err := g.buildAndExecRequestRaw("DELETE", url, opaque, nil)
	if err == nil {
		result, err = strconv.ParseBool(string(contents[:]))
	}

	return result, err
}

/*
Get a specific project, identified by project ID or NAME,
which is owned by the authentication user.
Namespaced project may be retrieved by specifying the namespace
and its project name like this:

	`namespace%2Fproject-name`

*/
func (g *Gitlab) Project(id string) (*Project, error) {

	url, opaque := g.ResourceUrlRaw(project_url, map[string]string{":id": id})

	var project *Project

	contents, err := g.buildAndExecRequestRaw("GET", url, opaque, nil)
	if err == nil {
		err = json.Unmarshal(contents, &project)
	}

	return project, err
}

/*
Update a specific project, identified by project ID or NAME,
which is owned by the authentication user.
Namespaced project may be retrieved by specifying the namespace
and its project name like this:

	`namespace%2Fproject-name`

*/
func (g *Gitlab) UpdateProject(id string, project *Project) (*Project, error) {

	url := g.ResourceUrl(project_url, map[string]string{":id": id})

	encodedRequest, err := json.Marshal(project)
	if err != nil {
		return nil, err
	}
	var result *Project

	contents, err := g.buildAndExecRequest("PUT", url, encodedRequest)
	if err == nil {
		err = json.Unmarshal(contents, &result)
	}

	return result, err
}

/*
Lists all branches of a project.
*/
func (g *Gitlab) ProjectBranches(id string) ([]*Branch, error) {

	url, opaque := g.ResourceUrlRaw(project_url_branches, map[string]string{":id": id})

	var branches []*Branch

	contents, err := g.buildAndExecRequestRaw("GET", url, opaque, nil)
	if err == nil {
		err = json.Unmarshal(contents, &branches)
	}

	return branches, err
}

func (g *Gitlab) ProjectMembers(id string) ([]*Member, error) {
	url, opaque := g.ResourceUrlRaw(project_url_members, map[string]string{":id": id})

	var members []*Member

	contents, err := g.buildAndExecRequestRaw("GET", url, opaque, nil)
	if err == nil {
		err = json.Unmarshal(contents, &members)
	}

	return members, err
}
