package gogitlab

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	repo_url_branches = "/projects/:id/repository/branches"         // List repository branches
	repo_url_branch   = "/projects/:id/repository/branches/:branch" // Get a specific branch of a project.
	repo_url_tags     = "/projects/:id/repository/tags"             // List project repository tags
	repo_url_commits  = "/projects/:id/repository/commits"          // List repository commits
	repo_url_tree     = "/projects/:id/repository/tree"             // List repository tree
	repo_url_raw_file = "/projects/:id/repository/blobs/:sha"       // Get raw file content for specific commit/branch
)

/*
Get a list of repository branches from a project, sorted by name alphabetically.

    GET /projects/:id/repository/branches

Parameters:

    id The ID of a project

Usage:

	branches, err := gitlab.RepoBranches("your_projet_id")
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, branch := range branches {
		fmt.Printf("%+v\n", branch)
	}
*/
func (g *Gitlab) RepoBranches(id string) ([]*Branch, error) {

	url := g.ResourceUrl(repo_url_branches, map[string]string{ ":id": id })

	var branches []*Branch

	contents, err := g.buildAndExecRequest("GET", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &branches)
	}

	return branches, err
}

/*
Get a single project repository branch.

    GET /projects/:id/repository/branches/:branch

Parameters:

    id     The ID of a project
    branch The name of the branch

*/
func (g *Gitlab) RepoBranch(id string, refName string) {

	url := g.ResourceUrl(repo_url_branch, map[string]string{
		":id":     id,
		":branch": refName,
	})
	fmt.Println(url)

	// @todo
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
func (g *Gitlab) RepoTags(id string) ([]*Tag, error) {

	url := g.ResourceUrl(repo_url_tags, map[string]string{ ":id": id })

	var tags []*Tag

	contents, err := g.buildAndExecRequest("GET", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &tags)
	}

	return tags, err
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
func (g *Gitlab) RepoCommits(id string) ([]*Commit, error) {

	url := g.ResourceUrl(repo_url_commits, map[string]string{ ":id": id })

	var commits []*Commit

	contents, err := g.buildAndExecRequest("GET", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &commits)
		if err == nil {
			for _, commit := range commits {
				t, _ := time.Parse(dateLayout, commit.Created_At)
				commit.CreatedAt = t
			}
		}
	}

	return commits, err
}

/*
Get Raw file content
*/
func (g *Gitlab) RepoRawFile(id, sha, filepath string) ([]byte, error) {

	url := g.ResourceUrl(repo_url_raw_file, map[string]string{
		":id":  id,
		":sha": sha,
	})
	url += "&filepath=" + filepath

	var raw []byte

	contents, err := g.buildAndExecRequest("GET", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, raw)
	}

	return raw, err
}
