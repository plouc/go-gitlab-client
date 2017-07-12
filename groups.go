package gogitlab

import (
	"encoding/json"
	"strconv"
)

const (
	groups_url         = "/groups?page=:page&per_page=:per_page" // Get a list of groups. (As user: my groups or all available, as admin: all groups)
	groups_add_url     = "/groups"                               // POST to add a group
	group_url          = "/groups/:id"                           // Get all details of a group
	group_projects_url = "/groups/:id/projects"                  // Get a list of projects in this group
	group_url_members  = "/groups/:id/members"                   // Get a list of members in this group
)

// A gitlab group
type Group struct {
	Id                        int        `json:"id,omitempty"`
	Name                      string     `json:"name,omitempty"`
	Path                      string     `json:"path,omitempty"`
	Description               string     `json:"description,omitempty"`
	Visibility                Visibility `json:"visibility,omitempty"`
	LfsEnabled                bool       `json:"lfs_enabled,omitempty"`
	AvatarUrl                 string     `json:"avatar_url,omitempty"`
	WebURL                    string     `json:"web_url,omitempty"`
	RequestAccessEnabled      bool       `json:"request_access_enabled,omitempty"`
	FullName                  string     `json:"full_name,omitempty"`
	FullPath                  string     `json:"full_path,omitempty"`
	ParentId                  int        `json:"parent_id,omitempty"`
	SharedRunnersMinutesLimit int        `json:"shared_runners_minutes_limit,omitempty"`
	Projects                  []*Project `json:"projects,omitempty"`
}

/*
Get a list of groups. (As user: my groups or all available, as admin: all groups)
*/
func (g *Gitlab) Groups(pageNum, resPerPage int) ([]*Group, error) {
	url := g.ResourceUrl(groups_url, map[string]string{":page": strconv.Itoa(pageNum), ":per_page": strconv.Itoa(resPerPage)})
	var groups []*Group

	contents, err := g.buildAndExecRequest("GET", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &groups)
	}

	return groups, err
}

/*
Get all details of a group
*/
func (g *Gitlab) Group(id string) (*Group, error) {
	url, opaque := g.ResourceUrlRaw(group_url, map[string]string{":id": id})

	var group *Group

	contents, err := g.buildAndExecRequestRaw("GET", url, opaque, nil)
	if err == nil {
		err = json.Unmarshal(contents, &group)
	}

	return group, err
}

/*
Creates a new project group. Available only for users who can create groups.

Required fields on group:
	* Name
	* Path

Optional fields on group:
	* Description
	* Visibility
	* LfsEnabled
	* RequestAccessEnabled
	* ParentId

Other fields on group are not supported by the GitLab API
*/
func (g *Gitlab) AddGroup(group *Group) (*Group, error) {
	url := g.ResourceUrl(groups_add_url, nil)

	encodedRequest, err := json.Marshal(group)
	if err != nil {
		return nil, err
	}
	var result *Group
	contents, err := g.buildAndExecRequest("POST", url, encodedRequest)
	if err == nil {
		err = json.Unmarshal(contents, &result)
	}

	return result, err
}

/*
Updates the project group. Only available to group owners and administrators.
*/
func (g *Gitlab) UpdateGroup(id string, group *Group) (*Group, error) {
	url := g.ResourceUrl(group_url, map[string]string{":id": id})

	encodedRequest, err := json.Marshal(group)
	if err != nil {
		return nil, err
	}
	var result *Group

	contents, err := g.buildAndExecRequest("PUT", url, encodedRequest)
	if err == nil {
		err = json.Unmarshal(contents, &result)
	}

	return result, err
}

/*
Remove a group.
*/
func (g *Gitlab) RemoveGroup(id string) (bool, error) {
	url, opaque := g.ResourceUrlRaw(group_url, map[string]string{":id": id})
	result := false

	contents, err := g.buildAndExecRequestRaw("DELETE", url, opaque, nil)
	if err == nil {
		result, err = strconv.ParseBool(string(contents[:]))
	}

	return result, err
}

/*
Get a list of projects in this group.
*/
func (g *Gitlab) GroupProjects(id string) ([]*Project, error) {
	url, opaque := g.ResourceUrlRaw(group_projects_url, map[string]string{":id": id})

	var projects []*Project

	contents, err := g.buildAndExecRequestRaw("GET", url, opaque, nil)
	if err == nil {
		err = json.Unmarshal(contents, &projects)
	}

	return projects, err
}

/*
Gets a list of group or project members viewable by the authenticated user
*/
func (g *Gitlab) GroupMembers(id string) ([]*Member, error) {
	url, opaque := g.ResourceUrlRaw(group_url_members, map[string]string{":id": id})

	var members []*Member

	contents, err := g.buildAndExecRequestRaw("GET", url, opaque, nil)
	if err == nil {
		err = json.Unmarshal(contents, &members)
	}

	return members, err
}
