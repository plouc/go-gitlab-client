package gogitlab

import (
	"encoding/json"
	"strconv"
)

const (
	namespacesUrl = "/namespaces"
	namespaceUrl  = "/namespaces/:id"
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
	Search string // Returns a list of namespaces the user is authorized to see based on the search criteria
}

func (g *Gitlab) Namespaces(o *NamespacesOptions) ([]*Namespace, *ResponseMeta, error) {
	u := g.ResourceUrl(namespacesUrl, nil)

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

	var namespaces []*Namespace

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &namespaces)
	}

	return namespaces, meta, err
}

func (g *Gitlab) Namespace(id string) (*Namespace, *ResponseMeta, error) {
	u := g.ResourceUrl(namespaceUrl, map[string]string{":id": id})

	var namespace *Namespace

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &namespace)
	}

	return namespace, meta, err
}
