package gogitlab

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProjects(t *testing.T) {
	projects_stub, err := ioutil.ReadFile("stubs/projects/index.json")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(projects_stub))
	}))
	defer ts.Close()

	gitlab := NewGitlab(ts.URL, "", "")
	projects, err := gitlab.Projects()
	project := projects[0]

	assert.Equal(t, err, nil)
	assert.Equal(t, len(projects), 2)
	assert.Equal(t, project.SshRepoUrl, "git@example.com:diaspora/diaspora-client.git")
	assert.Equal(t, project.HttpRepoUrl, "http://example.com/diaspora/diaspora-client.git")
}
