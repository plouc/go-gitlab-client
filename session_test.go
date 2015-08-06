package gogitlab

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSession(t *testing.T) {
	ts, gitlab := Stub("stubs/session/new.json")
	session, err := gitlab.NewSession("john_smith", "", "pw")

	assert.Equal(t, err, nil)
	assert.IsType(t, new(Session), session)
	assert.Equal(t, session.Id, 1)
	assert.Equal(t, session.Username, "john_smith")
	assert.Equal(t, session.Name, "John Smith")
	assert.Equal(t, session.PrivateToken, "dd34asd13as")
	assert.Equal(t, session.Blocked, false)
	assert.Equal(t, session.CreatedAt, "2012-05-23T08:00:58Z")
	assert.Equal(t, session.Bio, "")
	assert.Equal(t, session.Skype, "")
	assert.Equal(t, session.LinkedIn, "")
	assert.Equal(t, session.Twitter, "")
	assert.Equal(t, session.ThemeId, 1)
	assert.Equal(t, session.DarkScheme, false)
	assert.Equal(t, session.IsAdmin, false)
	assert.Equal(t, session.CanCreateGroup, true)
	assert.Equal(t, session.CanCreateTeam, true)
	assert.Equal(t, session.CanCreateProject, true)
	defer ts.Close()
}
