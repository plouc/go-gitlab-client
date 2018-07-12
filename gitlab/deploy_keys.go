package gitlab

import (
	"encoding/json"
	"net/url"
)

const (
	ProjectDeployKeysApiPath = "/projects/:id/keys"
	ProjectDeployKeyApiPath  = "/projects/:id/keys/:key_id"
)

/*
Get list of project deploy keys.

    GET /projects/:id/keys

Parameters:

    id The ID of a project

*/
func (g *Gitlab) ProjectDeployKeys(id string) ([]*SshKey, *ResponseMeta, error) {
	u := g.ResourceUrl(ProjectDeployKeysApiPath, map[string]string{":id": id})

	var deployKeys []*SshKey

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &deployKeys)
	}

	return deployKeys, meta, err
}

/*
Get single project deploy key.

    GET /projects/:id/keys/:key_id

Parameters:

    id    The ID of a project
    keyId The ID of a key

*/
func (g *Gitlab) ProjectDeployKey(id, keyId string) (*SshKey, *ResponseMeta, error) {
	u := g.ResourceUrl(ProjectDeployKeyApiPath, map[string]string{
		":id":     id,
		":key_id": keyId,
	})

	var deployKey *SshKey

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &deployKey)
	}

	return deployKey, meta, err
}

/*
Add deploy key to project.

    POST /projects/:id/keys

Parameters:

    id    The ID of a project
    title The key title
    key   The key value

*/
func (g *Gitlab) AddProjectDeployKey(id, title, key string) (*ResponseMeta, error) {
	u := g.ResourceUrl(ProjectDeployKeysApiPath, map[string]string{":id": id})

	var err error

	v := url.Values{}
	v.Set("title", title)
	v.Set("key", key)

	body := v.Encode()

	_, meta, err := g.buildAndExecRequest("POST", u.String(), []byte(body))

	return meta, err
}

/*
Remove deploy key from project

    DELETE /projects/:id/keys/:key_id

Parameters:

    id    The ID of a project
    keyId The ID of a key

*/
func (g *Gitlab) RemoveProjectDeployKey(id, keyId string) (*ResponseMeta, error) {
	u := g.ResourceUrl(ProjectDeployKeyApiPath, map[string]string{
		":id":     id,
		":key_id": keyId,
	})

	var err error

	_, meta, err := g.buildAndExecRequest("DELETE", u.String(), nil)

	return meta, err
}
