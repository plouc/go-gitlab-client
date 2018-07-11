package gitlab

import (
	"encoding/json"
)

const (
	projectIssuesUrl = "/projects/:id/issues"
)

type Issue struct {
	Id          int        `json:"id"`
	IId         int        `json:"iid"`
	ProjectId   int        `json:"project_id,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Labels      []string   `json:"labels,omitempty"`
	Milestone   *Milestone `json:"milestone,omitempty"`
	Assignee    *User      `json:"assignee,omitempty"`
	Author      *User      `json:"author,omitempty"`
	State       string     `json:"state,omitempty"`
	CreatedAt   string     `json:"created_at,omitempty"`
	UpdatedAt   string     `json:"updated_at,omitempty"`
}

type IssueRequest struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	AssigneeId  int    `json:"assignee_id,omitempty"`
	MilestoneId int    `json:"milestone_id,omitempty"`
	Labels      string `json:"labels,omitempty"`
}

func (g *Gitlab) AddIssue(projectId string, req *IssueRequest) (issue *Issue, meta *ResponseMeta, err error) {
	params := map[string]string{
		":id": projectId,
	}
	u := g.ResourceUrl(projectIssuesUrl, params)

	encodedRequest, err := json.Marshal(req)
	if err != nil {
		return
	}

	data, _, err := g.buildAndExecRequest("POST", u.String(), encodedRequest)
	if err != nil {
		return
	}

	issue = new(Issue)
	err = json.Unmarshal(data, issue)
	if err != nil {
		panic(err)
	}

	return
}
