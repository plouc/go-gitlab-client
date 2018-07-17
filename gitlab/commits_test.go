package gitlab

import (
	"testing"

	"github.com/plouc/go-gitlab-client/test"
	"github.com/stretchr/testify/assert"
)

func TestGitlab_ProjectCommits(t *testing.T) {
	ts := test.CreateMockServer(t, []string{
		"commits/project_1_commits",
	})
	defer ts.Close()
	gitlab := NewGitlab(ts.URL, "", "")

	c, meta, err := gitlab.ProjectCommits("1", nil)

	assert.NoError(t, err)

	assert.Equal(t, 2, len(c.Items))
	assert.Equal(t, "Replace sanitize with escape once", c.Items[0].Title)
	assert.Equal(t, "Sanitize for network graph", c.Items[1].Title)

	assert.IsType(t, new(ResponseMeta), meta)
	assert.Equal(t, 1, meta.Page)
	assert.Equal(t, 10, meta.PerPage)
}

func TestGitlab_ProjectCommit(t *testing.T) {
	ts := test.CreateMockServer(t, []string{
		"commits/project_1_commit_6104942438c14ec7bd21c6cd5bd995272b3faff6",
	})
	defer ts.Close()
	gitlab := NewGitlab(ts.URL, "", "")

	c, meta, err := gitlab.ProjectCommit("1", "6104942438c14ec7bd21c6cd5bd995272b3faff6")

	assert.NoError(t, err)

	assert.IsType(t, new(Commit), c)
	assert.Equal(t, "Sanitize for network graph", c.Title)

	assert.IsType(t, new(ResponseMeta), meta)
}

func TestGitlab_ProjectMergeRequestCommits(t *testing.T) {
	ts := test.CreateMockServer(t, []string{
		"commits/project_1_merge_request_1_commits",
	})
	defer ts.Close()
	gitlab := NewGitlab(ts.URL, "", "")

	c, meta, err := gitlab.ProjectMergeRequestCommits("1", 1, nil)

	assert.NoError(t, err)

	assert.Equal(t, 2, len(c.Items))
	assert.Equal(t, "Replace sanitize with escape once", c.Items[0].Title)
	assert.Equal(t, "Sanitize for network graph", c.Items[1].Title)

	assert.IsType(t, new(ResponseMeta), meta)
	assert.Equal(t, 1, meta.Page)
	assert.Equal(t, 10, meta.PerPage)
}
