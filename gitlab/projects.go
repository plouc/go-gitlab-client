package gitlab

import (
	"encoding/json"
	"io"
)

const (
	ProjectsApiPath      = "/projects"            // Get a list of projects owned by the authenticated user
	ProjectApiPath       = "/projects/:id"        // Get a specific project, identified by project ID or NAME
	StarProjectApiPath   = "/projects/:id/star"   // Stars a given project. Returns status code 304 if the project is already starred
	UnstarProjectApiPath = "/projects/:id/unstar" // Unstars a given project. Returns status code 304 if the project is not starred
	ProjectEventsApiPath = "/projects/:id/events" // Get project events
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

type MinimalProject struct {
	Id                int    `json:"id" yaml:"id"`
	Name              string `json:"name" yaml:"name"`
	NameWithNamespace string `json:"name_with_namespace" yaml:"name_with_namespace"`
	Path              string `json:"path" yaml:"path"`
	PathWithNamespace string `json:"path_with_namespace" yaml:"path_with_namespace"`
}

type Project struct {
	MinimalProject
	Description                               string                 `json:"description" yaml:"description"`
	DefaultBranch                             string                 `json:"default_branch" yaml:"default_branch"`
	Owner                                     *Member                `json:"owner" yaml:"owner"`
	Public                                    bool                   `json:"public" yaml:"public"`
	Visibility                                Visibility             `json:"visibility" yaml:"visibility"`
	IssuesEnabled                             bool                   `json:"issues_enabled" yaml:"issues_enabled"`
	OpenIssuesCount                           int                    `json:"open_issues_count" yaml:"open_issues_count"`
	MergeRequestsEnabled                      bool                   `json:"merge_requests_enabled" yaml:"merge_requests_enabled"`
	WallEnabled                               bool                   `json:"wall_enabled" yaml:"wall_enabled"`
	WikiEnabled                               bool                   `json:"wiki_enabled" yaml:"wiki_enabled"`
	CreatedAtRaw                              string                 `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	Namespace                                 *Namespace             `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	NamespaceId                               int                    `json:"namespace_id" yaml:"namespace_id"`
	SshRepoUrl                                string                 `json:"ssh_url_to_repo" yaml:"ssh_url_to_repo"`
	HttpRepoUrl                               string                 `json:"http_url_to_repo" yaml:"http_url_to_repo"`
	WebUrl                                    string                 `json:"web_url" yaml:"web_url"`
	ReadmeUrl                                 string                 `json:"readme_url" yaml:"readme_url"`
	SharedRunnersEnabled                      bool                   `json:"shared_runners_enabled" yaml:"shared_runners_enabled"`
	Archived                                  bool                   `json:"archived" yaml:"archived"`
	OnlyAllowMergeIfPipelineSucceeds          bool                   `json:"only_allow_merge_if_pipeline_succeeds" yaml:"only_allow_merge_if_pipeline_succeeds"`
	OnlyAllowMergeIfAllDiscussionsAreResolved bool                   `json:"only_allow_merge_if_all_discussions_are_resolved" yaml:"only_allow_merge_if_all_discussions_are_resolved"`
	MergeMethod                               string                 `json:"merge_method" yaml:"merge_method"`
	TagList                                   []string               `json:"tag_list" yaml:"tag_list"`
	SharedWithGroups                          []*ProjectGroupSharing `json:"shared_with_groups" yaml:"shared_with_groups"`
	ForksCount                                int                    `json:"forks_count" yaml:"forks_count"`
	StarCount                                 int                    `json:"star_count" yaml:"star_count"`
	Statistics                                *ProjectStatistics     `json:"statistics" yaml:"statistics"`
}

type ProjectGroupSharing struct {
	GroupId          int    `json:"group_id" yaml:"group_id"`
	GroupName        string `json:"group_name" yaml:"group_name"`
	GroupAccessLevel int    `json:"group_access_level" yaml:"group_access_level"`
}

type ProjectStatistics struct {
	CommitCount      int `json:"commit_count" yaml:"commit_count"`
	StorageSize      int `json:"storage_size" yaml:"storage_size"`
	RepositorySize   int `json:"repository_size" yaml:"repository_size"`
	LfsObjectsSize   int `json:"lfs_objects_size" yaml:"lfs_objects_size"`
	JobArtifactsSize int `json:"job_artifacts_size" yaml:"job_artifacts_size"`
}

type ProjectCollection struct {
	Items []*Project
}

type ProjectsOptions struct {
	PaginationOptions
	SortOptions

	// Limit by archived status
	Archived bool `url:"archived,omitempty"`

	// Limit by visibility public, internal, or private
	Visibility Visibility `url:"visibility,omitempty"`

	// Return projects ordered by id, name, path, created_at, updated_at,
	// or last_activity_at fields. Default is created_at
	OrderBy ProjectsOrder `url:"order_by,omitempty"`

	// Return list of projects matching the search criteria
	Search string `url:"search,omitempty"`

	// Return only the ID, URL, name, and path of each project
	Simple bool `url:"simple,omitempty"`

	// Limit by projects owned by the current user
	Owned bool `url:"owned,omitempty"`

	// Limit by projects that the current user is a member of
	Membership bool `url:"membership,omitempty"`

	// Limit by projects starred by the current user
	Starred bool `url:"starred,omitempty"`

	// Include project statistics
	Statistics bool `url:"statistics,omitempty"`

	// Include custom attributes in response (admins only)
	WithCustomAttributes bool `url:"with_custom_attributes,omitempty"`

	// Limit by enabled issues feature
	WithIssuesEnabled bool `url:"with_issues_enabled,omitempty"`

	// Limit by enabled merge requests feature
	WithMergeRequestsEnabled bool `url:"with_merge_requests_enabled,omitempty"`
}

func (p *Project) RenderJson(w io.Writer) error {
	return renderJson(w, p)
}

func (p *Project) RenderYaml(w io.Writer) error {
	return renderYaml(w, p)
}

func (c *ProjectCollection) RenderJson(w io.Writer) error {
	return renderJson(w, c.Items)
}

func (c *ProjectCollection) RenderYaml(w io.Writer) error {
	return renderYaml(w, c.Items)
}

func (g *Gitlab) Projects(o *ProjectsOptions) (*ProjectCollection, *ResponseMeta, error) {
	u := g.ResourceUrlQ(ProjectsApiPath, nil, o)

	collection := new(ProjectCollection)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &collection.Items)
	}

	return collection, meta, err
}

type ProjectAddPayload struct {
	Name string `json:"name"` // The name of the new project. Equals path if not provided
	Path string `json:"path"` // Repository name for new project. Generated based on name if not provided (generated lowercased with dashes)
}

func (g *Gitlab) AddProject(project *ProjectAddPayload) (*Project, *ResponseMeta, error) {
	u := g.ResourceUrl(ProjectsApiPath, nil)

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
	u := g.ResourceUrl(ProjectApiPath, map[string]string{":id": id})

	var responseWithMessage *ResponseWithMessage
	contents, meta, err := g.buildAndExecRequest("DELETE", u.String(), nil)
	if err != nil {
		return "", meta, err
	}

	err = json.Unmarshal(contents, &responseWithMessage)

	return responseWithMessage.Message, meta, err
}

func (g *Gitlab) Project(id string, withStatistics bool) (*Project, *ResponseMeta, error) {
	u := g.ResourceUrl(ProjectApiPath, map[string]string{":id": id})
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
	u := g.ResourceUrl(ProjectApiPath, map[string]string{":id": id})

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
	u := g.ResourceUrl(StarProjectApiPath, map[string]string{":id": id})

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
	u := g.ResourceUrl(UnstarProjectApiPath, map[string]string{":id": id})

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
