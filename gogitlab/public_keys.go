package gogitlab

import (
	"encoding/json"
	"net/url"
)

const (
	// ID
	user_keys        = "/user/keys"       // Get current user keys
	user_key         = "/user/keys/:id"   // Get user key by id
	list_keys        = "/users/:uid/keys" // Get keys for the user id
	custom_user_keys = "/user/:id/keys"   // Create key for user with :id
)

type PublicKey struct {
	Id           int    `json:"id,omitempty"`
	Title        string `json:"title,omitempty"`
	Key          string `json:"key,omitempty"`
	CreatedAtRaw string `json:"created_at,omitempty"`
}

func (g *Gitlab) UserKeys() ([]*PublicKey, *ResponseMeta, error) {
	u := g.ResourceUrl(user_keys, nil)

	var keys []*PublicKey

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &keys)
	}

	return keys, meta, err
}

func (g *Gitlab) ListKeys(id string) ([]*PublicKey, *ResponseMeta, error) {
	u := g.ResourceUrl(list_keys, map[string]string{":uid": id})

	var keys []*PublicKey

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &keys)
	}

	return keys, meta, err
}

func (g *Gitlab) UserKey(id string) (*PublicKey, *ResponseMeta, error) {
	u := g.ResourceUrl(user_key, map[string]string{":id": id})

	var key *PublicKey

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &key)
	}

	return key, meta, err
}

func (g *Gitlab) AddKey(title, key string) (*ResponseMeta, error) {
	u := g.ResourceUrl(user_keys, nil)

	var err error

	v := url.Values{}
	v.Set("title", title)
	v.Set("key", key)

	body := v.Encode()

	_, meta, err := g.buildAndExecRequest("POST", u.String(), []byte(body))

	return meta, err
}

func (g *Gitlab) AddUserKey(id, title, key string) (*ResponseMeta, error) {
	u := g.ResourceUrl(user_keys, map[string]string{":id": id})

	var err error

	v := url.Values{}
	v.Set("title", title)
	v.Set("key", key)

	body := v.Encode()

	_, meta, err := g.buildAndExecRequest("POST", u.String(), []byte(body))

	return meta, err
}

func (g *Gitlab) DeleteKey(id string) error {
	u := g.ResourceUrl(user_key, map[string]string{":id": id})
	var err error
	_, _, err = g.buildAndExecRequest("DELETE", u.String(), nil)
	return err
}
