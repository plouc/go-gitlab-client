package gogitlab

import (
	"encoding/json"
)

const (
	namespaces_url        = "/namespaces"        // Get a list of namespaces associated of the authenticated user
	namespaces_search_url = "/namespaces/:query" // Get all namespaces matching a string in their name/path
)

type nNamespace struct {
	Id       int
	Path     string
	Kind     string
	FullPath string `json:"full_path,omitempty"`
}

func namespaces(u string, g *Gitlab) ([]*nNamespace, error) {
	url := g.ResourceUrl(u, nil)

	var namespaces []*nNamespace

	contents, err := g.buildAndExecRequest("GET", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &namespaces)
	}

	return namespaces, err
}

func (g *Gitlab) Namespaces() ([]*nNamespace, error) {
	return namespaces(namespaces_url, g)
}

func (g *Gitlab) SearchNamespaces(query string) ([]*nNamespace, error) {
	url, opaque := g.ResourceUrlRaw(
		namespaces_search_url,
		map[string]string{":query": query},
	)

	var namespaces []*nNamespace

	contents, err := g.buildAndExecRequestRaw("GET", url, opaque, nil)
	if err == nil {
		err = json.Unmarshal(contents, &namespaces)
	}

	return namespaces, err
}
