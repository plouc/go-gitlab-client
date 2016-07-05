package gogitlab

import (
	"github.com/stretchr/testify/assert"
	"testing"
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

func TestCreateProject(t *testing.T) {
	ts, gitlab := Stub("stubs/projects/show.json")
	project := Project{
		Name: "Diaspora Project Site",
	}

	result, err := gitlab.CreateProject(&project)
	assert.NoError(t, err)
	assert.Equal(t, project.Name, result.Name)
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
