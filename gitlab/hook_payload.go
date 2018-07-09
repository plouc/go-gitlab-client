package gitlab

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type HookObjAttr struct {
	Id              int       `json:"id,omitempty"`
	Ref             string    `json:"ref,omitempty"`
	Tag             bool      `json:"tag,omitempty"`
	Sha             string    `json:"sha,omitempty"`
	BeforeSha       string    `json:"before_sha,omitempty"`
	Title           string    `json:"title,omitempty"`
	AssigneeId      int       `json:"assignee_id,omitempty"`
	AuthorId        int       `json:"author_id,omitempty"`
	ProjectId       int       `json:"project_id,omitempty"`
	Status          string    `json:"status,omitempty"`
	Stages          []string  `json:"stages,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	FinishedAt      time.Time `json:"finished_at,omitempty"`
	Duration        int       `json:"duration,omitempty"`
	Position        int       `json:"position,omitempty"`
	BranchName      string    `json:"branch_name,omitempty"`
	Description     string    `json:"description,omitempty"`
	MilestoneId     int       `json:"milestone_id,omitempty"`
	State           string    `json:"state,omitempty"`
	IId             int       `json:"iid,omitempty"`
	TargetBranch    string    `json:"target_branch,omitempty"`
	SourceBranch    string    `json:"source_branch,omitempty"`
	SourceProjectId int       `json:"source_project_id,omitempty"`
	StCommits       string    `json:"st_commits,omitempty"`
	StDiffs         string    `json:"st_diffs,omitempty"`
	MergeStatus     string    `json:"merge_status,omitempty"`
	TargetProjectId int       `json:"target_project_id,omitempty"`
}

type hRepository struct {
	Name        string `json:"name,omitempty"`
	URL         string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
	Homepage    string `json:"homepage,omitempty"`
}

type hProject struct {
	Project
	// Overwrite type *gitlab.Namespace with type string,
	// otherwise the project hash passed for pipeline hooks is
	// identical
	Namespace string `json:"namespace,omitempty"`
}

type hCommit struct {
	Id        string    `json:"id,omitempty"`
	Message   string    `json:"message,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
	URL       string    `json:"url,omitempty"`
	Author    *Person   `json:"author,omitempty"`
}

type HookPayload struct {
	Before            string       `json:"before,omitempty"`
	After             string       `json:"after,omitempty"`
	Ref               string       `json:"ref,omitempty"`
	UserId            int          `json:"user_id,omitempty"`
	UserName          string       `json:"user_name,omitempty"`
	ProjectId         int          `json:"project_id,omitempty"`
	Project           *hProject    `json:"project,omitempty"`
	Repository        *hRepository `json:"repository,omitempty"`
	Commits           []hCommit    `json:"commits,omitempty"`
	Commit            *hCommit     `json:"commit,omitempty"`
	TotalCommitsCount int          `json:"total_commits_count,omitempty"`
	ObjectKind        string       `json:"object_kind,omitempty"`
	ObjectAttributes  *HookObjAttr `json:"object_attributes,omitempty"`
	Builds            []*Build     `json:"builds,omitempty"`
}

// ParseHook parses hook payload from GitLab
func ParseHook(payload []byte) (*HookPayload, error) {
	hp := HookPayload{}
	if err := json.Unmarshal(payload, &hp); err != nil {
		return nil, err
	}

	// Basic sanity check
	switch {
	case len(hp.ObjectKind) == 0:
		// Assume this is a post-receive within repository
		if len(hp.After) == 0 {
			return nil, fmt.Errorf("Invalid hook received, commit hash not found.")
		}
	case hp.ObjectKind == "pipeline":
		fallthrough
	case hp.ObjectKind == "issue":
		fallthrough
	case hp.ObjectKind == "merge_request":
		if hp.ObjectAttributes == nil {
			return nil, fmt.Errorf("Invalid hook received, attributes not found.")
		}
	default:
		return nil, fmt.Errorf("Invalid hook received, payload format not recognized.")
	}

	return &hp, nil
}

// Branch returns current branch for pipeline and push event hook
// payload
// This function returns empty string for any other events
func (h *HookPayload) Branch() string {
	ref := h.Ref
	if h.ObjectAttributes != nil && len(h.ObjectAttributes.Ref) > 0 {
		ref = h.ObjectAttributes.Ref
	}

	return strings.TrimPrefix(ref, "refs/heads/")
}

// Head returns the latest changeset for push event hook payload
func (h *HookPayload) Head() hCommit {
	c := hCommit{}
	for _, cm := range h.Commits {
		if h.After == cm.Id {
			return cm
		}
	}
	return c
}
