package gogitlab

import (
	"encoding/json"
	"strconv"
)

const (
	projectsUrl      = "/projects"            // Get a list of projects owned by the authenticated user
	projectUrl       = "/projects/:id"        // Get a specific project, identified by project ID or NAME
	starProjectUrl   = "/projects/:id/star"   // Stars a given project. Returns status code 304 if the project is already starred
	unstarProjectUrl = "/projects/:id/unstar" // Unstars a given project. Returns status code 304 if the project is not starred
	projectUrlEvents = "/projects/:id/events" // Get project events
)

type Visibility string

const (
	// VisibilityPrivate indicates project access must be granted explicitly for each user.
	VisibilityPrivate Visibility = "private"

	// VisibilityInternal indicates the project can be cloned by any logged in user.
	VisibilityInternal Visibility = "internal"

	// VisibilityPublic indicates the project can be cloned without any authentication.
	VisibilityPublic Visibility = "public"
)

type ProjectsOrder string

const (
	ProjectsOrderId   ProjectsOrder = "id"
	ProjectsOrderName ProjectsOrder = "name"
	ProjectsOrderPath ProjectsOrder = "path"
)

type ProjectGroupSharing struct {
	GroupId          int    `json:"group_id"`
	GroupName        string `json:"group_name"`
	GroupAccessLevel int    `json:"group_access_level"`
}

type ProjectStatistics struct {
	CommitCount      int `json:"commit_count"`
	StorageSize      int `json:"storage_size"`
	RepositorySize   int `json:"repository_size"`
	LfsObjectsSize   int `json:"lfs_objects_size"`
	JobArtifactsSize int `json:"job_artifacts_size"`
}

type MinimalProject struct {
	Id                int    `json:"id"`
	Name              string `json:"name"`
	NameWithNamespace string `json:"name_with_namespace"`
	Path              string `json:"path"`
	PathWithNamespace string `json:"path_with_namespace"`
}

type Project struct {
	MinimalProject
	Description                               string                 `json:"description"`
	DefaultBranch                             string                 `json:"default_branch"`
	Owner                                     *Member                `json:"owner"`
	Public                                    bool                   `json:"public"`
	Visibility                                Visibility             `json:"visibility"`
	IssuesEnabled                             bool                   `json:"issues_enabled"`
	OpenIssuesCount                           int                    `json:"open_issues_count"`
	MergeRequestsEnabled                      bool                   `json:"merge_requests_enabled"`
	WallEnabled                               bool                   `json:"wall_enabled"`
	WikiEnabled                               bool                   `json:"wiki_enabled"`
	CreatedAtRaw                              string                 `json:"created_at,omitempty"`
	Namespace                                 *Namespace             `json:"namespace,omitempty"`
	NamespaceId                               int                    `json:"namespace_id"`
	SshRepoUrl                                string                 `json:"ssh_url_to_repo"`
	HttpRepoUrl                               string                 `json:"http_url_to_repo"`
	WebUrl                                    string                 `json:"web_url"`
	ReadmeUrl                                 string                 `json:"readme_url"`
	SharedRunnersEnabled                      bool                   `json:"shared_runners_enabled"`
	Archived                                  bool                   `json:"archived"`
	OnlyAllowMergeIfPipelineSucceeds          bool                   `json:"only_allow_merge_if_pipeline_succeeds"`
	OnlyAllowMergeIfAllDiscussionsAreResolved bool                   `json:"only_allow_merge_if_all_discussions_are_resolved"`
	MergeMethod                               string                 `json:"merge_method"`
	TagList                                   []string               `json:"tag_list"`
	SharedWithGroups                          []*ProjectGroupSharing `json:"shared_with_groups"`
	ForksCount                                int                    `json:"forks_count"`
	StarCount                                 int                    `json:"star_count"`
	Statistics                                *ProjectStatistics     `json:"statistics"`
}

type ProjectsOptions struct {
	PaginationOptions
	Archived                 bool          // Limit by archived status
	Visibility               Visibility    // Limit by visibility public, internal, or private
	OrderBy                  ProjectsOrder // Return projects ordered by id, name, path, created_at, updated_at, or last_activity_at fields. Default is created_at
	Sort                     SortDirection // Return projects sorted in asc or desc order. Default is desc
	Search                   string        // Return list of projects matching the search criteria
	Simple                   bool          // Return only the ID, URL, name, and path of each project
	Owned                    bool          // Limit by projects owned by the current user
	Membership               bool          // Limit by projects that the current user is a member of
	Starred                  bool          // Limit by projects starred by the current user
	Statistics               bool          // Include project statistics
	WithCustomAttributes     bool          // Include custom attributes in response (admins only)
	WithIssuesEnabled        bool          // Limit by enabled issues feature
	WithMergeRequestsEnabled bool          // Limit by enabled merge requests feature
}

func (g *Gitlab) Projects(o *ProjectsOptions) ([]*Project, *ResponseMeta, error) {
	u := g.ResourceUrl(projectsUrl, nil)
	if o != nil {
		q := u.Query()

		if o.Page != 1 {
			q.Set("page", strconv.Itoa(o.Page))
		}
		if o.PerPage != 0 {
			q.Set("per_page", strconv.Itoa(o.PerPage))
		}
		if o.Archived {
			q.Set("archived", "true")
		}
		// @todo Visibility
		// @todo OrderBy
		// @todo Sort
		if o.Search != "" {
			q.Set("search", o.Search)
		}
		if o.Simple {
			q.Set("simple", "true")
		}
		if o.Owned {
			q.Set("owned", "true")
		}
		if o.Membership {
			q.Set("membership", "true")
		}
		if o.Starred {
			q.Set("starred", "true")
		}
		if o.Statistics {
			q.Set("statistics", "true")
		}
		if o.WithCustomAttributes {
			q.Set("with_custom_attributes", "true")
		}
		if o.WithIssuesEnabled {
			q.Set("with_issues_enabled", "true")
		}
		if o.WithMergeRequestsEnabled {
			q.Set("with_merge_requests_enabled", "true")
		}

		u.RawQuery = q.Encode()
	}

	var projects []*Project

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &projects)
	}

	return projects, meta, err
}

type ProjectAddPayload struct {
	Name string `json:"name"` // The name of the new project. Equals path if not provided
	Path string `json:"path"` // Repository name for new project. Generated based on name if not provided (generated lowercased with dashes)
}

func (g *Gitlab) AddProject(project *ProjectAddPayload) (*Project, *ResponseMeta, error) {
	u := g.ResourceUrl(projectsUrl, nil)

	projectJson, err := json.Marshal(project)
	if err != nil {
		return nil, nil, err
	}

	var createdProject *Project
	contents, meta, err := g.buildAndExecRequest("POST", u.String(), projectJson)
	if err == nil {
		err = json.Unmarshal(contents, &createdProject)
	}

	return createdProject, meta, err
}

func (g *Gitlab) RemoveProject(id string) (string, *ResponseMeta, error) {
	u := g.ResourceUrl(projectUrl, map[string]string{":id": id})

	var responseWithMessage *ResponseWithMessage
	contents, meta, err := g.buildAndExecRequest("DELETE", u.String(), nil)
	if err != nil {
		return "", meta, err
	}

	err = json.Unmarshal(contents, &responseWithMessage)

	return responseWithMessage.Message, meta, err
}

func (g *Gitlab) Project(id string, withStatistics bool) (*Project, *ResponseMeta, error) {
	u := g.ResourceUrl(projectUrl, map[string]string{":id": id})
	q := u.Query()

	if withStatistics {
		q.Set("statistics", "true")
	}

	u.RawQuery = q.Encode()

	var project *Project

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &project)
	}

	return project, meta, err
}

func (g *Gitlab) UpdateProject(id string, project *Project) (*Project, *ResponseMeta, error) {
	u := g.ResourceUrl(projectUrl, map[string]string{":id": id})

	encodedRequest, err := json.Marshal(project)
	if err != nil {
		return nil, nil, err
	}
	var result *Project

	contents, meta, err := g.buildAndExecRequest("PUT", u.String(), encodedRequest)
	if err == nil {
		err = json.Unmarshal(contents, &result)
	}

	return result, meta, err
}

func (g *Gitlab) StarProject(id string) (*Project, *ResponseMeta, error) {
	u := g.ResourceUrl(starProjectUrl, map[string]string{":id": id})

	contents, meta, err := g.buildAndExecRequest("POST", u.String(), nil)
	if err != nil {
		return nil, meta, err
	}

	var project *Project
	if meta.StatusCode == 201 {
		err = json.Unmarshal(contents, &project)
	}

	return project, meta, err
}

func (g *Gitlab) UnstarProject(id string) (*Project, *ResponseMeta, error) {
	u := g.ResourceUrl(unstarProjectUrl, map[string]string{":id": id})

	contents, meta, err := g.buildAndExecRequest("POST", u.String(), nil)
	if err != nil {
		return nil, meta, err
	}

	var project *Project
	if meta.StatusCode == 201 {
		err = json.Unmarshal(contents, &project)
	}

	return project, meta, err
}
