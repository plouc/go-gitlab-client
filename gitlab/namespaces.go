package gitlab

import (
	"encoding/json"
)

const (
	NamespacesApiPath = "/namespaces"
	NamespaceApiPath  = "/namespaces/:id"
)

type Namespace struct {
	Id                          int    `json:"id,omitempty"`
	Name                        string `json:"name,omitempty,omitempty"`
	Path                        string `json:"path,omitempty,omitempty"`
	Kind                        string `json:"kind,omitempty,omitempty"`
	FullPath                    string `json:"full_path,omitempty,omitempty"`
	ParentId                    int    `json:"parent_id,omitempty"`
	MembersCountWithDescendants int    `json:"members_count_with_descendants,omitempty"`
	Plan                        string `json:"plan,omitempty"`
	Description                 string `json:"description,omitempty"`
	OwnerId                     int    `json:"owner_id,omitempty"`
	CreatedAt                   string `json:"createdAt,omitempty"`
	UpdatedAt                   string `json:"updatedAt,omitempty"`
}

type NamespacesOptions struct {
	PaginationOptions
	SortOptions

	// Returns a list of namespaces the user is authorized to see
	// based on the search criteria
	Search string `url:"search,omitempty"`
}

func (g *Gitlab) Namespaces(o *NamespacesOptions) ([]*Namespace, *ResponseMeta, error) {
	u := g.ResourceUrlQ(NamespacesApiPath, nil, o)

	var namespaces []*Namespace

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &namespaces)
	}

	return namespaces, meta, err
}

func (g *Gitlab) Namespace(id string) (*Namespace, *ResponseMeta, error) {
	u := g.ResourceUrl(NamespaceApiPath, map[string]string{":id": id})

	var namespace *Namespace

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &namespace)
	}

	return namespace, meta, err
}
