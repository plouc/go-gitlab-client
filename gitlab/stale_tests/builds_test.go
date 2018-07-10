package stale_tests

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestProjectCommitBuilds(t *testing.T) {
	ts, gitlab := Stub("stubs/builds/commit_builds_list.json")
	defer ts.Close()

	builds, err := gitlab.ProjectBuilds("3")

	assert.Nil(t, err)
	assert.Equal(t, len(builds), 2)
	assert.Equal(t, builds[1].User.AvatarUrl, "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon")
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
