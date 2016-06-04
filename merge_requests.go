package gogitlab

import (
	"encoding/json"
	"time"
)

const (
	project_url_merge_requests             = "/projects/:id/merge_requests"                                                    // Get project merge requests
	project_url_merge_request              = "/projects/:id/merge_requests/:merge_request_id"                                  // Get information about a single merge request
	project_url_merge_request_commits      = "/projects/:id/merge_requests/:merge_request_id/commits"                          // Get a list of merge request commits
	project_url_merge_request_changes      = "/projects/:id/merge_requests/:merge_request_id/changes"                          // Shows information about the merge request including its files and changes
	project_url_merge_request_merge        = "/projects/:id/merge_requests/:merge_request_id/merge"                            // Merge changes submitted with MR
	project_url_merge_request_cancel_merge = "/projects/:id/merge_requests/:merge_request_id/cancel_merge_when_build_succeeds" // Cancel Merge When Build Succeeds
	project_url_merge_request_comments     = "/projects/:id/merge_requests/:merge_request_id/comments"                         // Lists all comments associated with a merge request
)

type MergeRequest struct {
	Id              int    `json:"id,omitempty"`
	Iid             int    `json:"iid,omitempty"`
	TargetBranch    string `json:"target_branch,omitempty"`
	SourceBranch    string `json:"source_branch,omitempty"`
	ProjectId       int    `json:"project_id,omitempty"`
	Title           string `json:"title,omitempty"`
	State           string `json:"state,omitempty"`
	CreatedAt       string `json:"created_at,omitempty"`
	UpdatedAt       string `json:"updated_at,omitempty"`
	Upvotes         int    `json:"upvotes,omitempty"`
	Downvotes       int    `json:"downvotes,omitempty"`
	Author          *User  `json:"author,omitempty"`
	Assignee        *User  `json:"assignee,omitempty"`
	Description     string `json:"description,omitempty"`
	WorkInProgress  bool   `json:"work_in_progress,omitempty"`
	MergeStatus     string `json:"merge_status,omitempty"`
	SourceProjectID int    `json:"source_project_id,omitempty"`
	TargetProjectID int    `json:"target_project_id,omitempty"`
}

type ChangeItem struct {
	OldPath     string `json:"old_path,omitempty"`
	NewPath     string `json:"new_path,omitempty"`
	AMode       string `json:"a_mode,omitempty"`
	BMode       string `json:"b_mode,omitempty"`
	Diff        string `json:"diff,omitempty"`
	NewFile     bool   `json:"new_file,omitempty"`
	RenamedFile bool   `json:"renamed_file,omitempty"`
	DeletedFile bool   `json:"deleted_file,omitempty"`
}

type MergeRequestChanges struct {
	*MergeRequest
	CreatedAt       string       `json:"created_at,omitempty"`
	UpdatedAt       string       `json:"updated_at,omitempty"`
	SourceProjectId int          `json:"source_project_id,omitempty"`
	TargetProjectId int          `json:"target_project_id,omitempty"`
	Labels          []string     `json:"labels,omitempty"`
	Milestone       Milestone    `json:"milestone,omitempty"`
	Changes         []ChangeItem `json:"changes,omitempty"`
}

type AddMergeRequestRequest struct {
	SourceBranch    string   `json:"source_branch"`
	TargetBranch    string   `json:"target_branch"`
	AssigneeId      int      `json:"assignee_id,omitempty"`
	Title           string   `json:"title"`
	Description     string   `json:"description,omitempty"`
	TargetProjectId int      `json:"target_project_id,omitempty"`
	Lables          []string `json:"lables,omitempty"`
}

type AcceptMergeRequestRequest struct {
	MergeCommitMessage       string `json:"merge_commit_message,omitempty"`
	ShouldRemoveSourceBranch bool   `json:"should_remove_source_branch,omitempty"`
	MergedWhenBuildSucceeds  bool   `json:"merged_when_build_succeeds,omitempty"`
}

/*
Get list of project merge requests.

    GET /projects/:id/merge_requests

Parameters:

    id The ID of a project

Params:
	iid (optional) - Return the request having the given iid
	state (optional) - Return all requests or just those that are merged, opened or closed
	order_by (optional) - Return requests ordered by created_at or updated_at fields. Default is created_at
	sort (optional) - Return requests sorted in asc or desc order. Default is desc

*/
func (g *Gitlab) ProjectMergeRequests(id string, params map[string]string) ([]*MergeRequest, error) {
	url, opaque := g.ResourceUrlRaw(project_url_merge_requests, map[string]string{":id": id})

	for name, value := range params {
		url = url + "&" + name + "=" + value
	}

	var err error
	var mergeRequests []*MergeRequest

	contents, err := g.buildAndExecRequestRaw("GET", url, opaque, nil)
	if err != nil {
		return mergeRequests, err
	}

	err = json.Unmarshal(contents, &mergeRequests)

	return mergeRequests, err
}

/*
Get single project merge request.

    GET /projects/:id/merge_requests/:merge_request_id

Parameters:

    id               The ID of a project
    merge_request_id The ID of a merge request

*/
func (g *Gitlab) ProjectMergeRequest(id, merge_request_id string) (*MergeRequest, error) {
	url, opaque := g.ResourceUrlRaw(project_url_merge_request, map[string]string{
		":id":               id,
		":merge_request_id": merge_request_id,
	})

	var err error
	mr := new(MergeRequest)

	contents, err := g.buildAndExecRequestRaw("GET", url, opaque, nil)
	if err != nil {
		return mr, err
	}

	err = json.Unmarshal(contents, &mr)

	return mr, err
}

/*
Get a list of merge request commits.

    GET /projects/:id/merge_request/:merge_request_id/commits

Parameters:

    id               The ID of a project
    merge_request_id The ID of a merge request

*/
func (g *Gitlab) ProjectMergeRequestCommits(id, merge_request_id string) ([]*Commit, error) {
	url, opaque := g.ResourceUrlRaw(project_url_merge_request_commits, map[string]string{
		":id":               id,
		":merge_request_id": merge_request_id,
	})

	var err error
	var commits []*Commit

	contents, err := g.buildAndExecRequestRaw("GET", url, opaque, nil)
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
Get information about the merge request including its files and changes.

    GET /projects/:id/merge_request/:merge_request_id/changes

Parameters:

    id               The ID of a project
    merge_request_id The ID of a merge request

*/
func (g *Gitlab) ProjectMergeRequestChanges(id, merge_request_id string) (*MergeRequestChanges, error) {
	url, opaque := g.ResourceUrlRaw(project_url_merge_request_changes, map[string]string{
		":id":               id,
		":merge_request_id": merge_request_id,
	})

	var err error
	changes := new(MergeRequestChanges)

	contents, err := g.buildAndExecRequestRaw("GET", url, opaque, nil)
	if err != nil {
		return changes, err
	}

	err = json.Unmarshal(contents, &changes)

	return changes, err
}

/*
Creates a new merge request.

    POST /projects/:id/merge_requests

Parameters:

    id               The ID of a project

*/
func (g *Gitlab) AddMergeRequest(req *AddMergeRequestRequest) (*MergeRequest, error) {
	url, _ := g.ResourceUrlRaw(project_url_merge_requests, map[string]string{
		":id": string(req.TargetProjectId),
	})

	encodedRequest, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	data, err := g.buildAndExecRequest("POST", url, encodedRequest)
	if err != nil {
		return nil, err
	}

	mr := new(MergeRequest)
	err = json.Unmarshal(data, mr)
	if err != nil {
		panic(err)
	}
	return mr, nil
}

/*
Updates an existing merge request.

    PUT /projects/:id/merge_request/:merge_request_id

Parameters:

    id               The ID of a project

*/
func (g *Gitlab) EditMergeRequest(mr *MergeRequest) error {
	url, _ := g.ResourceUrlRaw(project_url_merge_request, map[string]string{
		":id":               string(mr.ProjectId),
		":merge_request_id": string(mr.Id),
	})

	encodedRequest, err := json.Marshal(mr)
	if err != nil {
		return err
	}

	data, err := g.buildAndExecRequest("PUT", url, encodedRequest)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, mr)
	if err != nil {
		panic(err)
	}
	return nil
}

/*
Merge changes submitted with MR.

    PUT /projects/:id/merge_request/:merge_request_id/merge

Parameters:

    id               The ID of a project
    merge_request_id The ID of a merge request

*/
func (g *Gitlab) ProjectMergeRequestAccept(id, merge_request_id string, req *AcceptMergeRequestRequest) (*MergeRequest, error) {
	url, _ := g.ResourceUrlRaw(project_url_merge_request_merge, map[string]string{
		":id":               id,
		":merge_request_id": merge_request_id,
	})

	encodedRequest, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	data, err := g.buildAndExecRequest("PUT", url, encodedRequest)
	if err != nil {
		return nil, err
	}

	mr := new(MergeRequest)
	err = json.Unmarshal(data, mr)
	if err != nil {
		panic(err)
	}
	return mr, nil
}

/*
Cancel Merge When Build Succeeds.

    PUT /projects/:id/merge_request/:merge_request_id/cancel_merge_when_build_succeeds

Parameters:

    id               The ID of a project
    merge_request_id The ID of a merge request

*/
func (g *Gitlab) ProjectMergeRequestCancelMerge(id, merge_request_id string) (*MergeRequest, error) {
	url, _ := g.ResourceUrlRaw(project_url_merge_request_cancel_merge, map[string]string{
		":id":               id,
		":merge_request_id": merge_request_id,
	})

	data, err := g.buildAndExecRequest("PUT", url, []byte{})
	if err != nil {
		return nil, err
	}

	mr := new(MergeRequest)
	err = json.Unmarshal(data, mr)
	if err != nil {
		panic(err)
	}
	return mr, nil
}
