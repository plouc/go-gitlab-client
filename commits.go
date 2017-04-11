package gogitlab

import (
	"encoding/json"
	"time"
)

const (
	commit_status = "/projects/:id/repository/commits/:sha/statuses" // Get the statuses of a commit in a project
)

type CommitStatus struct {
	Status       string     `json:"status"`
	CreatedAt    time.Time  `json:"created_at"`
	StartedAt    *time.Time `json:"started_at"`
	Name         string     `json:"name"`
	AllowFailure bool       `json:"allow_failure"`
	Author       Author     `json:"author"`
	Description  *string    `json:"description"`
	Sha          string     `json:"sha"`
	TargetURL    string     `json:"target_url"`
	FinishedAt   *time.Time `json:"finished_at"`
	ID           int        `json:"id"`
	Ref          string     `json:"ref"`
}

type Author struct {
	Username  string `json:"username"`
	State     string `json:"state"`
	WebURL    string `json:"web_url"`
	AvatarURL string `json:"avatar_url"`
	ID        int    `json:"id"`
	Name      string `json:"name"`
}

func (g *Gitlab) ProjectCommitStatuses(id, sha1 string) ([]*CommitStatus, error) {
	url, opaque := g.ResourceUrlRaw(commit_status, map[string]string{
		":id":  id,
		":sha": sha1,
	})

	statuses := make([]*CommitStatus, 0)

	contents, err := g.buildAndExecRequestRaw("GET", url, opaque, nil)
	if err != nil {
		return statuses, err
	}

	err = json.Unmarshal(contents, &statuses)

	return statuses, err
}
