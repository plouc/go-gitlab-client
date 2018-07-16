package gitlab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGitlab_Projects(t *testing.T) {
	ts, gitlab := mockServerFromMapping(t, "projects/projects.json")
	defer ts.Close()

	projects, meta, err := gitlab.Projects(nil)

	assert.NoError(t, err)

	assert.Equal(t, len(projects.Items), 2)

	assert.IsType(t, new(ResponseMeta), meta)
	assert.Equal(t, 1, meta.Page)
	assert.Equal(t, 10, meta.PerPage)
}

/*
func TestProject(t *testing.T) {
	ts, gitlab := Stub("stubs/projects/show.json")
	project, err := gitlab.Project("1")

	assert.NoError(t, err)
	assert.IsType(t, new(Project), project)
	assert.Equal(t, project.SshRepoUrl, "git@example.com:diaspora/diaspora-project-site.git")
	assert.Equal(t, project.HttpRepoUrl, "http://example.com/diaspora/diaspora-project-site.git")
	assert.Equal(t, project.WebUrl, "http://example.com/diaspora/diaspora-project-site")
	defer ts.Close()
}

func TestAddProject(t *testing.T) {
	ts, gitlab := Stub("stubs/projects/add.json")
	defer ts.Close()

	project := &Project{
		Name:        "Diaspora Client",
		Path:        "diaspora-client",
		NamespaceId: 3,
		Visibility:  VisibilityPrivate,
	}
	newProject, err := gitlab.AddProject(project)

	assert.NoError(t, err)
	assert.Equal(t, newProject.Name, project.Name)
	assert.Equal(t, newProject.Description, project.Description)
	assert.Equal(t, newProject.SshRepoUrl, "git@example.com:diaspora/diaspora-client.git")
	assert.Equal(t, newProject.HttpRepoUrl, "http://example.com/diaspora/diaspora-client.git")
	assert.Equal(t, newProject.WebUrl, "http://example.com/diaspora/diaspora-client")
}

func TestUpdateProject(t *testing.T) {
	ts, gitlab := Stub("stubs/projects/show.json")
	project := Project{
		Description: "Project Description",
	}

	_, err := gitlab.UpdateProject("1", &project)
	assert.NoError(t, err)
	defer ts.Close()
}

func TestProjectBranches(t *testing.T) {
	ts, gitlab := Stub("stubs/projects/branches/index.json")
	branches, err := gitlab.ProjectBranches("1")

	assert.NoError(t, err)
	assert.Equal(t, len(branches), 2)
	defer ts.Close()
}

func TestRemoveProject(t *testing.T) {
	ts, gitlab := Stub("stubs/projects/remove.json")
	defer ts.Close()

	result, err := gitlab.RemoveProject("1")

	assert.NoError(t, err)
	assert.Equal(t, result, true)
}
*/
