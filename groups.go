package gogitlab

import "encoding/json"

const (
	group_members_url = "/groups/:id/members"
)

func (g *Gitlab) GroupMembers(id string) ([]*Member, error) {
	url, opaque := g.ResourceUrlRaw(group_members_url, map[string]string{":id": id})

	var members []*Member

	contents, err := g.buildAndExecRequestEx("GET", url, opaque, nil, true)
	if err == nil {
		err = json.Unmarshal(contents, &members)
	}

	return members, err
}
