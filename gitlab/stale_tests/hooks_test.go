package stale_tests

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestHook(t *testing.T) {
	ts, gitlab := Stub("stubs/hooks/show.json")
	hook, _, err := gitlab.ProjectHook("1", "2")

	assert.NoError(t, err)
	assert.IsType(t, new(Hook), hook)
	assert.Equal(t, hook.Url, "http://example.com/hook")
	defer ts.Close()
}

func TestParsePushHook(t *testing.T) {
	stub, _ := ioutil.ReadFile("stubs/hooks/push.json")
	p, err := ParseHook([]byte(stub))

	assert.NoError(t, err)
	assert.IsType(t, new(HookPayload), p)
	assert.Equal(t, p.After, "da1560886d4f094c3e6c9ef40349f7d38b5d27d7")
	assert.Equal(t, p.Repository.URL, "git@localhost:diaspora.git")
	assert.Equal(t, len(p.Commits), 2)
	assert.Equal(t, p.Commits[0].Author.Email, "jordi@softcatala.org")
	assert.Equal(t, p.Commits[1].Id, "da1560886d4f094c3e6c9ef40349f7d38b5d27d7")
	assert.Equal(t, p.Branch(), "master")
	assert.Equal(t, p.Head().Message, "fixed readme")
}

func TestParseIssueHook(t *testing.T) {
	stub, _ := ioutil.ReadFile("stubs/hooks/issue.json")
	p, err := ParseHook([]byte(stub))

	assert.NoError(t, err)
	assert.Equal(t, p.ObjectKind, "issue")
	assert.Equal(t, p.ObjectAttributes.Id, 301)
}

func TestParseMergeRequestHook(t *testing.T) {
	stub, _ := ioutil.ReadFile("stubs/hooks/merge_request.json")
	p, err := ParseHook([]byte(stub))

	assert.NoError(t, err)
	assert.Equal(t, p.ObjectKind, "merge_request")
	assert.Equal(t, p.ObjectAttributes.TargetBranch, "master")
	assert.Equal(t, p.ObjectAttributes.SourceProjectId, p.ObjectAttributes.TargetProjectId)
}

func TestParsePipelineHook(t *testing.T) {
	stub, _ := ioutil.ReadFile("stubs/hooks/pipeline.json")
	p, err := ParseHook([]byte(stub))

	assert.NoError(t, err)
	assert.IsType(t, new(HookPayload), p)
	assert.Equal(t, p.ObjectAttributes.Sha, "bcbb5ec396a2c0f828686f14fac9b80b780504f2")
	assert.Equal(t, p.Project.Description, "Atque in sunt eos similique dolores voluptatem.")
	assert.Equal(t, p.Commit.Id, "bcbb5ec396a2c0f828686f14fac9b80b780504f2")
	assert.Equal(t, p.Branch(), "master")
	assert.Equal(t, p.Builds[0].Id, 380)
}
