package gogitlab

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRepoBranches(t *testing.T) {
	stub, err := ioutil.ReadFile("stubs/branches/index.json")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(stub))
	}))
	defer ts.Close()

	gitlab := NewGitlab(ts.URL, "", "")
	branches, err := gitlab.RepoBranches("1")

	assert.Equal(t, err, nil)
	assert.Equal(t, len(branches), 1)
}

func TestRepoBranch(t *testing.T) {
	stub, err := ioutil.ReadFile("stubs/branches/show.json")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(stub))
	}))
	defer ts.Close()

	gitlab := NewGitlab(ts.URL, "", "")
	branch, err := gitlab.RepoBranch("1", "master")

	assert.Equal(t, err, nil)
	assert.IsType(t, new(Branch), branch)
	assert.Equal(t, branch.Name, "master")
}

func TestRepoTags(t *testing.T) {
	stub, err := ioutil.ReadFile("stubs/tags/index.json")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(stub))
	}))
	defer ts.Close()

	gitlab := NewGitlab(ts.URL, "", "")
	tags, err := gitlab.RepoTags("1")

	assert.Equal(t, err, nil)
	assert.Equal(t, len(tags), 1)
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
