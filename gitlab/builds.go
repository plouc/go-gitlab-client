package gitlab

import (
	"encoding/json"
	"io"
)

const (
	ProjectCommitBuildsApiPath   = "/projects/:id/repository/commits/:sha/builds"
	ProjectBuildArtifactsApiPath = "/projects/:id/builds/:build_id/artifacts"
)

type ArtifactsFile struct {
	Filename string `json:"filename"`
	Size     int    `json:"size"`
}

type Build struct {
	Id            int           `json:"id"`
	ArtifactsFile ArtifactsFile `json:"artifacts_file"`
	Commit        Commit        `json:"commit,omitempty"`
	CreatedAt     string        `json:"created_at"`
	DownloadURL   string        `json:"download_url"`
	FinishedAt    string        `json:"finished_at"`
	Name          string        `json:"name"`
	Ref           string        `json:"ref"`
	Stage         string        `json:"stage"`
	StartedAt     string        `json:"started_at"`
	Status        string        `json:"status"`
	Tag           bool          `json:"tag"`
	User          User          `json:"user"`
	When          string        `json:"when,omitempty"`
	Manual        bool          `json:"manual,omitempty"`
}

func (g *Gitlab) ProjectCommitBuilds(id, sha1 string) ([]*Build, *ResponseMeta, error) {
	u := g.ResourceUrl(ProjectCommitBuildsApiPath, map[string]string{
		":id":  id,
		":sha": sha1,
	})

	builds := make([]*Build, 0)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err != nil {
		return builds, meta, err
	}

	err = json.Unmarshal(contents, &builds)

	return builds, meta, err
}

func (g *Gitlab) ProjectBuildArtifacts(id, buildId string) (io.ReadCloser, error) {
	u := g.ResourceUrl(ProjectBuildArtifactsApiPath, map[string]string{
		":id":       id,
		":build_id": buildId,
	})

	resp, err := g.execRequest("GET", u.String(), nil)

	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}
