package stale_tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddIssue(t *testing.T) {
	ts, gitlab := Stub("stubs/projects/issues/post.json")
	defer ts.Close()

	req := &IssueRequest{
		Title: "Test Issue",
	}
	issue, err := gitlab.AddIssue("1", req)

	assert.NoError(t, err)
	assert.Equal(t, issue.Id, 1)
	assert.Equal(t, issue.IId, 1)
	assert.Equal(t, issue.ProjectId, 1)
	assert.Equal(t, issue.Title, "Test Issue")
	assert.Equal(t, issue.Description, "")
	assert.Equal(t, issue.Labels, []string{})
	assert.Equal(t, issue.Milestone, (*Milestone)(nil))
	assert.Equal(t, issue.Assignee, (*User)(nil))
	assert.NotEqual(t, issue.Author, (*User)(nil))
	assert.Equal(t, issue.State, "opened")
	assert.Equal(t, issue.CreatedAt, "2014-07-13T19:00:00.000Z")
	assert.Equal(t, issue.UpdatedAt, issue.CreatedAt)
}
