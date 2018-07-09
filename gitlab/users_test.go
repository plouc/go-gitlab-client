package gitlab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUsers(t *testing.T) {
	ts, gitlab := Stub("stubs/users/index.json")
	users, _, err := gitlab.Users(nil)

	assert.NoError(t, err)
	assert.Equal(t, len(users), 2)
	defer ts.Close()
}

func TestUser(t *testing.T) {
	ts, gitlab := Stub("stubs/users/show.json")
	user, _, err := gitlab.User("plouc")

	assert.NoError(t, err)
	assert.IsType(t, new(User), user)
	assert.Equal(t, user.Id, 6)
	assert.Equal(t, user.Username, "plouc")
	assert.Equal(t, user.Name, "Raphaël Benitte")
	assert.Equal(t, user.Bio, "")
	assert.Equal(t, user.Skype, "")
	assert.Equal(t, user.LinkedIn, "")
	assert.Equal(t, user.Twitter, "")
	assert.Equal(t, user.ThemeId, 2)
	assert.Equal(t, user.State, "active")
	assert.Equal(t, user.CreatedAt, "2001-01-01T00:00:00Z")
	defer ts.Close()
}

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
