package gitlab

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

const (
	projectJobsUrl         = "/projects/:id/jobs"
	projectPipelineJobsUrl = "/projects/:id/pipelines/:pipeline_id/jobs"
	projectJobUrl          = "/projects/:id/jobs/:job_id"
	projectJobTraceUrl     = "/projects/:id/jobs/:job_id/trace"
	cancelProjectJobUrl    = "/projects/:id/jobs/:job_id/cancel"
	retryProjectJobUrl     = "/projects/:id/jobs/:job_id/retry"
	eraseProjectJobUrl     = "/projects/:id/jobs/:job_id/erase"
)

type Job struct {
	Id                int     `json:"id,omitempty" yaml:"id,omitempty"`
	Name              string  `json:"name,omitempty" yaml:"name,omitempty"`
	Status            string  `json:"status,omitempty" yaml:"status,omitempty"`
	Stage             string  `json:"stage,omitempty" yaml:"stage,omitempty"`
	Ref               string  `json:"ref,omitempty" yaml:"ref,omitempty"`
	Tag               bool    `json:"tag,omitempty" yaml:"tag,omitempty"`
	Coverage          float64 `json:"coverage,omitempty" yaml:"coverage,omitempty"`
	CreatedAt         string  `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	StartedAt         string  `json:"started_at,omitempty" yaml:"started_at,omitempty"`
	FinishedAt        string  `json:"finished_at,omitempty" yaml:"finished_at,omitempty"`
	Duration          float64 `json:"duration,omitempty" yaml:"duration,omitempty"`
	ArtifactsExpireAt string  `json:"artifacts_expire_at,omitempty" yaml:"artifacts_expire_at,omitempty"`
	Pipeline          struct {
		Id     int    `json:"id,omitempty" yaml:"id,omitempty"`
		Status string `json:"status,omitempty" yaml:"status,omitempty"`
		Ref    string `json:"ref,omitempty" yaml:"ref,omitempty"`
		Sha    string `json:"sha,omitempty" yaml:"sha,omitempty"`
	} `json:"pipeline" yaml:"pipeline"`
	User struct {
		Id           int    `json:"id,omitempty" yaml:"id,omitempty"`
		Name         string `json:"name,omitempty" yaml:"name,omitempty"`
		Username     string `json:"username,omitempty" yaml:"username,omitempty"`
		State        string `json:"state,omitempty" yaml:"state,omitempty"`
		AvatarUrl    string `json:"avatar_url,omitempty" yaml:"avatar_url,omitempty"`
		WebUrl       string `json:"web_url,omitempty" yaml:"web_url,omitempty"`
		CreatedAt    string `json:"created_at,omitempty" yaml:"created_at,omitempty"`
		Bio          string `json:"bio,omitempty" yaml:"bio,omitempty"`
		Location     string `json:"location,omitempty" yaml:"location,omitempty"`
		Skype        string `json:"skype,omitempty" yaml:"skype,omitempty"`
		Linkedin     string `json:"linkedin,omitempty" yaml:"linkedin,omitempty"`
		Twitter      string `json:"twitter,omitempty" yaml:"twitter,omitempty"`
		WebsiteUrl   string `json:"website_url,omitempty" yaml:"website_url,omitempty"`
		Organization string `json:"organization,omitempty" yaml:"organization,omitempty"`
	} `json:"user" yaml:"user"`
	Commit struct {
		Id             string   `json:"id,omitempty" yaml:"id,omitempty"`
		ShortId        string   `json:"short_id,omitempty" yaml:"short_id,omitempty"`
		Title          string   `json:"title,omitempty" yaml:"title,omitempty"`
		Message        string   `json:"message,omitempty" yaml:"message,omitempty"`
		AuthorName     string   `json:"author_name,omitempty" yaml:"author_name,omitempty"`
		AuthorEmail    string   `json:"author_email,omitempty" yaml:"author_email,omitempty"`
		CommitterName  string   `json:"committer_name,omitempty" yaml:"committer_name,omitempty"`
		CommitterEmail string   `json:"committer_email,omitempty" yaml:"committer_email,omitempty"`
		CreatedAt      string   `json:"created_at,omitempty" yaml:"created_at,omitempty"`
		AuthoredDate   string   `json:"authored_date,omitempty" yaml:"authored_date,omitempty"`
		CommittedDate  string   `json:"committed_date,omitempty" yaml:"committed_date,omitempty"`
		ParentIds      []string `json:"parent_ids,omitempty" yaml:"parent_ids,omitempty"`
	} `json:"commit" yaml:"commit"`
	Runner struct {
		Id          int    `json:"id,omitempty" yaml:"id,omitempty"`
		Name        string `json:"name,omitempty" yaml:"name,omitempty"`
		Description string `json:"description,omitempty" yaml:"description,omitempty"`
		IpAddress   string `json:"ip_address,omitempty" yaml:"ip_address,omitempty"`
		Active      bool   `json:"active,omitempty" yaml:"active,omitempty"`
		IsShared    bool   `json:"is_shared,omitempty" yaml:"is_shared,omitempty"`
		Online      bool   `json:"online,omitempty" yaml:"online,omitempty"`
		Status      string `json:"status,omitempty" yaml:"status,omitempty"`
	} `json:"runner" yaml:"runner"`
}

type JobsOptions struct {
	PaginationOptions
	// The scope of jobs to show, one or array of:
	// created, pending, running, failed, success, canceled, skipped, manual;
	// showing all jobs if none provided
	Scope []string
}

func (g *Gitlab) getJobs(u *url.URL, o *JobsOptions) ([]*Job, *ResponseMeta, error) {
	if o != nil {
		q := u.Query()

		if o.Page != 1 {
			q.Set("page", strconv.Itoa(o.Page))
		}
		if o.PerPage != 0 {
			q.Set("per_page", strconv.Itoa(o.PerPage))
		}
		if len(o.Scope) > 0 {
			q.Set("scope", strings.Join(o.Scope, ","))
		}

		u.RawQuery = q.Encode()
	}

	jobs := make([]*Job, 0)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err != nil {
		return jobs, meta, err
	}

	err = json.Unmarshal(contents, &jobs)

	return jobs, meta, err
}

func (g *Gitlab) ProjectJobs(projectId string, o *JobsOptions) ([]*Job, *ResponseMeta, error) {
	u := g.ResourceUrl(projectJobsUrl, map[string]string{
		":id": projectId,
	})

	return g.getJobs(u, o)
}

func (g *Gitlab) ProjectPipelineJobs(projectId string, pipelineId int, o *JobsOptions) ([]*Job, *ResponseMeta, error) {
	u := g.ResourceUrl(projectPipelineJobsUrl, map[string]string{
		":id":          projectId,
		":pipeline_id": fmt.Sprintf("%d", pipelineId),
	})

	return g.getJobs(u, o)
}

func (g *Gitlab) ProjectJob(projectId string, jobId int) (*Job, *ResponseMeta, error) {
	u := g.ResourceUrl(projectJobUrl, map[string]string{
		":id":     projectId,
		":job_id": strconv.Itoa(jobId),
	})

	job := &Job{}

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err != nil {
		return nil, meta, err
	}

	err = json.Unmarshal(contents, &job)

	return job, meta, err
}

func (g *Gitlab) ProjectJobTrace(projectId string, jobId int) (string, *ResponseMeta, error) {
	u := g.ResourceUrl(projectJobTraceUrl, map[string]string{
		":id":     projectId,
		":job_id": strconv.Itoa(jobId),
	})

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err != nil {
		return "", meta, err
	}

	return string(contents), meta, err
}

func (g *Gitlab) projectJobAction(path, projectId string, jobId int) (*Job, *ResponseMeta, error) {
	u := g.ResourceUrl(path, map[string]string{
		":id":     projectId,
		":job_id": strconv.Itoa(jobId),
	})

	job := &Job{}

	contents, meta, err := g.buildAndExecRequest("POST", u.String(), nil)
	if err != nil {
		return nil, meta, err
	}

	err = json.Unmarshal(contents, &job)

	return job, meta, err
}

func (g *Gitlab) CancelProjectJob(projectId string, jobId int) (*Job, *ResponseMeta, error) {
	return g.projectJobAction(cancelProjectJobUrl, projectId, jobId)
}

func (g *Gitlab) RetryProjectJob(projectId string, jobId int) (*Job, *ResponseMeta, error) {
	return g.projectJobAction(retryProjectJobUrl, projectId, jobId)
}

func (g *Gitlab) EraseProjectJob(projectId string, jobId int) (*Job, *ResponseMeta, error) {
	return g.projectJobAction(eraseProjectJobUrl, projectId, jobId)
}
