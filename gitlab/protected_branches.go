package gitlab

import (
	"encoding/json"
	"strconv"
)

const (
	protectedBranchesUrl = "/projects/:id/protected_branches"                    // Gets a list of protected branches from a project.
	protectBranchUrl     = "/projects/:id/repository/branches/:branch/protect"   // Protects a single project repository branch.
	unprotectBranchUrl   = "/projects/:id/repository/branches/:branch/unprotect" // Unprotects a single project repository branch.
)

type AccessLevelInfo struct {
	AccessLevel            int    `json:"access_level,omitempty"`
	AccessLevelDescription string `json:"access_level_description,omitempty"`
	GroupId                int    `json:"group_id,omitempty"`
	UserId                 int    `json:"user_id,omitempty"`
}

type ProtectedBranch struct {
	Name              string             `json:"name,omitempty"`
	PushAccessLevels  []*AccessLevelInfo `json:"push_access_levels,omitempty"`
	MergeAccessLevels []*AccessLevelInfo `json:"merge_access_levels,omitempty"`
}

func (g *Gitlab) ProtectedBranches(projectId string, o *PaginationOptions) ([]*ProtectedBranch, *ResponseMeta, error) {
	u := g.ResourceUrl(protectedBranchesUrl, map[string]string{":id": projectId})

	if o != nil {
		q := u.Query()

		if o.Page != 1 {
			q.Set("page", strconv.Itoa(o.Page))
		}
		if o.PerPage != 0 {
			q.Set("per_page", strconv.Itoa(o.PerPage))
		}

		u.RawQuery = q.Encode()
	}

	var protectedBranches []*ProtectedBranch

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &protectedBranches)
	}

	return protectedBranches, meta, err
}

func (g *Gitlab) ProtectBranch(projectId, branchName string) (*ResponseMeta, error) {
	u := g.ResourceUrl(protectBranchUrl, map[string]string{
		":id":     projectId,
		":branch": branchName,
	})

	var err error

	_, meta, err := g.buildAndExecRequest("PUT", u.String(), nil)

	return meta, err
}

func (g *Gitlab) UnprotectBranch(projectId, branchName string) (*ResponseMeta, error) {
	u := g.ResourceUrl(unprotectBranchUrl, map[string]string{
		":id":     projectId,
		":branch": branchName,
	})

	var err error

	_, meta, err := g.buildAndExecRequest("PUT", u.String(), nil)

	return meta, err
}
