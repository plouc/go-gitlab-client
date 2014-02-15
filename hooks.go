package gogitlab

import (
	"encoding/json"
	"net/url"
	"strings"
)

const (
	project_url_hooks = "/projects/:id/hooks"          // Get list of project hooks
	project_url_hook  = "/projects/:id/hooks/:hook_id" // Get single project hook
)

/*
Get list of project hooks.

    GET /projects/:id/hooks

Parameters:

    id The ID of a project

*/
func (g *Gitlab) ProjectHooks(id string) ([]*Hook, error) {

	url := strings.Replace(project_url_hooks, ":id", id, -1)
	url = g.BaseUrl + g.ApiPath + url + "?private_token=" + g.Token

	var err error
	var hooks []*Hook

	contents, err := g.buildAndExecRequest("GET", url, nil)
	if err != nil {
		return hooks, err
	}

	err = json.Unmarshal(contents, &hooks)

	return hooks, err
}

/*
Get single project hook.

    GET /projects/:id/hooks/:hook_id

Parameters:

    id      The ID of a project
    hook_id The ID of a hook

*/
func (g *Gitlab) ProjectHook(id, hook_id string) (*Hook, error) {

	url := strings.Replace(project_url_hook, ":id", id, -1)
	url = strings.Replace(url, ":hook_id", hook_id, -1)
	url = g.BaseUrl + g.ApiPath + url + "?private_token=" + g.Token

	var err error
	var hook *Hook

	contents, err := g.buildAndExecRequest("GET", url, nil)
	if err != nil {
		return hook, err
	}

	err = json.Unmarshal(contents, &hook)

	return hook, err
}

/*
Add new project hook.

    POST /projects/:id/hooks

Parameters:

    id                    The ID or NAMESPACE/PROJECT_NAME of a project
    hook_url              The hook URL
    push_events           Trigger hook on push events
    issues_events         Trigger hook on issues events
    merge_requests_events Trigger hook on merge_requests events

*/
func (g *Gitlab) AddProjectHook(id, hook_url string, push_events, issues_events, merge_requests_events bool) error {

	path := strings.Replace(project_url_hooks, ":id", id, -1)
	path = g.BaseUrl + g.ApiPath + path + "?private_token=" + g.Token

	var err error

	body := buildHookQuery(hook_url, push_events, issues_events, merge_requests_events)
	_, err = g.buildAndExecRequest("POST", path, []byte(body))

	return err
}

/*
Edit existing project hook.

    PUT /projects/:id/hooks/:hook_id

Parameters:

    id                    The ID or NAMESPACE/PROJECT_NAME of a project
    hook_id               The ID of a project hook
    hook_url              The hook URL
    push_events           Trigger hook on push events
    issues_events         Trigger hook on issues events
    merge_requests_events Trigger hook on merge_requests events

*/
func (g *Gitlab) EditProjectHook(id, hook_id, hook_url string, push_events, issues_events, merge_requests_events bool) error {

	path := strings.Replace(project_url_hook, ":id", id, -1)
	path = strings.Replace(path, ":hook_id", hook_id, -1)
	path = g.BaseUrl + g.ApiPath + path + "?private_token=" + g.Token

	var err error

	body := buildHookQuery(hook_url, push_events, issues_events, merge_requests_events)
	_, err = g.buildAndExecRequest("PUT", path, []byte(body))

	return err
}

/*
Remove hook from project.

    DELETE /projects/:id/hooks/:hook_id

Parameters:

    id      The ID or NAMESPACE/PROJECT_NAME of a project
    hook_id The ID of hook to delete

*/
func (g *Gitlab) RemoveProjectHook(id, hook_id string) error {

	url := strings.Replace(project_url_hook, ":id", id, -1)
	url = strings.Replace(url, ":hook_id", hook_id, -1)
	url = g.BaseUrl + g.ApiPath + url + "?private_token=" + g.Token

	var err error

	_, err = g.buildAndExecRequest("DELETE", url, nil)

	return err
}

/*
Build HTTP query to add or edit hook
*/
func buildHookQuery(hook_url string, push_events, issues_events, merge_requests_events bool) string {

	v := url.Values{}
	v.Set("url", hook_url)

	if push_events {
		v.Set("push_events", "true")
	} else {
		v.Set("push_events", "false")
	}
	if issues_events {
		v.Set("issues_events", "true")
	} else {
		v.Set("issues_events", "false")
	}
	if merge_requests_events {
		v.Set("merge_requests_events", "true")
	} else {
		v.Set("merge_requests_events", "false")
	}

	return v.Encode()
}
