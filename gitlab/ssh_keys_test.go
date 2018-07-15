package gitlab

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrentUserSshKeys(t *testing.T) {
	ts, gitlab := mockServerFromMapping(t, "ssh_keys/current_user_ssh_keys.json")
	defer ts.Close()

	keys, meta, err := gitlab.CurrentUserSshKeys(nil)

	assert.NoError(t, err)

	assert.Equal(t, 3, len(keys.Items))

	assert.IsType(t, new(ResponseMeta), meta)
	assert.Equal(t, 1, meta.Page)
	assert.Equal(t, 10, meta.PerPage)
}

func TestUserSshKeys(t *testing.T) {
	ts, gitlab := mockServerFromMapping(t, "ssh_keys/user_1_ssh_keys.json")
	defer ts.Close()

	keys, meta, err := gitlab.UserSshKeys(1, nil)

	assert.NoError(t, err)

	assert.Equal(t, 3, len(keys.Items))

	assert.IsType(t, new(ResponseMeta), meta)
	assert.Equal(t, 1, meta.Page)
	assert.Equal(t, 10, meta.PerPage)
}
