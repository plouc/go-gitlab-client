package gogitlab

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroups(t *testing.T) {
	ts, gitlab := Stub("stubs/groups/index.json")
	defer ts.Close()

	groups, err := gitlab.Groups(1, 5)

	assert.NoError(t, err)
	assert.Equal(t, len(groups), 2)
	assert.Equal(t, groups[0].Id, 1)
	assert.Equal(t, groups[0].Name, "Foobar Group")
	assert.Equal(t, groups[0].Path, "foo-bar")
	assert.Equal(t, groups[1].ParentId, 1)
	assert.Equal(t, groups[1].FullPath, "foo-bar/another")
}
func TestGroup(t *testing.T) {
	ts, gitlab := Stub("stubs/groups/show.json")
	defer ts.Close()

	group, err := gitlab.Group("2")

	assert.NoError(t, err)
	assert.Equal(t, group.Id, 2)
	assert.Equal(t, group.Name, "Another Group")
	assert.Equal(t, group.Path, "another")
	assert.Equal(t, group.ParentId, 1)
	assert.Equal(t, group.FullPath, "foo-bar/another")
}

func TestAddGroup(t *testing.T) {
	ts, gitlab := Stub("stubs/groups/add.json")
	defer ts.Close()

	result, err := gitlab.AddGroup(&Group{
		Name:                 "New Group",
		Path:                 "new-group",
		Description:          "A new group added for testing",
		Visibility:           "internal",
		LfsEnabled:           true,
		RequestAccessEnabled: true,
	})

	assert.NoError(t, err)
	assert.Equal(t, result.Name, "New Group")
	assert.Equal(t, result.Path, "new-group")
	assert.Equal(t, result.Description, "A new group added for testing")
	assert.Equal(t, result.Visibility, "internal")
	assert.Equal(t, result.LfsEnabled, true)
	assert.Equal(t, result.RequestAccessEnabled, true)
}

func TestUpdateGroup(t *testing.T) {
	ts, gitlab := Stub("stubs/groups/update.json")
	defer ts.Close()

	result, err := gitlab.UpdateGroup("1", &Group{
		Name:        "Updated Name",
		Description: "A new group description",
	})

	assert.NoError(t, err)
	assert.Equal(t, result.Name, "Updated Name")
	assert.Equal(t, result.Description, "A new group description")

}

func TestRemoveGroup(t *testing.T) {
	ts, gitlab := Stub("stubs/groups/remove.json")
	defer ts.Close()

	result, err := gitlab.RemoveGroup("1")

	assert.NoError(t, err)
	assert.Equal(t, result, true)
}

func TestGroupProjects(t *testing.T) {
	ts, gitlab := Stub("stubs/groups/projects/index.json")
	defer ts.Close()

	projects, err := gitlab.GroupProjects("1")

	assert.NoError(t, err)
	assert.Equal(t, len(projects), 2)
	assert.Equal(t, projects[1].Id, 6)
	assert.Equal(t, projects[1].Name, "Puppet")
	assert.Equal(t, projects[1].DefaultBranch, "master")
	assert.Equal(t, projects[1].SshRepoUrl, "git@example.com:brightbox/puppet.git")
	assert.Equal(t, projects[1].Owner.Name, "Brightbox")
}

func TestGroupMembers(t *testing.T) {
	ts, gitlab := Stub("stubs/groups/members/index.json")
	defer ts.Close()

	members, err := gitlab.GroupMembers("1")

	assert.NoError(t, err)
	assert.Equal(t, len(members), 2)
	assert.Equal(t, members[0].Id, 1)
	assert.Equal(t, members[0].Username, "raymond_smith")
	assert.Equal(t, members[0].Name, "Raymond Smith")
	assert.Equal(t, members[1].State, "active")
	assert.Equal(t, members[1].CreatedAt, "2012-10-22T14:13:35Z")
}
