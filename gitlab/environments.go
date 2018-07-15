package gitlab

import (
	"encoding/json"
	"strconv"
)

const (
	ProjectEnvironmentsApiPath = "/projects/:id/environments"
	ProjectEnvironmentApiPath  = "/projects/:id/environments/:environment_id"
)

type EnvironmentAddPayload struct {
	Name        string `json:"name,omitempty" yaml:"name,omitempty"`
	ExternalUrl string `json:"external_url,omitempty" yaml:"external_url,omitempty"`
}

type Environment struct {
	Id                    int `json:"id,omitempty" yaml:"id,omitempty"`
	EnvironmentAddPayload `yaml:",inline"`
	Slug                  string              `json:"slug,omitempty" yaml:"slug,omitempty"`
	Project               *EnvironmentProject `json:"project,omitempty" yaml:"project,omitempty"`
}

type EnvironmentProject struct {
	Id                int      `json:"id,omitempty" yaml:"id,omitempty"`
	Name              string   `json:"name,omitempty" yaml:"name,omitempty"`
	NameWithNamespace string   `json:"name_with_namespace,omitempty" yaml:"name_with_namespace,omitempty"`
	Path              string   `json:"path,omitempty" yaml:"path,omitempty"`
	PathWithNamespace string   `json:"path_with_namespace,omitempty" yaml:"path_with_namespace,omitempty"`
	Description       string   `json:"description" yaml:"description"`
	CreatedAtRaw      string   `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	DefaultBranch     string   `json:"default_branch" yaml:"default_branch"`
	TagList           []string `json:"tag_list" yaml:"tag_list"`
	SshUrlToRepo      string   `json:"ssh_url_to_repo" yaml:"ssh_url_to_repo"`
	HttpUrlToRepo     string   `json:"http_url_to_repo" yaml:"http_url_to_repo"`
	WebUrl            string   `json:"web_url" yaml:"web_url"`
	AvatarUrl         string   `json:"avatar_url" yaml:"avatar_url"`
	ForksCount        int      `json:"forks_count" yaml:"forks_count"`
	StarCount         int      `json:"star_count" yaml:"star_count"`
	LastActivityAtRaw string   `json:"last_activity_at,omitempty" yaml:"last_activity_at,omitempty"`
}

func (g *Gitlab) ProjectEnvironments(projectId string, o *PaginationOptions) ([]*Environment, *ResponseMeta, error) {
	u := g.ResourceUrlQ(ProjectEnvironmentsApiPath, map[string]string{
		":id": projectId,
	}, o)

	var environments []*Environment

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &environments)
	}

	return environments, meta, err
}

func (g *Gitlab) AddProjectEnvironment(projectId string, environment *EnvironmentAddPayload) (*Environment, *ResponseMeta, error) {
	u := g.ResourceUrl(ProjectEnvironmentsApiPath, map[string]string{":id": projectId})

	environmentJson, err := json.Marshal(environment)
	if err != nil {
		return nil, nil, err
	}

	var createdEnvironment *Environment
	contents, meta, err := g.buildAndExecRequest("POST", u.String(), environmentJson)
	if err == nil {
		err = json.Unmarshal(contents, &createdEnvironment)
	}

	return createdEnvironment, meta, err
}

func (g *Gitlab) RemoveProjectEnvironment(projectId string, id int) (*ResponseMeta, error) {
	u := g.ResourceUrl(ProjectEnvironmentApiPath, map[string]string{
		":id":             projectId,
		":environment_id": strconv.Itoa(id),
	})

	_, meta, err := g.buildAndExecRequest("DELETE", u.String(), nil)

	return meta, err
}
