package gogitlab

import (
	"encoding/json"
	"net/url"
)

const (
	// ID
	user_keys        = "/user/keys"     // Get current user keys
	user_key         = "/user/keys/:id" // Get user key by id
	custom_user_keys = "/user/:id/keys" // Create key for user with :id
)

type UserKey struct {
	Id    int    `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Key   string `json:"key,omitempty"`
}

func (g *Gitlab) UserKeys() ([]*UserKey, error) {

	url := g.ResourceUrl(user_keys, nil)

	var keys []*UserKey

	contents, err := g.buildAndExecRequest("GET", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &keys)
	}

	return keys, err
}

func (g *Gitlab) UserKey(id string) (*UserKey, error) {

	url := g.ResourceUrl(user_key, map[string]string{":id": id})

	var key *UserKey

	contents, err := g.buildAndExecRequest("GET", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &key)
	}

	return key, err
}

func (g *Gitlab) AddKey(title, key string) error {

	path := g.ResourceUrl(user_keys, nil)

	var err error

	v := url.Values{}
	v.Set("title", title)
	v.Set("key", key)

	body := v.Encode()

	_, err = g.buildAndExecRequest("POST", path, []byte(body))

	return err
}

func (g *Gitlab) AddUserKey(id, title, key string) error {

	path := g.ResourceUrl(user_keys, map[string]string{":id": id})

	var err error

	v := url.Values{}
	v.Set("title", title)
	v.Set("key", key)

	body := v.Encode()

	_, err = g.buildAndExecRequest("POST", path, []byte(body))

	return err
}

func (g *Gitlab) RemoveKey(id string) error {

	url := g.ResourceUrl(user_key, map[string]string{":id": id})

	var err error

	_, err = g.buildAndExecRequest("DELETE", url, nil)

	return err
}
