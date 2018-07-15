package gitlab

import (
	"encoding/json"
	"io"
)

const (
	GroupsApiPath        = "/groups"
	GroupApiPath         = "/groups/:id"
	GroupProjectsApiPath = "/groups/:id/projects"
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

type GroupCollection struct {
	Items []*Group
}

type GroupsOptions struct {
	PaginationOptions
	SortOptions

	// Skip the group IDs passed
	SkipGroups []string `url:"skip_groups,omitempty,comma"`

	// Show all the groups you have access to
	// (defaults to false for authenticated users, true for admin)
	AllAvailable bool `url:"all_available,omitempty"`

	// Return the list of authorized groups matching the search criteria
	Search string `url:"search,omitempty"`

	// Include group statistics (admins only)
	Statistics bool `url:"statistics,omitempty"`

	// Include custom attributes in response (admins only)
	WithCustomAttributes bool `url:"with_custom_attributes,omitempty"`

	// Limit to groups owned by the current user
	Owned bool `url:"owned,omitempty"`
}

func (g *Group) RenderJson(w io.Writer) error {
	return renderJson(w, g)
}

func (g *Group) RenderYaml(w io.Writer) error {
	return renderYaml(w, g)
}

func (c *GroupCollection) RenderJson(w io.Writer) error {
	return renderJson(w, c.Items)
}

func (c *GroupCollection) RenderYaml(w io.Writer) error {
	return renderYaml(w, c.Items)
}

func (g *Gitlab) Groups(o *GroupsOptions) (*GroupCollection, *ResponseMeta, error) {
	u := g.ResourceUrlQ(GroupsApiPath, nil, o)

	collection := new(GroupCollection)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &collection.Items)
	}

	return collection, meta, err
}

func (g *Gitlab) Group(id string, withCustomAttributes bool) (*GroupWithDetails, *ResponseMeta, error) {
	u := g.ResourceUrl(GroupApiPath, map[string]string{":id": id})
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
	u := g.ResourceUrl(GroupsApiPath, nil)

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
	u := g.ResourceUrl(GroupApiPath, map[string]string{":id": id})

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
	u := g.ResourceUrl(GroupApiPath, map[string]string{":id": id})

	var responseWithMessage *ResponseWithMessage
	contents, meta, err := g.buildAndExecRequest("DELETE", u.String(), nil)
	if err != nil {
		return "", meta, err
	}

	err = json.Unmarshal(contents, &responseWithMessage)

	return responseWithMessage.Message, meta, err
}

func (g *Gitlab) GroupProjects(id string) ([]*Project, *ResponseMeta, error) {
	u := g.ResourceUrl(GroupProjectsApiPath, map[string]string{":id": id})

	var projects []*Project

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &projects)
	}

	return projects, meta, err
}
