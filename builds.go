package gogitlab

import (
	"encoding/json"
	"io"
)

const (
	project_builds          = "/projects/:id/builds"                         // List project builds
	project_build           = "/projects/:id/builds/:build_id"               // Get a single build
	project_commit_builds   = "/projects/:id/repository/commits/:sha/builds" // List commit builds
	project_build_artifacts = "/projects/:id/builds/:build_id/artifacts"     // Get build artifacts
	project_build_cancel    = "/projects/:id/builds/:build_id/cancel"        // Cancel a build
	project_build_retry     = "/projects/:id/builds/:build_id/retry"         // Retry a build
	project_build_erase     = "/projects/:id/builds/:build_id/erase"         // Erase a build
)

type ArtifactsFile struct {
	Filename string `json:"filename"`
	Size     int    `json:"size"`
}

type Build struct {
	Id            int           `json:"id"`
	ArtifactsFile ArtifactsFile `json:"artifacts_file"`
	Commit        Commit        `json:"commit"`
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
}

func (g *Gitlab) ProjectBuilds(id string) ([]*Build, error) {
	url, opaque := g.ResourceUrlRaw(project_builds, map[string]string{
		":id": id,
	})

	builds := make([]*Build, 0)

	contents, err := g.buildAndExecRequestRaw("GET", url, opaque, nil)
	if err != nil {
		return builds, err
	}

	err = json.Unmarshal(contents, &builds)

	return builds, err
}

func (g *Gitlab) ProjectCommitBuilds(id, sha1 string) ([]*Build, error) {
	url, opaque := g.ResourceUrlRaw(project_commit_builds, map[string]string{
		":id":  id,
		":sha": sha1,
	})

	builds := make([]*Build, 0)

	contents, err := g.buildAndExecRequestRaw("GET", url, opaque, nil)
	if err != nil {
		return builds, err
	}

	err = json.Unmarshal(contents, &builds)

	return builds, err
}

func (g *Gitlab) ProjectBuild(id, buildId string) (*Build, error) {
	url, opaque := g.ResourceUrlRaw(project_build, map[string]string{
		":id":       id,
		":build_id": buildId,
	})

	build := &Build{}

	contents, err := g.buildAndExecRequestRaw("GET", url, opaque, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &build)

	return build, err
}

func (g *Gitlab) ProjectBuildArtifacts(id, buildId string) (io.ReadCloser, error) {
	url, _ := g.ResourceUrlRaw(project_build_artifacts, map[string]string{
		":id":       id,
		":build_id": buildId,
	})

	resp, err := g.execRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func (g *Gitlab) ProjectCancelBuild(id, buildId string) (*Build, error) {
	url, opaque := g.ResourceUrlRaw(project_build_cancel, map[string]string{
		":id":       id,
		":build_id": buildId,
	})

	build := &Build{}

	contents, err := g.buildAndExecRequestRaw("POST", url, opaque, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &build)

	return build, err
}

func (g *Gitlab) ProjectRetryBuild(id, buildId string) (*Build, error) {
	url, opaque := g.ResourceUrlRaw(project_build_retry, map[string]string{
		":id":       id,
		":build_id": buildId,
	})

	build := &Build{}

	contents, err := g.buildAndExecRequestRaw("POST", url, opaque, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &build)

	return build, err
}

func (g *Gitlab) ProjectEraseBuild(id, buildId string) (*Build, error) {
	url, opaque := g.ResourceUrlRaw(project_build_erase, map[string]string{
		":id":       id,
		":build_id": buildId,
	})

	build := &Build{}

	contents, err := g.buildAndExecRequestRaw("POST", url, opaque, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &build)

	return build, err
}
