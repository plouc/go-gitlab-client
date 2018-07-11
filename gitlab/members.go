package gitlab

import (
	"encoding/json"
)

const (
	MambersApiPath = "/:type/:id/members"          // List group or project team members
	MamberApiPath  = "/:type/:id/members/:user_id" // Get group or project team member
)

type Member struct {
	Id          int    `json:"id"`
	Username    string `json:"username"`
	Name        string `json:"name"`
	State       string `json:"state"`
	AvatarUrl   string `json:"avatar_url"`
	WebUrl      string `json:"web_url"`
	CreatedAt   string `json:"created_at"`
	ExpiresAt   string `json:"expires_at"`
	AccessLevel int    `json:"access_level"`
}

type MembersOptions struct {
	PaginationOptions

	Query string `url:"query,omitempty"`
}

func (g *Gitlab) getResourceMembers(resourceType, projectId string, o *MembersOptions) ([]*Member, *ResponseMeta, error) {
	u := g.ResourceUrlQ(MambersApiPath, map[string]string{
		":type": resourceType,
		":id":   projectId,
	}, o)

	var members []*Member

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &members)
	}

	return members, meta, err
}

func (g *Gitlab) ProjectMembers(projectId string, o *MembersOptions) ([]*Member, *ResponseMeta, error) {
	return g.getResourceMembers("projects", projectId, o)
}

func (g *Gitlab) GroupMembers(groupId string, o *MembersOptions) ([]*Member, *ResponseMeta, error) {
	return g.getResourceMembers("groups", groupId, o)
}
