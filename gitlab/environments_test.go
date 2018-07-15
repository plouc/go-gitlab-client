package gitlab

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectEnvironments(t *testing.T) {
	ts, gitlab := mockServerFromMapping(t, "environments/project_1_environments.json")
	defer ts.Close()

	environments, meta, err := gitlab.ProjectEnvironments("1", nil)

	assert.NoError(t, err)

	assert.Equal(t, 5, len(environments.Items))

	assert.IsType(t, new(ResponseMeta), meta)
	assert.Equal(t, 1, meta.Page)
	assert.Equal(t, 10, meta.PerPage)
}
