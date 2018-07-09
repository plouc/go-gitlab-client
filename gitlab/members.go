package gitlab

import (
	"encoding/json"
	"strconv"
)

const (
	membersUrl = "/:type/:id/members"          // List group or project team members
	memberUrl  = "/:type/:id/members/:user_id" // Get group or project team member
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
	Query string
}

func (g *Gitlab) getResourceMembers(resourceType, projectId string, o *MembersOptions) ([]*Member, *ResponseMeta, error) {
	u := g.ResourceUrl(membersUrl, map[string]string{
		":type": resourceType,
		":id":   projectId,
	})
	if o != nil {
		q := u.Query()

		if o.Page != 1 {
			q.Set("page", strconv.Itoa(o.Page))
		}
		if o.PerPage != 0 {
			q.Set("per_page", strconv.Itoa(o.PerPage))
		}
		if o.Query != "" {
			q.Set("query", o.Query)
		}

		u.RawQuery = q.Encode()
	}

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
