package gogitlab

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

const (
	project_url_hooks = "/projects/:id/hooks"          // Get list of project hooks
	project_url_hook  = "/projects/:id/hooks/:hook_id" // Get single project hook
)

type Hook struct {
	Id           int    `json:"id,omitempty"`
	Url          string `json:"url,omitempty"`
	CreatedAtRaw string `json:"created_at,omitempty"`
}

type HookObjAttr struct {
	Id              int       `json:"id,omitempty"`
	Title           string    `json:"title,omitempty"`
	AssigneeId      int       `json:"assignee_id,omitempty"`
	AuthorId        int       `json:"author_id,omitempty"`
	ProjectId       int       `json:"project_id,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	Position        int       `json:"position,omitempty"`
	BranchName      string    `json:"branch_name,omitempty"`
	Description     string    `json:"description,omitempty"`
	MilestoneId     int       `json:"milestone_id,omitempty"`
	State           string    `json:"state,omitempty"`
	IId             int       `json:"iid,omitempty"`
	TargetBranch    string    `json:"target_branch,omitempty"`
	SourceBranch    string    `json:"source_branch,omitempty"`
	SourceProjectId int       `json:"source_project_id,omitempty"`
	StCommits       string    `json:"st_commits,omitempty"`
	StDiffs         string    `json:"st_diffs,omitempty"`
	MergeStatus     string    `json:"merge_status,omitempty"`
	TargetProjectId int       `json:"target_project_id,omitempty"`
}

type hRepository struct {
	Name        string `json:"name,omitempty"`
	URL         string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
	Homepage    string `json:"homepage,omitempty"`
}

type hCommit struct {
	Id        string  `json:"id,omitempty"`
	Message   string  `json:"message,omitempty"`
	Timestamp string  `json:"timestamp,omitempty"`
	URL       string  `json:"url,omitempty"`
	Author    *Person `json:"author,omitempty"`
}

type HookPayload struct {
	Before            string       `json:"before,omitempty"`
	After             string       `json:"after,omitempty"`
	Ref               string       `json:"ref,omitempty"`
	UserId            int          `json:"user_id,omitempty"`
	UserName          string       `json:"user_name,omitempty"`
	ProjectId         int          `json:"project_id,omitempty"`
	Repository        *hRepository `json:"repository,omitempty"`
	Commits           []hCommit    `json:"commits,omitempty"`
	TotalCommitsCount int          `json:"total_commits_count,omitempty"`
	ObjectKind        string       `json:"object_kind,omitempty"`
	ObjectAttributes  *HookObjAttr `json:"object_attributes,omitempty"`
}

/*
Get list of project hooks.

    GET /projects/:id/hooks

Parameters:

    id The ID of a project

*/
func (g *Gitlab) ProjectHooks(id string) ([]*Hook, error) {

	url, opaque := g.ResourceUrlRaw(project_url_hooks, map[string]string{":id": id})

	var err error
	var hooks []*Hook

	contents, err := g.buildAndExecRequestRaw("GET", url, opaque, nil)
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

	url, opaque := g.ResourceUrlRaw(project_url_hook, map[string]string{
		":id":      id,
		":hook_id": hook_id,
	})

	var err error
	hook := new(Hook)

	contents, err := g.buildAndExecRequestRaw("GET", url, opaque, nil)
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

	url, opaque := g.ResourceUrlRaw(project_url_hooks, map[string]string{":id": id})

	var err error

	body := buildHookQuery(hook_url, push_events, issues_events, merge_requests_events)
	_, err = g.buildAndExecRequestRaw("POST", url, opaque, []byte(body))

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

	url, opaque := g.ResourceUrlRaw(project_url_hook, map[string]string{
		":id":      id,
		":hook_id": hook_id,
	})

	var err error

	body := buildHookQuery(hook_url, push_events, issues_events, merge_requests_events)
	_, err = g.buildAndExecRequestRaw("PUT", url, opaque, []byte(body))

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

	url, opaque := g.ResourceUrlRaw(project_url_hook, map[string]string{
		":id":      id,
		":hook_id": hook_id,
	})

	var err error

	_, err = g.buildAndExecRequestRaw("DELETE", url, opaque, nil)

	return err
}

func ParseHook(payload []byte) (*HookPayload, error) {
	hp := HookPayload{}
	if err := json.Unmarshal(payload, &hp); err != nil {
		return nil, err
	}

	// Basic sanity check
	switch {
	case len(hp.ObjectKind) == 0:
		// Assume this is a post-receive within repository
		if len(hp.After) == 0 {
			return nil, fmt.Errorf("Invalid hook received, commit hash not found.")
		}
	case hp.ObjectKind == "issue":
		fallthrough
	case hp.ObjectKind == "merge_request":
		if hp.ObjectAttributes == nil {
			return nil, fmt.Errorf("Invalid hook received, attributes not found.")
		}
	default:
		return nil, fmt.Errorf("Invalid hook received, payload format not recognized.")
	}

	return &hp, nil
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
