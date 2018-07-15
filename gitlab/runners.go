package gitlab

import (
	"encoding/json"
	"io"
	"strconv"
)

const (
	RunnersApiPath        = "/runners"                          // Get current users runner list.
	AllRunnersApiPath     = "/runners/all"                      // Get ALL runners list.
	RunnerApiPath         = "/runners/:id"                      // Get a single runner.
	ProjectRunnersApiPath = "/projects/:project_id/runners"     // Get ALL project runners.
	ProjectRunnerApiPath  = "/projects/:project_id/runners/:id" // Get a single project runner.
)

type Runner struct {
	Id          int    `json:"id,omitempty" yaml:"id,omitempty"`
	Name        string `json:"name,omitempty" yaml:"name,omitempty"`
	Active      bool   `json:"active,omitempty" yaml:"active,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	IpAddress   string `json:"ip_address,omitempty" yaml:"ip_address,omitempty"`
	IsShared    bool   `json:"is_shared,omitempty" yaml:"is_shared,omitempty"`
	Online      bool   `json:"online,omitempty" yaml:"online,omitempty"`
	Status      string `json:"status,omitempty" yaml:"status,omitempty"`
}

type RunnerWithDetails struct {
	Runner         `yaml:",inline"`
	Architecture   string     `json:"architecture,omitempty" yaml:"architecture,omitempty"`
	Platform       string     `json:"platform,omitempty" yaml:"platform,omitempty"`
	Token          string     `json:"token,omitempty" yaml:"token,omitempty"`
	Revision       string     `json:"revision,omitempty" yaml:"revision,omitempty"`
	ContactedAt    string     `json:"contacted_at,omitempty" yaml:"contacted_at,omitempty"`
	Version        string     `json:"version,omitempty" yaml:"version,omitempty"`
	Projects       []*Project `json:"projects,omitempty" yaml:"projects,omitempty"`
	TagList        []string   `json:"tag_list,omitempty" yaml:"tag_list,omitempty"`
	AccessLevel    string     `json:"access_level,omitempty" yaml:"access_level,omitempty"`
	MaximumTimeout int        `json:"maximum_timeout,omitempty" yaml:"maximum_timeout,omitempty"`
}

type RunnerCollection struct {
	Items []*Runner
}

type RunnerScope string

const (
	RunnerScopeSpecific RunnerScope = "specific"
	RunnerScopeShared   RunnerScope = "shared"
	RunnerScopeActive   RunnerScope = "active"
	RunnerScopePaused   RunnerScope = "paused"
	RunnerScopeOnline   RunnerScope = "online"
)

type RunnersOptions struct {
	PaginationOptions
	SortOptions

	// Get a list of all runners in the GitLab instance (specific and shared).
	// Access is restricted to users with admin privileges
	All bool `url:"-"`

	// The scope of runners to show, one of: specific, shared, active, paused, online;
	// showing all runners if none provided
	Scope RunnerScope `url:"scope,omitempty"`
}

func (r *RunnerWithDetails) RenderJson(w io.Writer) error {
	return renderJson(w, r)
}

func (r *RunnerWithDetails) RenderYaml(w io.Writer) error {
	return renderYaml(w, r)
}

func (c *RunnerCollection) RenderJson(w io.Writer) error {
	return renderJson(w, c.Items)
}

func (c *RunnerCollection) RenderYaml(w io.Writer) error {
	return renderYaml(w, c.Items)
}

func (g *Gitlab) Runners(o *RunnersOptions) (*RunnerCollection, *ResponseMeta, error) {
	p := RunnersApiPath
	if o != nil && o.All {
		p = AllRunnersApiPath
	}
	u := g.ResourceUrlQ(p, nil, o)

	collection := new(RunnerCollection)
	runners := make([]*Runner, 0)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &runners)
	}

	collection.Items = runners

	return collection, meta, err
}

func (g *Gitlab) Runner(id int) (*RunnerWithDetails, *ResponseMeta, error) {
	u := g.ResourceUrl(RunnerApiPath, map[string]string{":id": strconv.Itoa(id)})

	runner := new(RunnerWithDetails)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &runner)
	}

	return runner, meta, err
}

func (g *Gitlab) ProjectRunners(projectId string, page, per_page int) ([]*Runner, *ResponseMeta, error) {
	u := g.ResourceUrl(ProjectRunnersApiPath, map[string]string{":project_id": projectId,
		":page":     strconv.Itoa(page),
		":per_page": strconv.Itoa(per_page),
	})

	var runners []*Runner

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &runners)
	}

	return runners, meta, err
}

func (g *Gitlab) UpdateRunner(id int, runner *Runner) (*Runner, *ResponseMeta, error) {
	u := g.ResourceUrl(RunnerApiPath, map[string]string{":id": strconv.Itoa(id)})

	encodedRequest, err := json.Marshal(runner)
	if err != nil {
		return nil, nil, err
	}

	var result *Runner

	contents, meta, err := g.buildAndExecRequest("PUT", u.String(), encodedRequest)
	if err == nil {
		err = json.Unmarshal(contents, &result)
	}

	return result, meta, err
}

func (g *Gitlab) EnableProjectRunner(projectId string, id int) (*Runner, *ResponseMeta, error) {
	u := g.ResourceUrl(ProjectRunnerApiPath, map[string]string{":project_id": projectId, ":id": strconv.Itoa(id)})

	request := map[string]int{"runner_id": id}

	encodedRequest, err := json.Marshal(request)
	if err != nil {
		return nil, nil, err
	}
	var result *Runner

	contents, meta, err := g.buildAndExecRequest("PUT", u.String(), encodedRequest)
	if err == nil {
		err = json.Unmarshal(contents, &result)
	}

	return result, meta, err
}

func (g *Gitlab) DisableProjectRunner(projectId string, id int) (*Runner, *ResponseMeta, error) {
	u := g.ResourceUrl(ProjectRunnerApiPath, map[string]string{":project_id": projectId, ":id": strconv.Itoa(id)})

	var result *Runner

	contents, meta, err := g.buildAndExecRequest("DELETE", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &result)
	}

	return result, meta, err
}

func (g *Gitlab) DeleteRunner(id int) (*Runner, *ResponseMeta, error) {
	u := g.ResourceUrl(RunnerApiPath, map[string]string{":id": strconv.Itoa(id)})

	var result *Runner

	contents, meta, err := g.buildAndExecRequest("DELETE", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &result)
	}

	return result, meta, err
}
