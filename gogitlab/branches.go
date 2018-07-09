package gogitlab

import (
	"encoding/json"
	"strconv"
)

const (
	projectBranchesUrl       = "/projects/:id/repository/branches"         // List repository branches.
	projectBranchUrl         = "/projects/:id/repository/branches/:branch" // Get a specific branch of a project.
	projectMergedBranchesUrl = "/projects/:id/repository/merged_branches"
)

type BranchCommit struct {
	Id               string  `json:"id,omitempty"`
	Tree             string  `json:"tree,omitempty"`
	AuthoredDateRaw  string  `json:"authored_date,omitempty"`
	CommittedDateRaw string  `json:"committed_date,omitempty"`
	Message          string  `json:"message,omitempty"`
	Author           *Person `json:"author,omitempty"`
	Committer        *Person `json:"committer,omitempty"`
}

type Branch struct {
	Name               string        `json:"name,omitempty"`
	Protected          bool          `json:"protected,omitempty"`
	Merged             bool          `json:"merged,omitempty"`
	DevelopersCanPush  bool          `json:"developers_can_push,omitempty"`
	DevelopersCanMerge bool          `json:"developers_can_merge,omitempty"`
	Commit             *BranchCommit `json:"commit,omitempty"`
}

type BranchesOptions struct {
	PaginationOptions
	Search string // Return list of branches matching the search criteria
}

func (g *Gitlab) ProjectBranches(projectId string, o *BranchesOptions) ([]*Branch, *ResponseMeta, error) {
	u := g.ResourceUrl(projectBranchesUrl, map[string]string{":id": projectId})

	if o != nil {
		q := u.Query()

		if o.Page != 1 {
			q.Set("page", strconv.Itoa(o.Page))
		}
		if o.PerPage != 0 {
			q.Set("per_page", strconv.Itoa(o.PerPage))
		}
		if o.Search != "" {
			q.Set("search", o.Search)
		}

		u.RawQuery = q.Encode()
	}

	var branches []*Branch

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &branches)
	}

	return branches, meta, err
}

func (g *Gitlab) ProjectBranch(projectId, branchName string) (*Branch, *ResponseMeta, error) {
	u := g.ResourceUrl(projectBranchUrl, map[string]string{
		":id":     projectId,
		":branch": branchName,
	})

	branch := new(Branch)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &branch)
	}

	return branch, meta, err
}

func (g *Gitlab) AddProjectBranch(projectId string, branchName, ref string) (*Branch, *ResponseMeta, error) {
	u := g.ResourceUrl(projectBranchesUrl, map[string]string{":id": projectId})
	q := u.Query()
	q.Set("branch", branchName)
	q.Set("ref", ref)
	u.RawQuery = q.Encode()

	var createdBranch *Branch
	contents, meta, err := g.buildAndExecRequest("POST", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &createdBranch)
	}

	return createdBranch, meta, err
}

func (g *Gitlab) RemoveProjectBranch(projectId, branchName string) (*ResponseMeta, error) {
	u := g.ResourceUrl(projectBranchUrl, map[string]string{
		":id":     projectId,
		":branch": branchName,
	})

	// response does not contain any message, but some other resource do :/
	_, meta, err := g.buildAndExecRequest("DELETE", u.String(), nil)

	return meta, err
}

func (g *Gitlab) RemoveProjectMergedBranches(projectId string) (string, *ResponseMeta, error) {
	u := g.ResourceUrl(projectMergedBranchesUrl, map[string]string{":id": projectId})

	var responseWithMessage *ResponseWithMessage
	contents, meta, err := g.buildAndExecRequest("DELETE", u.String(), nil)
	if err != nil {
		return "", meta, err
	}

	err = json.Unmarshal(contents, &responseWithMessage)

	return responseWithMessage.Message, meta, err
}
