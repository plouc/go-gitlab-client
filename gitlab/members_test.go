package gitlab

import (
	"testing"

	"github.com/plouc/go-gitlab-client/test"
	"github.com/stretchr/testify/assert"
)

func TestGitlab_ProjectMembers(t *testing.T) {
	ts := test.CreateMockServer(t, []string{
		"members/project_1_members",
	})
	defer ts.Close()
	gitlab := NewGitlab(ts.URL, "", "")

	members, meta, err := gitlab.ProjectMembers("1", nil)

	assert.NoError(t, err)

	assert.Equal(t, 10, len(members.Items))

	assert.IsType(t, new(ResponseMeta), meta)
	assert.Equal(t, 1, meta.Page)
	assert.Equal(t, 10, meta.PerPage)
}

func TestGitlab_GroupMembers(t *testing.T) {
	ts := test.CreateMockServer(t, []string{
		"members/group_1_members",
	})
	defer ts.Close()
	gitlab := NewGitlab(ts.URL, "", "")

	members, meta, err := gitlab.GroupMembers("1", nil)

	assert.NoError(t, err)

	assert.Equal(t, 10, len(members.Items))

	assert.IsType(t, new(ResponseMeta), meta)
	assert.Equal(t, 1, meta.Page)
	assert.Equal(t, 10, meta.PerPage)
}
