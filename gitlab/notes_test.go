package gitlab

import (
	"github.com/plouc/go-gitlab-client/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGitlab_ProjectIssueNotes(t *testing.T) {
	ts := test.CreateMockServerFromDir(t, "notes")
	defer ts.Close()
	gitlab := NewGitlab(ts.URL, "", "")

	collection, meta, err := gitlab.ProjectIssueNotes("1", 5, nil)

	assert.NoError(t, err)

	assert.Equal(t, 2, len(collection.Items))

	assert.IsType(t, new(ResponseMeta), meta)
	assert.Equal(t, 1, meta.Page)
	assert.Equal(t, 10, meta.PerPage)
}

func TestGitlab_ProjectSnippetNotes(t *testing.T) {
	ts := test.CreateMockServerFromDir(t, "notes")
	defer ts.Close()
	gitlab := NewGitlab(ts.URL, "", "")

	collection, meta, err := gitlab.ProjectSnippetNotes("1", 7, nil)

	assert.NoError(t, err)

	assert.Equal(t, 5, len(collection.Items))

	assert.IsType(t, new(ResponseMeta), meta)
	assert.Equal(t, 1, meta.Page)
	assert.Equal(t, 10, meta.PerPage)
}

func TestGitlab_ProjectMergeRequestNotes(t *testing.T) {
	ts := test.CreateMockServerFromDir(t, "notes")
	defer ts.Close()
	gitlab := NewGitlab(ts.URL, "", "")

	collection, meta, err := gitlab.ProjectMergeRequestNotes("1", 3, nil)

	assert.NoError(t, err)

	assert.Equal(t, 6, len(collection.Items))

	assert.IsType(t, new(ResponseMeta), meta)
	assert.Equal(t, 1, meta.Page)
	assert.Equal(t, 10, meta.PerPage)
}
