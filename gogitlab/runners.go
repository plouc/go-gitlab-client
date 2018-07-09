package gogitlab

import (
	"encoding/json"
	"fmt"
	"strconv"
)

const (
	runnersUrl        = "/runners"                          // Get current users runner list.
	allRunnersUrl     = "/runners/all"                      // Get ALL runners list.
	runnerUrl         = "/runners/:id"                      // Get a single runner.
	projectRunnersUrl = "/projects/:project_id/runners"     // Get ALL project runners.
	projectRunnerUrl  = "/projects/:project_id/runners/:id" // Get a single project runner.
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
	All   bool        // Get a list of all runners in the GitLab instance (specific and shared). Access is restricted to users with admin privileges
	Scope RunnerScope // The scope of runners to show, one of: specific, shared, active, paused, online; showing all runners if none provided
}

func (g *Gitlab) Runners(o *RunnersOptions) ([]*Runner, *ResponseMeta, error) {
	p := runnersUrl
	if o != nil && o.All {
		p = allRunnersUrl
	}
	u := g.ResourceUrl(p, nil)
	if o != nil {
		q := u.Query()

		if o.Page > 1 {
			q.Set("page", strconv.Itoa(o.Page))
		}
		if o.PerPage != 0 {
			q.Set("per_page", strconv.Itoa(o.PerPage))
		}
		if o.Scope != "" {
			q.Set("scope", fmt.Sprintf("%s", o.Scope))
		}

		u.RawQuery = q.Encode()
	}

	var runners []*Runner

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &runners)
	}

	return runners, meta, err
}

func (g *Gitlab) Runner(id int) (*RunnerWithDetails, *ResponseMeta, error) {
	u := g.ResourceUrl(runnerUrl, map[string]string{":id": strconv.Itoa(id)})

	runner := new(RunnerWithDetails)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &runner)
	}

	return runner, meta, err
}

func (g *Gitlab) ProjectRunners(projectId string, page, per_page int) ([]*Runner, *ResponseMeta, error) {
	u := g.ResourceUrl(projectRunnersUrl, map[string]string{":project_id": projectId,
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
	u := g.ResourceUrl(runnerUrl, map[string]string{":id": strconv.Itoa(id)})

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
	u := g.ResourceUrl(projectRunnerUrl, map[string]string{":project_id": projectId, ":id": strconv.Itoa(id)})

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
	u := g.ResourceUrl(projectRunnerUrl, map[string]string{":project_id": projectId, ":id": strconv.Itoa(id)})

	var result *Runner

	contents, meta, err := g.buildAndExecRequest("DELETE", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &result)
	}

	return result, meta, err
}

func (g *Gitlab) DeleteRunner(id int) (*Runner, *ResponseMeta, error) {
	u := g.ResourceUrl(runnerUrl, map[string]string{":id": strconv.Itoa(id)})

	var result *Runner

	contents, meta, err := g.buildAndExecRequest("DELETE", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &result)
	}

	return result, meta, err
}
