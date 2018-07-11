package gitlab

import (
	"encoding/json"
	"net/url"
)

const (
	CurrentUserKeysApiPath = "/user/keys"
	CurrentUserKeyApiPath  = "/user/keys/:id"
	UserKeysApiPath        = "/users/:id/keys"
)

type PublicKey struct {
	Id           int    `json:"id,omitempty"`
	Title        string `json:"title,omitempty"`
	Key          string `json:"key,omitempty"`
	CreatedAtRaw string `json:"created_at,omitempty"`
}

func (g *Gitlab) UserKeys(userId string) ([]*PublicKey, *ResponseMeta, error) {
	u := g.ResourceUrl(UserKeysApiPath, map[string]string{":id": userId})

	var keys []*PublicKey

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &keys)
	}

	return keys, meta, err
}

func (g *Gitlab) CurrentUserKeys() ([]*PublicKey, *ResponseMeta, error) {
	u := g.ResourceUrl(CurrentUserKeysApiPath, nil)

	var keys []*PublicKey

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &keys)
	}

	return keys, meta, err
}

func (g *Gitlab) CurrentUserKey(id string) (*PublicKey, *ResponseMeta, error) {
	u := g.ResourceUrl(CurrentUserKeyApiPath, map[string]string{":id": id})

	var key *PublicKey

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &key)
	}

	return key, meta, err
}

func (g *Gitlab) addKey(u *url.URL, title, key string) (*ResponseMeta, error) {
	var err error

	v := url.Values{}
	v.Set("title", title)
	v.Set("key", key)

	body := v.Encode()

	_, meta, err := g.buildAndExecRequest("POST", u.String(), []byte(body))

	return meta, err
}

func (g *Gitlab) AddUserKey(userId, title, key string) (*ResponseMeta, error) {
	u := g.ResourceUrl(UserKeysApiPath, map[string]string{":id": userId})

	return g.addKey(u, title, key)
}

func (g *Gitlab) AddCurrentUserKey(title, key string) (*ResponseMeta, error) {
	u := g.ResourceUrl(CurrentUserKeysApiPath, nil)

	return g.addKey(u, title, key)

}

func (g *Gitlab) DeleteCurrentUserKey(id string) error {
	u := g.ResourceUrl(CurrentUserKeyApiPath, map[string]string{":id": id})

	var err error
	_, _, err = g.buildAndExecRequest("DELETE", u.String(), nil)

	return err
}
