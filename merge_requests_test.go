package gogitlab

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectMergeRequests(t *testing.T) {
	ts, gitlab := Stub("stubs/merge_requests/index.json")
	mrs, err := gitlab.ProjectMergeRequests("3", nil)

	assert.NoError(t, err)
	assert.Equal(t, len(mrs), 1)
	defer ts.Close()
}

func TestProjectMergeRequest(t *testing.T) {
	ts, gitlab := Stub("stubs/merge_requests/show.json")
	mr, err := gitlab.ProjectMergeRequest("3", "1")

	assert.NoError(t, err)
	assert.Equal(t, mr.TargetBranch, "master")
	assert.Equal(t, mr.MergeStatus, "can_be_merged")
	assert.Equal(t, mr.SourceProjectID, 2)
	assert.Equal(t, mr.TargetProjectID, 3)
	defer ts.Close()
}

func TestProjectMergeRequestCommits(t *testing.T) {
	ts, gitlab := Stub("stubs/merge_requests/commits.json")
	commits, err := gitlab.ProjectMergeRequestCommits("3", "1")

	assert.NoError(t, err)
	assert.Equal(t, len(commits), 2)
	defer ts.Close()
}

func TestProjectMergeRequestChanges(t *testing.T) {
	ts, gitlab := Stub("stubs/merge_requests/changes.json")
	mr, err := gitlab.ProjectMergeRequestChanges("3", "1")

	assert.NoError(t, err)
	assert.Equal(t, len(mr.Changes), 1)
	defer ts.Close()
}

func TestAddMergeRequest(t *testing.T) {
	ts, gitlab := Stub("stubs/merge_requests/show.json")
	req := AddMergeRequestRequest{
		TargetProjectId: 3,
	}
	_, err := gitlab.AddMergeRequest(&req)

	assert.NoError(t, err)
	defer ts.Close()
}

func TestEditMergeRequest(t *testing.T) {
	ts, gitlab := Stub("stubs/merge_requests/show.json")
	req := MergeRequest{
		ProjectId: 3,
		Id:        1,
	}
	err := gitlab.EditMergeRequest(&req)

	assert.NoError(t, err)
	defer ts.Close()
}

func TestProjectMergeRequestAccept(t *testing.T) {
	ts, gitlab := Stub("stubs/merge_requests/show.json")
	req := AcceptMergeRequestRequest{}
	_, err := gitlab.ProjectMergeRequestAccept("3", "1", &req)
	assert.NoError(t, err)
	defer ts.Close()
}

func TestProjectMergeRequestCancelMerge(t *testing.T) {
	ts, gitlab := Stub("stubs/merge_requests/show.json")
	_, err := gitlab.ProjectMergeRequestCancelMerge("3", "1")
	assert.NoError(t, err)
	defer ts.Close()
}
