package gogitlab

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
  "id":6,
  "username":"plouc",
  "email":"plouc@plouc.com",
  "name":"Raphaël Benitte",
  "bio":null,
  "skype":"",
  "linkedin":"",
  "twitter":"",
  "theme_id":2,
  "state":"active",
  "created_at":"2001-01-01T00:00:00Z",
  "extern_uid":"uid=plouc",
  "provider":"ldap"
		}`)
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