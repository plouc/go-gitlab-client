package gitlab

import (
	"encoding/json"
	"io"
	"strconv"
)

const (
	ProjectCommitsApiPath             = "/projects/:id/repository/commits"
	ProjectMergeRequestCommitsApiPath = "/projects/:id/merge_requests/:merge_request_iid/commits"
	ProjectCommitApiPath              = "/projects/:id/repository/commits/:sha"
	ProjectCommitDiffApiPath          = "/projects/:id/repository/commits/:sha/diff"
	ProjectCommitRefsApiPath          = "/projects/:id/repository/commits/:sha/refs"
	ProjectCommitStatusesApiPath      = "/projects/:id/repository/commits/:sha/statuses"
)

type MinimalCommit struct {
	Id           string `json:"id" yaml:"id"`
	ShortId      string `json:"short_id" yaml:"short_id"`
	Title        string `json:"title" yaml:"title"`
	Message      string `json:"message" yaml:"message"`
	AuthorName   string `json:"author_name" yaml:"author_name"`
	AuthorEmail  string `json:"author_email" yaml:"author_email"`
	CreatedAtRaw string `json:"created_at" yaml:"created_at"`
}

func (c *MinimalCommit) RenderJson(w io.Writer) error {
	return renderJson(w, c)
}

func (c *MinimalCommit) RenderYaml(w io.Writer) error {
	return renderYaml(w, c)
}

type MinimalCommitCollection struct {
	Items []*MinimalCommit
}

func (c *MinimalCommitCollection) RenderJson(w io.Writer) error {
	return renderJson(w, c.Items)
}

func (c *MinimalCommitCollection) RenderYaml(w io.Writer) error {
	return renderYaml(w, c.Items)
}

type Commit struct {
	MinimalCommit    `yaml:",inline"`
	AuthoredDateRaw  string   `json:"authored_date" yaml:"authored_date"`
	CommitterName    string   `json:"committer_name" yaml:"committer_name"`
	CommitterEmail   string   `json:"committer_email" yaml:"committer_email"`
	CommittedDateRaw string   `json:"committed_date" yaml:"committed_date"`
	ParentIds        []string `json:"parent_ids" yaml:"parent_ids"`
	Status           string   `json:"status,omitempty" yaml:"status,omitempty"`
	Stats            struct {
		Additions int `json:"additions" yaml:"additions"`
		Deletions int `json:"deletions" yaml:"deletions"`
		Total     int `json:"total" yaml:"total"`
	} `json:"stats,omitempty" yaml:"stats,omitempty"`
	LastPipeline struct {
		Id     int    `json:"id" yaml:"id"`
		Ref    string `json:"ref" yaml:"ref"`
		Sha    string `json:"sha" yaml:"sha"`
		Status string `json:"status" yaml:"status"`
	} `json:"last_pipeline,omitempty" yaml:"last_pipeline,omitempty"`
}

func (c *Commit) RenderJson(w io.Writer) error {
	return renderJson(w, c)
}

func (c *Commit) RenderYaml(w io.Writer) error {
	return renderYaml(w, c)
}

type CommitCollection struct {
	Items []*Commit
}

func (c *CommitCollection) RenderJson(w io.Writer) error {
	return renderJson(w, c.Items)
}

func (c *CommitCollection) RenderYaml(w io.Writer) error {
	return renderYaml(w, c.Items)
}

type CommitsOptions struct {
	PaginationOptions

	// The file path
	Path string `url:"path,omitempty"`

	// Only commits after or on this date will be returned in ISO 8601 format YYYY-MM-DDTHH:MM:SSZ
	Since string `url:"since,omitempty"`

	// Only commits before or on this date will be returned in ISO 8601 format YYYY-MM-DDTHH:MM:SSZ
	Until string `url:"until,omitempty"`

	// Retrieve every commit from the repository
	All bool `url:"all,omitempty"`

	// Stats about each commit will be added to the response
	WithStats bool `url:"with_stats,omitempty"`
}

func (g *Gitlab) ProjectCommits(projectId string, o *CommitsOptions) (*CommitCollection, *ResponseMeta, error) {
	u := g.ResourceUrlQ(ProjectCommitsApiPath, map[string]string{
		":id": projectId,
	}, o)

	collection := new(CommitCollection)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &collection.Items)
	}

	return collection, meta, err
}

func (g *Gitlab) ProjectMergeRequestCommits(projectId string, mergeRequestIid int, o *PaginationOptions) (*MinimalCommitCollection, *ResponseMeta, error) {
	u := g.ResourceUrlQ(ProjectMergeRequestCommitsApiPath, map[string]string{
		":id":                projectId,
		":merge_request_iid": strconv.Itoa(mergeRequestIid),
	}, o)

	collection := new(MinimalCommitCollection)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &collection.Items)
	}

	return collection, meta, err
}

func (g *Gitlab) ProjectCommit(projectId, commitSha string) (*Commit, *ResponseMeta, error) {
	u := g.ResourceUrl(ProjectCommitApiPath, map[string]string{
		":id":  projectId,
		":sha": commitSha,
	})

	commit := new(Commit)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &commit)
	}

	return commit, meta, err
}

type CommitRef struct {
	Id   string `json:"id" yaml:"id"`
	Sha  string `json:"sha" yaml:"sha"`
	Type string `json:"type" yaml:"type"`
}

type CommitRefCollection struct {
	Items []*CommitRef
}

func (c *CommitRefCollection) RenderJson(w io.Writer) error {
	return renderJson(w, c.Items)
}

func (c *CommitRefCollection) RenderYaml(w io.Writer) error {
	return renderYaml(w, c.Items)
}

func (g *Gitlab) ProjectCommitRefs(projectId, sha string, o *PaginationOptions) (*CommitRefCollection, *ResponseMeta, error) {
	u := g.ResourceUrlQ(ProjectCommitRefsApiPath, map[string]string{
		":id":  projectId,
		":sha": sha,
	}, o)

	collection := new(CommitRefCollection)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &collection.Items)
	}

	return collection, meta, err
}

type CommitStatus struct {
	Id           int    `json:"id" yaml:"id"`
	Ref          string `json:"ref" yaml:"ref"`
	Status       string `json:"status" yaml:"status"`
	CreatedAtRaw string `json:"created_at" yaml:"created_at"`
	StartedAtRaw string `json:"started_at" yaml:"started_at"`
	Name         string `json:"name" yaml:"name"`
	AllowFailure bool   `json:"allow_failure" yaml:"allow_failure"`
	Author       struct {
		Id        int    `json:"id"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		State     string `json:"state"`
		WebUrl    string `json:"web_url"`
		AvatarUrl string `json:"avatar_url"`
	} `json:"author" yaml:"author"`
	Description   string `json:"description" yaml:"description"`
	Sha           string `json:"sha" yaml:"sha"`
	TargetURL     string `json:"target_url" yaml:"target_url"`
	FinishedAtRaw string `json:"finished_at" yaml:"finished_at"`
}

type CommitStatusCollection struct {
	Items []*CommitStatus
}

func (c *CommitStatusCollection) RenderJson(w io.Writer) error {
	return renderJson(w, c.Items)
}

func (c *CommitStatusCollection) RenderYaml(w io.Writer) error {
	return renderYaml(w, c.Items)
}

func (g *Gitlab) ProjectCommitStatuses(projectId, sha string, o *PaginationOptions) (*CommitStatusCollection, *ResponseMeta, error) {
	u := g.ResourceUrlQ(ProjectCommitStatusesApiPath, map[string]string{
		":id":  projectId,
		":sha": sha,
	}, o)

	collection := new(CommitStatusCollection)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &collection.Items)
	}

	return collection, meta, err
}
