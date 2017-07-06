package gogitlab

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjects(t *testing.T) {
	ts, gitlab := Stub("stubs/projects/index.json")
	projects, err := gitlab.Projects()

	assert.Equal(t, err, nil)
	assert.Equal(t, len(projects), 2)
	defer ts.Close()
}

func TestProject(t *testing.T) {
	ts, gitlab := Stub("stubs/projects/show.json")
	project, err := gitlab.Project("1")

	assert.Equal(t, err, nil)
	assert.IsType(t, new(Project), project)
	assert.Equal(t, project.SshRepoUrl, "git@example.com:diaspora/diaspora-project-site.git")
	assert.Equal(t, project.HttpRepoUrl, "http://example.com/diaspora/diaspora-project-site.git")
	assert.Equal(t, project.WebUrl, "http://example.com/diaspora/diaspora-project-site")
	defer ts.Close()
}

func TestUpdateProject(t *testing.T) {
	ts, gitlab := Stub("stubs/projects/show.json")
	project := Project{
		Description: "Project Description",
	}

	_, err := gitlab.UpdateProject("1", &project)
	assert.Equal(t, err, nil)
	defer ts.Close()
}

func TestProjectBranches(t *testing.T) {
	ts, gitlab := Stub("stubs/projects/branches/index.json")
	branches, err := gitlab.ProjectBranches("1")

	assert.Equal(t, err, nil)
	assert.Equal(t, len(branches), 2)
	defer ts.Close()
}

func TestRemoveProject(t *testing.T) {
	ts, gitlab := Stub("stubs/projects/remove.json")
	defer ts.Close()

	result, err := gitlab.RemoveProject("1")

	assert.Equal(t, err, nil)
	assert.Equal(t, result, true)
}

func TestIdParameter(t *testing.T) {
	var validTests = []struct {
		project *Project
		id      string
	}{
		{&Project{Id: 7}, "7"},
		{&Project{PathWithNamespace: "my-project"}, "my-project"},
		{&Project{PathWithNamespace: "my-group/my-project"}, "my-group%2Fmy-project"},
		{&Project{Id: 7, PathWithNamespace: "my-project"}, "7"},
	}

	for _, tt := range validTests {
		id, err := tt.project.idParameter()

		assert.NoError(t, err)
		assert.Equal(t, id, tt.id)
	}

	var invalidTests = []*Project{
		{Name: "My Project"},
		{Path: "my-path"},
	}

	for _, tt := range invalidTests {
		id, err := tt.idParameter()

		assert.Error(t, err)
		assert.Equal(t, id, "")
	}
}

func TestArchiveProject(t *testing.T) {
	ts, gitlab := Stub("stubs/projects/archive.json")
	defer ts.Close()

	result, err := gitlab.ArchiveProject(&Project{Id: 7})

	assert.NoError(t, err)
	assert.Equal(t, result.Archived, true)
}

func TestUnrchiveProject(t *testing.T) {
	ts, gitlab := Stub("stubs/projects/unarchive.json")
	defer ts.Close()

	result, err := gitlab.UnarchiveProject(&Project{Id: 7})

	assert.NoError(t, err)
	assert.Equal(t, result.Archived, false)
}
