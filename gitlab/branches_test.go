package gitlab

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGitlab_ProjectBranches(t *testing.T) {
	ts, gitlab := mockServerFromMapping(t, "branches/project_1_branches.json")
	defer ts.Close()

	branches, meta, err := gitlab.ProtectedBranches("1", nil)

	assert.NoError(t, err)

	assert.Equal(t, 10, len(branches))

	assert.IsType(t, new(ResponseMeta), meta)
	assert.Equal(t, 1, meta.Page)
	assert.Equal(t, 10, meta.PerPage)
}
