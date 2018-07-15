package gitlab

import (
	"encoding/json"
	"io"
)

const (
	ProjectHooksApiPath = "/projects/:id/hooks"
	ProjectHookApiPath  = "/projects/:id/hooks/:hook_id"
)

type Hook struct {
	HookAddPayload
	Id        int    `json:"id,omitempty" yaml:"id,omitempty"`
	ProjectId int    `json:"project_id,omitempty" yaml:"project_id,omitempty"`
	CreatedAt string `json:"created_at,omitempty" yaml:"created_at,omitempty"`
}

type HookCollection struct {
	Items []*Hook
}

type HookAddPayload struct {
	Url                      string `json:"url" yaml:"url"`
	PushEvents               bool   `json:"push_events" yaml:"push_events"`
	IssuesEvents             bool   `json:"issues_events" yaml:"issues_events"`
	ConfidentialIssuesEvents bool   `json:"confidential_issues_events" yaml:"confidential_issues_events"`
	MergeRequestsEvents      bool   `json:"merge_requests_events" yaml:"merge_requests_events"`
	TagPushEvents            bool   `json:"tag_push_events" yaml:"tag_push_events"`
	NoteEvents               bool   `json:"note_events" yaml:"note_events"`
	JobEvents                bool   `json:"job_events" yaml:"job_events"`
	PipelineEvents           bool   `json:"pipeline_events" yaml:"pipeline_events"`
	WikiPageEvents           bool   `json:"wiki_page_events" yaml:"wiki_page_events"`
	EnableSslVerification    bool   `json:"enable_ssl_verification" yaml:"enable_ssl_verification"`
	Token                    string `json:"token" yaml:"token"`
}

func (h *Hook) RenderJson(w io.Writer) error {
	return renderJson(w, h)
}

func (h *Hook) RenderYaml(w io.Writer) error {
	return renderYaml(w, h)
}

func (c *HookCollection) RenderJson(w io.Writer) error {
	return renderJson(w, c.Items)
}

func (c *HookCollection) RenderYaml(w io.Writer) error {
	return renderYaml(w, c.Items)
}

func (g *Gitlab) ProjectHooks(projectId string) (*HookCollection, *ResponseMeta, error) {
	u := g.ResourceUrl(ProjectHooksApiPath, map[string]string{":id": projectId})

	collection := new(HookCollection)
	var err error

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err != nil {
		return collection, meta, err
	}

	err = json.Unmarshal(contents, &collection.Items)

	return collection, meta, err
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
