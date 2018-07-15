package gitlab

import (
	"encoding/json"
	"io"
)

const (
	MambersApiPath = "/:type/:id/members"          // List group or project team members
	MamberApiPath  = "/:type/:id/members/:user_id" // Get group or project team member
)

type Member struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	State       string `json:"state"`
	AvatarUrl   string `json:"avatar_url"`
	WebUrl      string `json:"web_url"`
	AccessLevel int    `json:"access_level"`
	ExpiresAt   string `json:"expires_at"`
}

type MemberCollection struct {
	Items []*Member
}

type MembersOptions struct {
	PaginationOptions

	Query string `url:"query,omitempty"`
}

func (m *Member) RenderJson(w io.Writer) error {
	return renderJson(w, m)
}

func (m *Member) RenderYaml(w io.Writer) error {
	return renderYaml(w, m)
}

func (c *MemberCollection) RenderJson(w io.Writer) error {
	return renderJson(w, c.Items)
}

func (c *MemberCollection) RenderYaml(w io.Writer) error {
	return renderYaml(w, c.Items)
}

func (g *Gitlab) getResourceMembers(resourceType, projectId string, o *MembersOptions) (*MemberCollection, *ResponseMeta, error) {
	u := g.ResourceUrlQ(MambersApiPath, map[string]string{
		":type": resourceType,
		":id":   projectId,
	}, o)

	collection := new(MemberCollection)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &collection.Items)
	}

	return collection, meta, err
}

func (g *Gitlab) ProjectMembers(projectId string, o *MembersOptions) (*MemberCollection, *ResponseMeta, error) {
	return g.getResourceMembers("projects", projectId, o)
}

func (g *Gitlab) GroupMembers(groupId string, o *MembersOptions) (*MemberCollection, *ResponseMeta, error) {
	return g.getResourceMembers("groups", groupId, o)
}
