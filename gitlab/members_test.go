package gitlab

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectMembers(t *testing.T) {
	ts, gitlab := mockServerFromMapping(t, "members/project_1_members.json")
	defer ts.Close()

	members, meta, err := gitlab.ProjectMembers("1", nil)

	assert.NoError(t, err)

	assert.Equal(t, 10, len(members))

	assert.IsType(t, new(ResponseMeta), meta)
	assert.Equal(t, 1, meta.Page)
	assert.Equal(t, 10, meta.PerPage)
}

func TestGroupMembers(t *testing.T) {
	ts, gitlab := mockServerFromMapping(t, "members/group_1_members.json")
	defer ts.Close()

	members, meta, err := gitlab.GroupMembers("1", nil)

	assert.NoError(t, err)

	assert.Equal(t, 10, len(members))

	assert.IsType(t, new(ResponseMeta), meta)
	assert.Equal(t, 1, meta.Page)
	assert.Equal(t, 10, meta.PerPage)
}
