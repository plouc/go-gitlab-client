package gogitlab

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProjects(t *testing.T) {
	stub, err := ioutil.ReadFile("stubs/projects/index.json")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(stub))
	}))
	defer ts.Close()

	gitlab := NewGitlab(ts.URL, "", "")
	projects, err := gitlab.Projects()

	assert.Equal(t, err, nil)
	assert.Equal(t, len(projects), 2)
}

func TestProject(t *testing.T) {
	stub, err := ioutil.ReadFile("stubs/projects/show.json")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(stub))
	}))
	defer ts.Close()

	gitlab := NewGitlab(ts.URL, "", "")
	project, err := gitlab.Project("1")

	assert.Equal(t, err, nil)
	assert.IsType(t, new(Project), project)
	assert.Equal(t, project.SshRepoUrl, "git@example.com:diaspora/diaspora-project-site.git")
	assert.Equal(t, project.HttpRepoUrl, "http://example.com/diaspora/diaspora-project-site.git")
}

func TestProjectBranches(t *testing.T) {
	stub, err := ioutil.ReadFile("stubs/projects/branches/index.json")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(stub))
	}))
	defer ts.Close()

	gitlab := NewGitlab(ts.URL, "", "")
	branches, err := gitlab.ProjectBranches("1")

	assert.Equal(t, err, nil)
	assert.Equal(t, len(branches), 2)
}
