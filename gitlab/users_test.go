package gitlab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGitlab_Users(t *testing.T) {
	ts, gitlab := mockServerFromMapping(t, "users/users.json")
	defer ts.Close()

	users, meta, err := gitlab.Users(nil)

	assert.NoError(t, err)

	assert.Equal(t, len(users.Items), 2)

	assert.IsType(t, new(ResponseMeta), meta)
	assert.Equal(t, 1, meta.Page)
	assert.Equal(t, 10, meta.PerPage)
}

func TestGitlab_User(t *testing.T) {
	ts, gitlab := mockServerFromMapping(t, "users/user_1.json")
	defer ts.Close()

	user, _, err := gitlab.User("plouc")

	assert.NoError(t, err)

	assert.IsType(t, new(User), user)
	assert.Equal(t, user.Id, 1)
	assert.Equal(t, user.Username, "plouc")
	assert.Equal(t, user.Name, "Raphaël Benitte")
	assert.Equal(t, user.Bio, "")
	assert.Equal(t, user.Skype, "")
	assert.Equal(t, user.LinkedIn, "")
	assert.Equal(t, user.Twitter, "")
	assert.Equal(t, user.ThemeId, 2)
	assert.Equal(t, user.State, "active")
	assert.Equal(t, user.CreatedAt, "2001-01-01T00:00:00Z")
}

/*
func TestDeleteUser(t *testing.T) {
	ts, gitlab := Stub("")
	_, err := gitlab.RemoveUser("1")

	assert.NoError(t, err)
	defer ts.Close()
}

func TestCurrentUser(t *testing.T) {
	ts, gitlab := Stub("stubs/users/current.json")
	user, _, err := gitlab.CurrentUser()

	assert.NoError(t, err)
	assert.Equal(t, user.Username, "john_smith")
	defer ts.Close()
}
*/
