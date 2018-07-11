package gitlab

import (
	"encoding/json"
	"time"
)

const (
	RepositoryTagsApiPath    = "/projects/:id/repository/tags"       // List project repository tags
	RepositoryCommitsApiPath = "/projects/:id/repository/commits"    // List repository commits
	RepositoryTreeApiPath    = "/projects/:id/repository/tree"       // List repository tree
	RawRepositoryFileApiPath = "/projects/:id/repository/blobs/:sha" // Get raw file content for specific commit/branch
)

type TreeNode struct {
	Name string
	Type string
	Mode string
	Id   string
}

type Tag struct {
	Name      string        `json:"name,omitempty"`
	Protected bool          `json:"protected,omitempty"`
	Commit    *BranchCommit `json:"commit,omitempty"`
}

type Commit struct {
	Id           string
	Short_Id     string
	Title        string
	Author_Name  string
	Author_Email string
	Created_At   string
	CreatedAt    time.Time
	Message      string
}

/*
Get a list of repository files and directories in a project.

    GET /projects/:id/repository/tree

Parameters:

    id (required) The ID of a project
    path (optional) The path inside repository. Used to get contend of subdirectories
		ref_name (optional) The name of a repository branch or tag or if not given the default branch

Usage:
		pass nil when not using optional parameters
*/
func (g *Gitlab) RepoTree(id, path, refName string) ([]*TreeNode, *ResponseMeta, error) {
	u := g.ResourceUrl(RepositoryTreeApiPath, map[string]string{":id": id})

	q := u.Query()
	q.Set("path", path)
	q.Set("ref_name", refName)
	u.RawQuery = q.Encode()

	var treeNodes []*TreeNode

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &treeNodes)
	}

	return treeNodes, meta, err
}

/*
Get a list of repository tags from a project, sorted by name in reverse alphabetical order.

    GET /projects/:id/repository/tags

Parameters:

    id The ID of a project

Usage:

	tags, err := gitlab.RepoTags("your_projet_id")
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, tag := range tags {
		fmt.Printf("%+v\n", tag)
	}
*/
func (g *Gitlab) RepoTags(id string) ([]*Tag, *ResponseMeta, error) {
	u := g.ResourceUrl(RepositoryTagsApiPath, map[string]string{":id": id})

	var tags []*Tag

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &tags)
	}

	return tags, meta, err
}

/*
Get a list of repository commits in a project.

    GET /projects/:id/repository/commits

Parameters:

    id      The ID of a project
	refName The name of a repository branch or tag or if not given the default branch

Usage:

	commits, err := gitlab.RepoCommits("your_projet_id")
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, commit := range commits {
		fmt.Printf("%+v\n", commit)
	}
*/
func (g *Gitlab) RepoCommits(id string) ([]*Commit, *ResponseMeta, error) {
	u := g.ResourceUrl(RepositoryCommitsApiPath, map[string]string{":id": id})

	var commits []*Commit

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &commits)
		if err == nil {
			for _, commit := range commits {
				t, _ := time.Parse(dateLayout, commit.Created_At)
				commit.CreatedAt = t
			}
		}
	}

	return commits, meta, err
}

/*
Get Raw file content
*/
func (g *Gitlab) RepoRawFile(id, sha, filepath string) ([]byte, *ResponseMeta, error) {
	u := g.ResourceUrl(RawRepositoryFileApiPath, map[string]string{
		":id":  id,
		":sha": sha,
	})

	q := u.Query()
	q.Set("filepath", filepath)
	u.RawQuery = q.Encode()

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)

	return contents, meta, err
}
