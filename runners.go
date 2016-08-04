package gogitlab

import (
	"encoding/json"
	"strconv"
)

const (
	runners_url         = "/runners?page=:page&per_page=:per_page"     // Get current users runner list.
	runners_all         = "/runners/all?page=:page&per_page=:per_page" // Get ALL runners list.
	runner_url          = "/runners/:id"                               // Get a single runner.
	project_runners_url = "/projects/:project_id/runners"              // Get ALL project runners.
	project_runner_url  = "/projects/:project_id/runners/:id"          // Get a single project runner.
)

type Runner struct {
	Id           int        `json:"id,omitempty"`
	Name         string     `json:"name,omitempty"`
	Description  string     `json:"description,omitempty"`
	Token        string     `json:"token,omitempty"`
	Revision     string     `json:"revision,omitempty"`
	ContactedAt  string     `json:"contacted_at,omitempty"`
	Platform     string     `json:"platform,omitempty"`
	Version      string     `json:"version,omitempty"`
	Architecture string     `json:"architecture,omitempty"`
	Projects     []*Project `json:"projects,omitempty"`
	TagList      []string   `json:"tag_list,omitempty"`
	Active       bool       `json:"active,omitempty"`
	IsShared     bool       `json:"is_shared,omitempty"`
}

/*
Get all runners owned by the authenticated user.

    GET /runners/:id

Parameters:

    id The ID of a runner

Usage:

  runner, err := gitlab.Runner(your_runner_id)
  if err != nil {
    fmt.Println(err.Error())
  }
  fmt.Printf("%+v\n", runner)
*/
func (g *Gitlab) Runners(page, per_page int) ([]*Runner, error) {

	url := g.ResourceUrl(runners_url, map[string]string{":page": strconv.Itoa(page), ":per_page": strconv.Itoa(per_page)})

	var runners []*Runner

	contents, err := g.buildAndExecRequest("GET", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &runners)
	}

	return runners, err
}

/*
Get a single runner.

    GET /runners/:id

Parameters:

    id The ID of a runner

Usage:

  runner, err := gitlab.Runner(your_runner_id)
  if err != nil {
    fmt.Println(err.Error())
  }
  fmt.Printf("%+v\n", runner)
*/
func (g *Gitlab) Runner(id int) (*Runner, error) {

	url := g.ResourceUrl(runner_url, map[string]string{":id": strconv.Itoa(id)})

	runner := new(Runner)

	contents, err := g.buildAndExecRequest("GET", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &runner)
	}

	return runner, err
}

/*
Get all runners.

    GET /runners/all

Parameters:

    page The start page
    per_page Number of runners per page

Usage:

  runners, err := gitlab.AllRunners(0,20)
  if err != nil {
    fmt.Println(err.Error())
  }
  fmt.Printf("%+v\n", runner)
*/
func (g *Gitlab) AllRunners(page, per_page int) ([]*Runner, error) {

	url := g.ResourceUrl(runners_all, map[string]string{":page": strconv.Itoa(page), ":per_page": strconv.Itoa(per_page)})

	var runners []*Runner

	contents, err := g.buildAndExecRequest("GET", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &runners)
	}

	return runners, err
}

/*
Get all projects runners.

    GET /projects/:id/runners

Parameters:

    page The start page
    per_page Number of runners per page

Usage:

  runners, err := gitlab.AllRunners(0,20)
  if err != nil {
    fmt.Println(err.Error())
  }
  fmt.Printf("%+v\n", runner)
*/
func (g *Gitlab) ProjectRunners(project_id string, page, per_page int) ([]*Runner, error) {

	url := g.ResourceUrl(project_runners_url, map[string]string{":project_id": project_id,
		":page":     strconv.Itoa(page),
		":per_page": strconv.Itoa(per_page),
	})

	var runners []*Runner

	contents, err := g.buildAndExecRequest("GET", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &runners)
	}

	return runners, err
}

/*
Update a specific runner, identified by runner ID,
which is owned by the authentication user.

*/
func (g *Gitlab) UpdateRunner(id int, runner *Runner) (*Runner, error) {

	url := g.ResourceUrl(runner_url, map[string]string{":id": strconv.Itoa(id)})

	encodedRequest, err := json.Marshal(runner)
	if err != nil {
		return nil, err
	}
	var result *Runner

	contents, err := g.buildAndExecRequest("PUT", url, encodedRequest)
	if err == nil {
		err = json.Unmarshal(contents, &result)
	}

	return result, err
}

/*
Enable a specific Project Runner

*/
func (g *Gitlab) EnableProjectRunner(project_id string, id int) (*Runner, error) {

	url := g.ResourceUrl(project_runner_url, map[string]string{":project_id": project_id, ":id": strconv.Itoa(id)})

	request := map[string]int{"runner_id": id}

	encodedRequest, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	var result *Runner

	contents, err := g.buildAndExecRequest("PUT", url, encodedRequest)
	if err == nil {
		err = json.Unmarshal(contents, &result)
	}

	return result, err
}

/*
Disable a specific Project Runner

*/
func (g *Gitlab) DisableProjectRunner(project_id string, id int) (*Runner, error) {

	url := g.ResourceUrl(project_runner_url, map[string]string{":project_id": project_id, ":id": strconv.Itoa(id)})

	var result *Runner

	contents, err := g.buildAndExecRequest("DELETE", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &result)
	}

	return result, err
}

/*
Delete a runner.

    DELETE /runners/:id

Parameters:

    id The id of a runner.

Usage:

  runner, err := gitlab.DeleteRunner(6)
  if err != nil {
    fmt.Println(err.Error())
  }
*/
func (g *Gitlab) DeleteRunner(id int) (*Runner, error) {
	url := g.ResourceUrl(runner_url, map[string]string{":id": strconv.Itoa(id)})

	var result *Runner

	contents, err := g.buildAndExecRequest("DELETE", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &result)
	}

	return result, err
}
