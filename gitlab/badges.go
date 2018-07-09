package gitlab

import (
	"encoding/json"
	"strconv"
)

const (
	projectBadgesUrl = "/projects/:id/badges"
	projectBadgeUrl  = "/projects/:id/badges/:badge_id"
)

type Badge struct {
	Id               int    `json:"id"`
	LinkUrl          string `json:"link_url"`
	ImageUrl         string `json:"image_url"`
	RenderedLinkUrl  string `json:"rendered_link_url"`
	RenderedImageUrl string `json:"rendered_image_url"`
	Kind             string `json:"kind"`
}

func (g *Gitlab) ProjectBadges(projectId string, o *PaginationOptions) ([]*Badge, *ResponseMeta, error) {
	u := g.ResourceUrl(projectBadgesUrl, map[string]string{":id": projectId})

	if o != nil {
		q := u.Query()

		if o.Page != 1 {
			q.Set("page", strconv.Itoa(o.Page))
		}
		if o.PerPage != 0 {
			q.Set("per_page", strconv.Itoa(o.PerPage))
		}

		u.RawQuery = q.Encode()
	}

	var badges []*Badge

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &badges)
	}

	return badges, meta, err
}

func (g *Gitlab) ProjectBadge(projectId, badgeId string) (*Badge, *ResponseMeta, error) {
	u := g.ResourceUrl(projectBadgeUrl, map[string]string{
		":id":       projectId,
		":badge_id": badgeId,
	})

	badge := new(Badge)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &badge)
	}

	return badge, meta, err
}

func (g *Gitlab) AddProjectBadge(projectId string, badge *Badge) (*Badge, *ResponseMeta, error) {
	u := g.ResourceUrl(projectBadgesUrl, map[string]string{":id": projectId})

	badgeJson, err := json.Marshal(badge)
	if err != nil {
		return nil, nil, err
	}

	var createdBadge *Badge
	contents, meta, err := g.buildAndExecRequest("POST", u.String(), badgeJson)
	if err == nil {
		err = json.Unmarshal(contents, &createdBadge)
	}

	return createdBadge, meta, err
}

func (g *Gitlab) RemoveProjectBadge(projectId, badgeId string) (*ResponseMeta, error) {
	u := g.ResourceUrl(projectBadgeUrl, map[string]string{
		":id":       projectId,
		":badge_id": badgeId,
	})

	_, meta, err := g.buildAndExecRequest("DELETE", u.String(), nil)

	return meta, err
}
