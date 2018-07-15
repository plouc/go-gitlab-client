package gitlab

import (
	"encoding/json"
	"io"
	"strconv"
)

const (
	ProjectBadgesApiPath = "/projects/:id/badges"
	ProjectBadgeApiPath  = "/projects/:id/badges/:badge_id"
)

type Badge struct {
	Id               int    `json:"id" yaml:"id"`
	LinkUrl          string `json:"link_url" yaml:"link_url"`
	ImageUrl         string `json:"image_url" yaml:"image_url"`
	RenderedLinkUrl  string `json:"rendered_link_url" yaml:"rendered_link_url"`
	RenderedImageUrl string `json:"rendered_image_url" yaml:"rendered_image_url"`
	Kind             string `json:"kind" yaml:"kind"`
}

type BadgeCollection struct {
	Items []*Badge
}

func (b *Badge) RenderJson(w io.Writer) error {
	return renderJson(w, b)
}

func (b *Badge) RenderYaml(w io.Writer) error {
	return renderYaml(w, b)
}

func (c *BadgeCollection) RenderJson(w io.Writer) error {
	return renderJson(w, c.Items)
}

func (c *BadgeCollection) RenderYaml(w io.Writer) error {
	return renderYaml(w, c.Items)
}

func (g *Gitlab) ProjectBadges(projectId string, o *PaginationOptions) (*BadgeCollection, *ResponseMeta, error) {
	u := g.ResourceUrlQ(ProjectBadgesApiPath, map[string]string{":id": projectId}, o)

	collection := new(BadgeCollection)
	badges := make([]*Badge, 0)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &badges)
	}

	collection.Items = badges

	return collection, meta, err
}

func (g *Gitlab) ProjectBadge(projectId string, badgeId int) (*Badge, *ResponseMeta, error) {
	u := g.ResourceUrl(ProjectBadgeApiPath, map[string]string{
		":id":       projectId,
		":badge_id": strconv.Itoa(badgeId),
	})

	badge := new(Badge)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &badge)
	}

	return badge, meta, err
}

func (g *Gitlab) AddProjectBadge(projectId string, badge *Badge) (*Badge, *ResponseMeta, error) {
	u := g.ResourceUrl(ProjectBadgesApiPath, map[string]string{":id": projectId})

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
	u := g.ResourceUrl(ProjectBadgeApiPath, map[string]string{
		":id":       projectId,
		":badge_id": badgeId,
	})

	_, meta, err := g.buildAndExecRequest("DELETE", u.String(), nil)

	return meta, err
}
