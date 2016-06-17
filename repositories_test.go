package gogitlab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepoBranches(t *testing.T) {
	ts, gitlab := Stub("stubs/branches/index.json")
	branches, err := gitlab.RepoBranches("1")

	assert.NoError(t, err)
	assert.Equal(t, len(branches), 1)
	defer ts.Close()
}

func TestRepoBranch(t *testing.T) {
	ts, gitlab := Stub("stubs/branches/show.json")
	branch, err := gitlab.RepoBranch("1", "master")

	assert.NoError(t, err)
	assert.IsType(t, new(Branch), branch)
	assert.Equal(t, branch.Name, "master")
	defer ts.Close()
}

func TestRepoTags(t *testing.T) {
	ts, gitlab := Stub("stubs/tags/index.json")
	tags, err := gitlab.RepoTags("1")

	assert.NoError(t, err)
	assert.Equal(t, len(tags), 1)
	defer ts.Close()
}

func TestRepoCommits(t *testing.T) {
	ts, gitlab := Stub("stubs/commits/index.json")
	commits, err := gitlab.RepoCommits("1")

	assert.NoError(t, err)
	assert.Equal(t, len(commits), 2)
	defer ts.Close()
}

func TestRepoTree(t *testing.T) {
	ts, gitlab := Stub("stubs/trees/show.json")
	tree, err := gitlab.RepoTree("1", "path", "ref_name")

	assert.NoError(t, err)
	assert.Equal(t, len(tree), 6)
	defer ts.Close()
}
