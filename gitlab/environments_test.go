package gitlab

import (
	"testing"

	"github.com/edison-moreland/go-gitlab-client/test"
	"github.com/stretchr/testify/assert"
)

func TestGitlab_ProjectEnvironments(t *testing.T) {
	ts := test.CreateMockServer(t, []string{
		"environments/project_1_environments",
	})
	defer ts.Close()
	gitlab := NewGitlab(ts.URL, "", "")

	environments, meta, err := gitlab.ProjectEnvironments("1", nil)

	assert.NoError(t, err)

	assert.Equal(t, 5, len(environments.Items))

	assert.IsType(t, new(ResponseMeta), meta)
	assert.Equal(t, 1, meta.Page)
	assert.Equal(t, 10, meta.PerPage)
}
