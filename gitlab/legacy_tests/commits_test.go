package legacy_tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectCommitStatuses(t *testing.T) {
	ts, gitlab := Stub("stubs/commits/statuses.json")
	defer ts.Close()

	statuses, err := gitlab.ProjectCommitStatuses("1", "18f3e63d05582537db6d183d9d557be09e1f90c8")

	assert.Nil(t, err)
	assert.Equal(t, len(statuses), 2)

	assert.Equal(t, statuses[0].Author.Username, "thedude")
	assert.Equal(t, statuses[0].Author.Name, "Jeff Lebowski")
	assert.Equal(t, statuses[0].Name, "bundler:audit")
	assert.Equal(t, statuses[0].Status, "pending")
	assert.Equal(t, statuses[0].ID, 91)
	assert.Equal(t, statuses[0].Sha, "18f3e63d05582537db6d183d9d557be09e1f90c8")
	assert.Equal(t, statuses[0].TargetURL, "https://gitlab.example.com/thedude/gitlab-ce/builds/91")

	assert.Equal(t, statuses[1].Author.Username, "thedude")
	assert.Equal(t, statuses[1].Author.Name, "Jeff Lebowski")
	assert.Equal(t, statuses[1].Name, "flay")
	assert.Equal(t, statuses[1].Status, "pending")
	assert.Equal(t, statuses[1].ID, 90)
	assert.Equal(t, statuses[1].Sha, "18f3e63d05582537db6d183d9d557be09e1f90c8")
	assert.Equal(t, statuses[1].TargetURL, "https://gitlab.example.com/thedude/gitlab-ce/builds/90")
}
