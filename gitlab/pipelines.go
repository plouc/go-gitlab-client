package gitlab

import (
	"encoding/json"
	"io"
)

const (
	ProjectPipelinesApiPath = "/projects/:id/pipelines"
	ProjectPipelineApiPath  = "/projects/:id/pipelines/:pipeline_id"
)

type Pipeline struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
	Ref    string `json:"ref"`
	Sha    string `json:"sha"`
}

type PipelineWithDetails struct {
	Pipeline
	BeforeSha   string `json:"before_sha"`
	Tag         bool   `json:"tag"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	StartedAt   string `json:"started_at"`
	FinishedAt  string `json:"finished_at"`
	CommittedAt string `json:"committed_at"`
	Duration    int    `json:"duration"`
	Coverage    string `json:"coverage"`
	YamlErrors  string `json:"yaml_errors"`
	User        struct {
		Id        int    `json:"id"`
		Username  string `json:"username"`
		Name      string `json:"name"`
		State     string `json:"state"`
		AvatarUrl string `json:"avatar_url"`
		WebUrl    string `json:"web_url"`
	}
}

type PipelineCollection struct {
	Items []*Pipeline
}

type PipelinesOptions struct {
	PaginationOptions
	SortOptions

	// The scope of pipelines, one of:
	// running, pending, finished, branches, tags
	Scope string `url:"scope,omitempty"`

	// The status of pipelines, one of:
	// running, pending, success, failed, canceled, skipped
	Status string `url:"status,omitempty"`

	// The ref of pipelines
	Ref string `url:"ref,omitempty"`

	// The sha or pipelines
	Sha string `url:"sha,omitempty"`

	// Returns pipelines with invalid configurations
	YamlErrors bool `url:"yaml_errors,omitempty"`

	// The name of the user who triggered pipelines
	Name string `url:"name,omitempty"`

	// The username of the user who triggered pipelines
	Username string `url:"username,omitempty"`
}

func (p *PipelineWithDetails) RenderJson(w io.Writer) error {
	return renderJson(w, p)
}

func (p *PipelineWithDetails) RenderYaml(w io.Writer) error {
	return renderYaml(w, p)
}

func (c *PipelineCollection) RenderJson(w io.Writer) error {
	return renderJson(w, c.Items)
}

func (c *PipelineCollection) RenderYaml(w io.Writer) error {
	return renderYaml(w, c.Items)
}

func (g *Gitlab) ProjectPipelines(projectId string, o *PipelinesOptions) (*PipelineCollection, *ResponseMeta, error) {
	u := g.ResourceUrlQ(ProjectPipelinesApiPath, map[string]string{":id": projectId}, o)

	collection := new(PipelineCollection)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &collection.Items)
	}

	return collection, meta, err
}

func (g *Gitlab) ProjectPipeline(projectId, pipelineId string) (*PipelineWithDetails, *ResponseMeta, error) {
	u := g.ResourceUrl(ProjectPipelineApiPath, map[string]string{
		":id":          projectId,
		":pipeline_id": pipelineId,
	})

	pipeline := new(PipelineWithDetails)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &pipeline)
	}

	return pipeline, meta, err
}
