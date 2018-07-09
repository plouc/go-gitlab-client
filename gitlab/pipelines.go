package gitlab

import (
	"encoding/json"
	"strconv"
)

const (
	projectPipelinesUrl = "/projects/:id/pipelines"
	projectPipelineUrl  = "/projects/:id/pipelines/:pipeline_id"
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

type PipelinesOptions struct {
	PaginationOptions
	Scope      string // The scope of pipelines, one of: running, pending, finished, branches, tags
	Status     string // The status of pipelines, one of: running, pending, success, failed, canceled, skipped
	Ref        string // The ref of pipelines
	Sha        string // The sha or pipelines
	YamlErrors bool   // Returns pipelines with invalid configurations
	Name       string // The name of the user who triggered pipelines
	Username   string // The username of the user who triggered pipelines
	OrderBy    string // Order pipelines by id, status, ref, or user_id (default: id)
	Sort       string // Sort pipelines in asc or desc order (default: desc)
}

func (g *Gitlab) ProjectPipelines(projectId string, o *PipelinesOptions) ([]*Pipeline, *ResponseMeta, error) {
	u := g.ResourceUrl(projectPipelinesUrl, map[string]string{":id": projectId})

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

	var pipelines []*Pipeline

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &pipelines)
	}

	return pipelines, meta, err
}

func (g *Gitlab) ProjectPipeline(projectId, pipelineId string) (*PipelineWithDetails, *ResponseMeta, error) {
	u := g.ResourceUrl(projectPipelineUrl, map[string]string{
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
