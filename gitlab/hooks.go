package gitlab

import (
	"encoding/json"
)

const (
	ProjectHooksApiPath = "/projects/:id/hooks"
	ProjectHookApiPath  = "/projects/:id/hooks/:hook_id"
)

type HookAddPayload struct {
	Url                      string `json:"url"`
	PushEvents               bool   `json:"push_events"`
	IssuesEvents             bool   `json:"issues_events"`
	ConfidentialIssuesEvents bool   `json:"confidential_issues_events"`
	MergeRequestsEvents      bool   `json:"merge_requests_events"`
	TagPushEvents            bool   `json:"tag_push_events"`
	NoteEvents               bool   `json:"note_events"`
	JobEvents                bool   `json:"job_events"`
	PipelineEvents           bool   `json:"pipeline_events"`
	WikiPageEvents           bool   `json:"wiki_page_events"`
	EnableSslVerification    bool   `json:"enable_ssl_verification"`
	Token                    string `json:"token"`
}

type Hook struct {
	HookAddPayload
	Id        int    `json:"id,omitempty"`
	ProjectId int    `json:"project_id,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

func (g *Gitlab) ProjectHooks(projectId string) ([]*Hook, *ResponseMeta, error) {
	u := g.ResourceUrl(ProjectHooksApiPath, map[string]string{":id": projectId})

	var err error
	var hooks []*Hook

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err != nil {
		return hooks, meta, err
	}

	err = json.Unmarshal(contents, &hooks)

	return hooks, meta, err
}

func (g *Gitlab) ProjectHook(projectId, hookId string) (*Hook, *ResponseMeta, error) {
	u := g.ResourceUrl(ProjectHookApiPath, map[string]string{
		":id":      projectId,
		":hook_id": hookId,
	})

	var err error
	hook := new(Hook)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err != nil {
		return hook, meta, err
	}

	err = json.Unmarshal(contents, &hook)

	return hook, meta, err
}

func (g *Gitlab) AddProjectHook(projectId string, hook *HookAddPayload) (*Hook, *ResponseMeta, error) {
	u := g.ResourceUrl(ProjectHooksApiPath, map[string]string{":id": projectId})

	hookJson, err := json.Marshal(hook)
	if err != nil {
		return nil, nil, err
	}

	var createdHook *Hook
	contents, meta, err := g.buildAndExecRequest("POST", u.String(), hookJson)
	if err == nil {
		err = json.Unmarshal(contents, &createdHook)
	}

	return createdHook, meta, err
}

/*
func (g *Gitlab) UpdateProjectHook(id, hook_id, hook_url string, push_events, issues_events, merge_requests_events bool) (*ResponseMeta, error) {
	u := g.ResourceUrl(ProjectHookApiPath, map[string]string{
		":id":      id,
		":hook_id": hook_id,
	})

	var err error

	body := buildHookQuery(hook_url, push_events, issues_events, merge_requests_events)
	_, meta, err := g.buildAndExecRequest("PUT", u.String(), []byte(body))

	return meta, err
}
*/

func (g *Gitlab) RemoveProjectHook(projectId, hookId string) (*ResponseMeta, error) {
	u := g.ResourceUrl(ProjectHookApiPath, map[string]string{
		":id":      projectId,
		":hook_id": hookId,
	})

	_, meta, err := g.buildAndExecRequest("DELETE", u.String(), nil)

	return meta, err
}
