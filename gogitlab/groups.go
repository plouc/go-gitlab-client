package gogitlab

import (
	"encoding/json"
	"strconv"
	"strings"
)

const (
	groupsUrl        = "/groups"
	groupUrl         = "/groups/:id"
	groupProjectsUrl = "/groups/:id/projects"
)

type Group struct {
	Id                        int        `json:"id,omitempty" yaml:"id,omitempty"`
	Name                      string     `json:"name" yaml:"name"`
	Path                      string     `json:"path" yaml:"path"`
	Description               string     `json:"description" yaml:"description"`
	Visibility                Visibility `json:"visibility" yaml:"visibility"`
	LfsEnabled                bool       `json:"lfs_enabled" yaml:"lfs_enabled"`
	RequestAccessEnabled      bool       `json:"request_access_enabled" yaml:"request_access_enabled"`
	ParentId                  int        `json:"parent_id" yaml:"parent_id"`
	SharedRunnersMinutesLimit int        `json:"shared_runners_minutes_limit" yaml:"shared_runners_minutes_limit"`
	AvatarUrl                 string     `json:"avatar_url" yaml:"avatar_url"`
	WebURL                    string     `json:"web_url" yaml:"web_url"`
	FullName                  string     `json:"full_name" yaml:"full_name"`
	FullPath                  string     `json:"full_path" yaml:"full_path"`
}

type GroupWithDetails struct {
	Group          `yaml:",inline"`
	Projects       []*Project `json:"projects" yaml:"projects"`
	SharedProjects []*Project `json:"shared_projects" yaml:"shared_projects"`
}

type GroupsOptions struct {
	PaginationOptions
	SkipGroups           []string // Skip the group IDs passed
	AllAvailable         bool     //Show all the groups you have access to (defaults to false for authenticated users, true for admin)
	Search               string   // Return the list of authorized groups matching the search criteria
	Statistics           bool     // Include group statistics (admins only)
	WithCustomAttributes bool     // Include custom attributes in response (admins only)
	Owned                bool     // Limit to groups owned by the current user
	// order_by	string	no	Order groups by name, path or id. Default is name
	// sort	string	no	Order groups in asc or desc order. Default is asc
}

func (g *Gitlab) Groups(o *GroupsOptions) ([]*Group, *ResponseMeta, error) {
	u := g.ResourceUrl(groupsUrl, nil)
	if o != nil {
		q := u.Query()

		if o.Page != 1 {
			q.Set("page", strconv.Itoa(o.Page))
		}
		if o.PerPage != 0 {
			q.Set("per_page", strconv.Itoa(o.PerPage))
		}
		if len(o.SkipGroups) > 0 {
			q.Set("skip_groups", strings.Join(o.SkipGroups, ","))
		}
		if o.AllAvailable {
			q.Set("all_available", "true")
		}
		if o.Search != "" {
			q.Set("search", o.Search)
		}
		if o.Statistics {
			q.Set("statistics", "true")
		}
		if o.WithCustomAttributes {
			q.Set("with_custom_attributes", "true")
		}
		if o.Owned {
			q.Set("owned", "true")
		}

		u.RawQuery = q.Encode()
	}

	var groups []*Group

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &groups)
	}

	return groups, meta, err
}

func (g *Gitlab) Group(id string, withCustomAttributes bool) (*GroupWithDetails, *ResponseMeta, error) {
	u := g.ResourceUrl(groupUrl, map[string]string{":id": id})
	q := u.Query()

	if withCustomAttributes {
		q.Set("with_custom_attributes", "true")
	}

	u.RawQuery = q.Encode()

	group := new(GroupWithDetails)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &group)
	}

	return group, meta, err
}

type GroupAddPayload struct {
	Name                      string     `json:"name"`
	Path                      string     `json:"path"`
	Description               string     `json:"description,omitempty"`
	Visibility                Visibility `json:"visibility,omitempty"`
	LfsEnabled                bool       `json:"lfs_enabled,omitempty"`
	RequestAccessEnabled      bool       `json:"request_access_enabled,omitempty"`
	ParentId                  int        `json:"parent_id,omitempty"`
	SharedRunnersMinutesLimit int        `json:"shared_runners_minutes_limit,omitempty"`
}

func (g *Gitlab) AddGroup(group *GroupAddPayload) (*GroupWithDetails, *ResponseMeta, error) {
	u := g.ResourceUrl(groupsUrl, nil)

	encodedRequest, err := json.Marshal(group)
	if err != nil {
		return nil, nil, err
	}

	var createdGroup *GroupWithDetails
	contents, meta, err := g.buildAndExecRequest("POST", u.String(), encodedRequest)
	if err == nil {
		err = json.Unmarshal(contents, &createdGroup)
	}

	return createdGroup, meta, err
}

type GroupUpdatePayload struct {
	Id                        int        `json:"id,omitempty"`
	Name                      string     `json:"name"`
	Path                      string     `json:"path"`
	Description               string     `json:"description,omitempty"`
	MembershipLock            bool       `json:"membership_lock,omitempty"`
	ShareWithGroupLock        bool       `json:"share_with_group_lock,omitempty"`
	Visibility                Visibility `json:"visibility,omitempty"`
	LfsEnabled                bool       `json:"lfs_enabled,omitempty"`
	RequestAccessEnabled      bool       `json:"request_access_enabled,omitempty"`
	SharedRunnersMinutesLimit int        `json:"shared_runners_minutes_limit,omitempty"`
}

func (g *Gitlab) UpdateGroup(id string, group *GroupUpdatePayload) (*GroupWithDetails, error) {
	u := g.ResourceUrl(groupUrl, map[string]string{":id": id})

	encodedRequest, err := json.Marshal(group)
	if err != nil {
		return nil, err
	}

	var updatedGroup *GroupWithDetails
	contents, _, err := g.buildAndExecRequest("PUT", u.String(), encodedRequest)
	if err == nil {
		err = json.Unmarshal(contents, &updatedGroup)
	}

	return updatedGroup, err
}

func (g *Gitlab) RemoveGroup(id string) (string, *ResponseMeta, error) {
	u := g.ResourceUrl(groupUrl, map[string]string{":id": id})

	var responseWithMessage *ResponseWithMessage
	contents, meta, err := g.buildAndExecRequest("DELETE", u.String(), nil)
	if err != nil {
		return "", meta, err
	}

	err = json.Unmarshal(contents, &responseWithMessage)

	return responseWithMessage.Message, meta, err
}

func (g *Gitlab) GroupProjects(id string) ([]*Project, *ResponseMeta, error) {
	url, opaque := g.ResourceUrlRaw(groupProjectsUrl, map[string]string{":id": id})

	var projects []*Project

	contents, meta, err := g.buildAndExecRequestRaw("GET", url, opaque, nil)
	if err == nil {
		err = json.Unmarshal(contents, &projects)
	}

	return projects, meta, err
}
