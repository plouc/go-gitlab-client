package gitlab

import (
	"encoding/json"
	"io"
)

const (
	ProjectBranchesApiPath       = "/projects/:id/repository/branches"
	ProjectBranchApiPath         = "/projects/:id/repository/branches/:branch"
	ProjectMergedBranchesApiPath = "/projects/:id/repository/merged_branches"
)

type Branch struct {
	Name               string        `json:"name,omitempty" yaml:"name,omitempty"`
	Protected          bool          `json:"protected,omitempty" yaml:"protected,omitempty"`
	Merged             bool          `json:"merged,omitempty" yaml:"merged,omitempty"`
	DevelopersCanPush  bool          `json:"developers_can_push,omitempty" yaml:"developers_can_push,omitempty"`
	DevelopersCanMerge bool          `json:"developers_can_merge,omitempty" yaml:"developers_can_merge,omitempty"`
	Commit             *BranchCommit `json:"commit,omitempty" yaml:"commit,omitempty"`
}

type BranchCommit struct {
	Id               string  `json:"id,omitempty" yaml:"id,omitempty"`
	Tree             string  `json:"tree,omitempty" yaml:"tree,omitempty"`
	AuthoredDateRaw  string  `json:"authored_date,omitempty" yaml:"authored_date,omitempty"`
	CommittedDateRaw string  `json:"committed_date,omitempty" yaml:"committed_date,omitempty"`
	Message          string  `json:"message,omitempty" yaml:"message,omitempty"`
	Author           *Person `json:"author,omitempty" yaml:"author,omitempty"`
	Committer        *Person `json:"committer,omitempty" yaml:"committer,omitempty"`
}

type BranchCollection struct {
	Items []*Branch
}

type BranchesOptions struct {
	PaginationOptions
	SortOptions

	// Return list of branches matching the search criteria
	Search string `url:"search,omitempty"`
}

func (b *Branch) RenderJson(w io.Writer) error {
	return renderJson(w, b)
}

func (b *Branch) RenderYaml(w io.Writer) error {
	return renderYaml(w, b)
}

func (c *BranchCollection) RenderJson(w io.Writer) error {
	return renderJson(w, c.Items)
}

func (c *BranchCollection) RenderYaml(w io.Writer) error {
	return renderYaml(w, c.Items)
}

func (g *Gitlab) ProjectBranches(projectId string, o *BranchesOptions) (*BranchCollection, *ResponseMeta, error) {
	u := g.ResourceUrlQ(ProjectBranchesApiPath, map[string]string{":id": projectId}, o)

	collection := new(BranchCollection)
	branches := make([]*Branch, 0)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &branches)
	}

	collection.Items = branches

	return collection, meta, err
}

func (g *Gitlab) ProjectBranch(projectId, branchName string) (*Branch, *ResponseMeta, error) {
	u := g.ResourceUrl(ProjectBranchApiPath, map[string]string{
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
	u := g.ResourceUrl(ProjectBranchesApiPath, map[string]string{":id": projectId})

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
	u := g.ResourceUrl(ProjectBranchApiPath, map[string]string{
		":id":     projectId,
		":branch": branchName,
	})

	// response does not contain any message, but some other resource do :/
	_, meta, err := g.buildAndExecRequest("DELETE", u.String(), nil)

	return meta, err
}

func (g *Gitlab) RemoveProjectMergedBranches(projectId string) (string, *ResponseMeta, error) {
	u := g.ResourceUrl(ProjectMergedBranchesApiPath, map[string]string{":id": projectId})

	var responseWithMessage *ResponseWithMessage
	contents, meta, err := g.buildAndExecRequest("DELETE", u.String(), nil)
	if err != nil {
		return "", meta, err
	}

	err = json.Unmarshal(contents, &responseWithMessage)

	return responseWithMessage.Message, meta, err
}
