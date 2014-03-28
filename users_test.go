package gogitlab

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUsers(t *testing.T) {
	stub, err := ioutil.ReadFile("stubs/users/index.json")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(stub))
	}))
	defer ts.Close()

	gitlab := NewGitlab(ts.URL, "", "")
	users, err := gitlab.Users()

	assert.Equal(t, err, nil)
	assert.Equal(t, len(users), 2)
}

func TestUser(t *testing.T) {
	stub, err := ioutil.ReadFile("stubs/users/show.json")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(stub))
	}))
	defer ts.Close()

	gitlab := NewGitlab(ts.URL, "", "")
	user, err := gitlab.User("plouc")

	assert.Equal(t, err, nil)
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
	assert.Equal(t, user.ExternUid, "uid=plouc")
	assert.Equal(t, user.Provider, "ldap")
}

func TestDeleteUser(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(""))
	}))
	defer ts.Close()

	gitlab := NewGitlab(ts.URL, "", "")
	err := gitlab.DeleteUser("1")

	assert.Equal(t, err, nil)
}

func TestCurrentUser(t *testing.T) {
	stub, err := ioutil.ReadFile("stubs/users/current.json")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(stub))
	}))
	defer ts.Close()

	gitlab := NewGitlab(ts.URL, "", "")
	user, err := gitlab.CurrentUser()

	assert.Equal(t, err, nil)
	assert.Equal(t, user.Username, "john_smith")
}
