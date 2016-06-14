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

func TestListVariables(t *testing.T) {
	ts, gitlab := Stub("stubs/projects/variables/index.json")

	variables, err := gitlab.ProjectVariables("1")

	assert.Equal(t, err, nil)
	assert.Equal(t, len(variables), 2)
	defer ts.Close()
}

func TestGetVariable(t *testing.T) {
	ts, gitlab := Stub("stubs/projects/variables/show.json")

	result, err := gitlab.ProjectVariable("1", "Somekey")

	assert.NoError(t, err)
	assert.Equal(t, result.Key, "somekey")
	assert.Equal(t, result.Value, "somevalue")
	defer ts.Close()
}

func TestAddVariable(t *testing.T) {
	ts, gitlab := Stub("stubs/projects/variables/show.json")
	req := Variable{
		Key:   "somekey",
		Value: "somevalue",
	}
	result, err := gitlab.AddProjectVariable("1", &req)

	assert.NoError(t, err)
	assert.Equal(t, result.Key, "somekey")
	assert.Equal(t, result.Value, "somevalue")
	defer ts.Close()
}

func TestUpdateVariable(t *testing.T) {
	ts, gitlab := Stub("stubs/projects/variables/show.json")
	req := Variable{
		Key:   "somekey",
		Value: "somevalue",
	}
	result, err := gitlab.UpdateProjectVariable("1", &req)

	assert.NoError(t, err)
	assert.Equal(t, result.Key, "somekey")
	assert.Equal(t, result.Value, "somevalue")
	defer ts.Close()
}

func TestDeleteVariable(t *testing.T) {
	ts, gitlab := Stub("stubs/projects/variables/show.json")

	result, err := gitlab.DeleteProjectVariable("1", "somekey")

	assert.NoError(t, err)
	assert.Equal(t, result.Key, "somekey")
	assert.Equal(t, result.Value, "somevalue")
	defer ts.Close()
}
