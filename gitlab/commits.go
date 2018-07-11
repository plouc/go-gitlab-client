package gitlab

import (
	"encoding/json"
	"time"
)

const (
	ProjectCommitStatusesApiPath = "/projects/:id/repository/commits/:sha/statuses"
)

type CommitStatus struct {
	Status       string     `json:"status"`
	CreatedAt    time.Time  `json:"created_at"`
	StartedAt    *time.Time `json:"started_at"`
	Name         string     `json:"name"`
	AllowFailure bool       `json:"allow_failure"`
	Author       User       `json:"author"`
	Description  *string    `json:"description"`
	Sha          string     `json:"sha"`
	TargetURL    string     `json:"target_url"`
	FinishedAt   *time.Time `json:"finished_at"`
	ID           int        `json:"id"`
	Ref          string     `json:"ref"`
}

func (g *Gitlab) ProjectCommitStatuses(id, sha1 string) ([]*CommitStatus, *ResponseMeta, error) {
	u := g.ResourceUrl(ProjectCommitStatusesApiPath, map[string]string{
		":id":  id,
		":sha": sha1,
	})

	statuses := make([]*CommitStatus, 0)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err != nil {
		return statuses, meta, err
	}

	err = json.Unmarshal(contents, &statuses)

	return statuses, meta, err
}
