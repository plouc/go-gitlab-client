package gitlab

import (
	"encoding/json"
	"io"
)

const (
	NamespacesApiPath = "/namespaces"
	NamespaceApiPath  = "/namespaces/:id"
)

type Namespace struct {
	Id                          int    `json:"id,omitempty" yaml:"id,omitempty"`
	Name                        string `json:"name,omitempty,omitempty" yaml:"name,omitempty,omitempty"`
	Path                        string `json:"path,omitempty,omitempty" yaml:"path,omitempty,omitempty"`
	Kind                        string `json:"kind,omitempty,omitempty" yaml:"kind,omitempty,omitempty"`
	FullPath                    string `json:"full_path,omitempty,omitempty" yaml:"full_path,omitempty,omitempty"`
	ParentId                    int    `json:"parent_id,omitempty" yaml:"parent_id,omitempty"`
	MembersCountWithDescendants int    `json:"members_count_with_descendants,omitempty" yaml:"members_count_with_descendants,omitempty"`
	Plan                        string `json:"plan,omitempty" yaml:"plan,omitempty"`
	Description                 string `json:"description,omitempty" yaml:"description,omitempty"`
	OwnerId                     int    `json:"owner_id,omitempty" yaml:"owner_id,omitempty"`
	CreatedAt                   string `json:"createdAt,omitempty" yaml:"createdAt,omitempty"`
	UpdatedAt                   string `json:"updatedAt,omitempty" yaml:"updatedAt,omitempty"`
}

type NamespaceCollection struct {
	Items []*Namespace
}

type NamespacesOptions struct {
	PaginationOptions
	SortOptions

	// Returns a list of namespaces the user is authorized to see
	// based on the search criteria
	Search string `url:"search,omitempty"`
}

func (ns *Namespace) RenderJson(w io.Writer) error {
	return renderJson(w, ns)
}

func (ns *Namespace) RenderYaml(w io.Writer) error {
	return renderYaml(w, ns)
}

func (c *NamespaceCollection) RenderJson(w io.Writer) error {
	return renderJson(w, c.Items)
}

func (c *NamespaceCollection) RenderYaml(w io.Writer) error {
	return renderYaml(w, c.Items)
}

func (g *Gitlab) Namespaces(o *NamespacesOptions) (*NamespaceCollection, *ResponseMeta, error) {
	u := g.ResourceUrlQ(NamespacesApiPath, nil, o)

	collection := new(NamespaceCollection)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &collection.Items)
	}

	return collection, meta, err
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
