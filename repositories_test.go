package gogitlab

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRepoBranches(t *testing.T) {
  ts, gitlab := Stub("stubs/branches/index.json")
	branches, err := gitlab.RepoBranches("1")

	assert.Equal(t, err, nil)
	assert.Equal(t, len(branches), 1)
  defer ts.Close()
}

func TestRepoBranch(t *testing.T) {
	ts, gitlab := Stub("stubs/branches/show.json")
	branch, err := gitlab.RepoBranch("1", "master")

	assert.Equal(t, err, nil)
	assert.IsType(t, new(Branch), branch)
	assert.Equal(t, branch.Name, "master")
  defer ts.Close()
}

func TestRepoTags(t *testing.T) {
	ts, gitlab := Stub("stubs/tags/index.json")
	tags, err := gitlab.RepoTags("1")

	assert.Equal(t, err, nil)
	assert.Equal(t, len(tags), 1)
  defer ts.Close()
}

func TestRepoCommits(t *testing.T) {
	stub, err := ioutil.ReadFile("stubs/commits/index.json")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(stub))
	}))
	defer ts.Close()

	gitlab := NewGitlab(ts.URL, "", "")
	commits, err := gitlab.RepoCommits("1")

	assert.Equal(t, err, nil)
	assert.Equal(t, len(commits), 2)
}
