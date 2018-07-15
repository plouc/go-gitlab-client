package gitlab

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strconv"
)

const (
	ProjectJobsApiPath         = "/projects/:id/jobs"
	ProjectPipelineJobsApiPath = "/projects/:id/pipelines/:pipeline_id/jobs"
	ProjectJobApiPath          = "/projects/:id/jobs/:job_id"
	ProjectJobTraceApiPath     = "/projects/:id/jobs/:job_id/trace"
	CancelProjectJobApiPath    = "/projects/:id/jobs/:job_id/cancel"
	RetryProjectJobApiPath     = "/projects/:id/jobs/:job_id/retry"
	EraseProjectJobApiPath     = "/projects/:id/jobs/:job_id/erase"
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

type JobCollection struct {
	Items []*Job
}

type JobsOptions struct {
	PaginationOptions
	SortOptions

	// The scope of jobs to show, one or array of:
	// created, pending, running, failed, success, canceled, skipped, manual;
	// showing all jobs if none provided
	Scope []string `url:"scope,omitempty"`
}

func (j *Job) RenderJson(w io.Writer) error {
	return renderJson(w, j)
}

func (j *Job) RenderYaml(w io.Writer) error {
	return renderYaml(w, j)
}

func (c *JobCollection) RenderJson(w io.Writer) error {
	return renderJson(w, c.Items)
}

func (c *JobCollection) RenderYaml(w io.Writer) error {
	return renderYaml(w, c.Items)
}

func (g *Gitlab) getJobs(u *url.URL) (*JobCollection, *ResponseMeta, error) {
	collection := new(JobCollection)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err != nil {
		return collection, meta, err
	}

	err = json.Unmarshal(contents, &collection.Items)

	return collection, meta, err
}

func (g *Gitlab) ProjectJobs(projectId string, o *JobsOptions) (*JobCollection, *ResponseMeta, error) {
	u := g.ResourceUrlQ(ProjectJobsApiPath, map[string]string{
		":id": projectId,
	}, o)

	return g.getJobs(u)
}

func (g *Gitlab) ProjectPipelineJobs(projectId string, pipelineId int, o *JobsOptions) (*JobCollection, *ResponseMeta, error) {
	u := g.ResourceUrlQ(ProjectPipelineJobsApiPath, map[string]string{
		":id":          projectId,
		":pipeline_id": fmt.Sprintf("%d", pipelineId),
	}, o)

	return g.getJobs(u)
}

func (g *Gitlab) ProjectJob(projectId string, jobId int) (*Job, *ResponseMeta, error) {
	u := g.ResourceUrl(ProjectJobApiPath, map[string]string{
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
	u := g.ResourceUrl(ProjectJobTraceApiPath, map[string]string{
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
	return g.projectJobAction(CancelProjectJobApiPath, projectId, jobId)
}

func (g *Gitlab) RetryProjectJob(projectId string, jobId int) (*Job, *ResponseMeta, error) {
	return g.projectJobAction(RetryProjectJobApiPath, projectId, jobId)
}

func (g *Gitlab) EraseProjectJob(projectId string, jobId int) (*Job, *ResponseMeta, error) {
	return g.projectJobAction(EraseProjectJobApiPath, projectId, jobId)
}

// Aggregate jobs by:
//   - pipeline
//   - stage
//   - job name
//
// The resulting aggregation can be used to built something similar as the GitLab's UI
// used to display pipeline details.
func AggregateJobs(jobs []*Job) map[int]map[string]map[string][]*Job {
	agg := map[int]map[string]map[string][]*Job{}

	for _, job := range jobs {
		_, ok := agg[job.Pipeline.Id]
		if !ok {
			agg[job.Pipeline.Id] = map[string]map[string][]*Job{}
		}

		_, ok = agg[job.Pipeline.Id][job.Stage]
		if !ok {
			agg[job.Pipeline.Id][job.Stage] = map[string][]*Job{}
		}

		_, ok = agg[job.Pipeline.Id][job.Stage][job.Name]
		if !ok {
			agg[job.Pipeline.Id][job.Stage][job.Name] = []*Job{}
		}

		agg[job.Pipeline.Id][job.Stage][job.Name] = append(agg[job.Pipeline.Id][job.Stage][job.Name], job)
	}

	return agg
}
