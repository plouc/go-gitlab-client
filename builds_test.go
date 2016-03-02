package gogitlab

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestProjectBuilds(t *testing.T) {
	ts, gitlab := Stub("stubs/builds/list.json")
	defer ts.Close()

	builds, err := gitlab.ProjectBuilds("3")

	assert.Nil(t, err)
	assert.Equal(t, len(builds), 2)
	assert.Equal(t, builds[0].ArtifactsFile.Filename, "artifacts.zip")
}

func TestProjectCommitBuilds(t *testing.T) {
	ts, gitlab := Stub("stubs/builds/commit_builds_list.json")
	defer ts.Close()

	builds, err := gitlab.ProjectBuilds("3")

	assert.Nil(t, err)
	assert.Equal(t, len(builds), 2)
	assert.Equal(t, builds[1].User.AvatarUrl, "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon")
}

func TestProjectBuild(t *testing.T) {
	ts, gitlab := Stub("stubs/builds/build.json")
	defer ts.Close()

	build, err := gitlab.ProjectBuild("3", "12")

	assert.Nil(t, err)
	assert.Equal(t, build.Commit.Id, "0ff3ae198f8601a285adcf5c0fff204ee6fba5fd")
}

func TestProjectCancelBuild(t *testing.T) {
	ts, gitlab := Stub("stubs/builds/cancel.json")
	defer ts.Close()

	build, err := gitlab.ProjectBuild("3", "12")

	assert.Nil(t, err)
	assert.Equal(t, build.Status, "canceled")
}

func TestProjectRetryBuild(t *testing.T) {
	ts, gitlab := Stub("stubs/builds/retry.json")
	defer ts.Close()

	build, err := gitlab.ProjectBuild("3", "12")

	assert.Nil(t, err)
	assert.Equal(t, build.Status, "pending")
}

func TestProjectEraseBuild(t *testing.T) {
	ts, gitlab := Stub("stubs/builds/erase.json")
	defer ts.Close()

	build, err := gitlab.ProjectBuild("3", "12")

	assert.Nil(t, err)
	assert.Equal(t, build.Status, "failed")
}

func TestProjectArtifact(t *testing.T) {
	ts, gitlab := Stub("stubs/builds/content.txt")
	defer ts.Close()

	r, err := gitlab.ProjectBuildArtifacts("3", "12")

	assert.Nil(t, err)

	defer r.Close()

	contents, err := ioutil.ReadAll(r)

	assert.Equal(t, string(contents), "a content")
}
