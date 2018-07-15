package gitlab

import (
	"encoding/json"
	"io"
	"net/url"
	"strconv"
)

const (
	CurrentUserSshKeysApiPath = "/user/keys"
	UserSshKeysApiPath        = "/users/:id/keys"
	CurrentUserSshKeyApiPath  = "/user/keys/:key_id"
	UserSshKeyApiPath         = "/user/:id/keys/:key_id"
)

type SshKey struct {
	Id           int    `json:"id,omitempty" yaml:"id,omitempty"`
	Title        string `json:"title,omitempty" yaml:"title,omitempty"`
	Key          string `json:"key,omitempty" yaml:"key,omitempty"`
	CreatedAtRaw string `json:"created_at,omitempty" yaml:"created_at,omitempty"`
}

type SshKeyCollection struct {
	Items []*SshKey
}

func (s *SshKey) RenderJson(w io.Writer) error {
	return renderJson(w, s)
}

func (s *SshKey) RenderYaml(w io.Writer) error {
	return renderYaml(w, s)
}

func (c *SshKeyCollection) RenderJson(w io.Writer) error {
	return renderJson(w, c.Items)
}

func (c *SshKeyCollection) RenderYaml(w io.Writer) error {
	return renderYaml(w, c.Items)
}

func (g *Gitlab) getSshKeys(u *url.URL) (*SshKeyCollection, *ResponseMeta, error) {
	collection := new(SshKeyCollection)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &collection.Items)
	}

	return collection, meta, err
}

func (g *Gitlab) UserSshKeys(userId int, o *PaginationOptions) (*SshKeyCollection, *ResponseMeta, error) {
	u := g.ResourceUrlQ(UserSshKeysApiPath, map[string]string{
		":id": strconv.Itoa(userId),
	}, o)

	return g.getSshKeys(u)
}

func (g *Gitlab) CurrentUserSshKeys(o *PaginationOptions) (*SshKeyCollection, *ResponseMeta, error) {
	u := g.ResourceUrlQ(CurrentUserSshKeysApiPath, nil, o)

	return g.getSshKeys(u)
}

func (g *Gitlab) CurrentUserSshKey(id string) (*SshKey, *ResponseMeta, error) {
	u := g.ResourceUrl(CurrentUserSshKeyApiPath, map[string]string{":id": id})

	var key *SshKey

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &key)
	}

	return key, meta, err
}

func (g *Gitlab) addSshKey(u *url.URL, title, key string) (*ResponseMeta, error) {
	var err error

	v := url.Values{}
	v.Set("title", title)
	v.Set("key", key)

	body := v.Encode()

	_, meta, err := g.buildAndExecRequest("POST", u.String(), []byte(body))

	return meta, err
}

func (g *Gitlab) AddUserSshKey(userId, title, key string) (*ResponseMeta, error) {
	u := g.ResourceUrl(UserSshKeysApiPath, map[string]string{":id": userId})

	return g.addSshKey(u, title, key)
}

func (g *Gitlab) AddCurrentUserSshKey(title, key string) (*ResponseMeta, error) {
	u := g.ResourceUrl(CurrentUserSshKeysApiPath, nil)

	return g.addSshKey(u, title, key)

}

func (g *Gitlab) deleteSshKey(u *url.URL) (*ResponseMeta, error) {
	_, meta, err := g.buildAndExecRequest("DELETE", u.String(), nil)

	return meta, err
}

func (g *Gitlab) DeleteCurrentUserSshKey(keyId int) (*ResponseMeta, error) {
	u := g.ResourceUrl(CurrentUserSshKeyApiPath, map[string]string{
		":key_id": strconv.Itoa(keyId),
	})

	return g.deleteSshKey(u)
}

func (g *Gitlab) DeleteUserSshKey(userId, keyId int) (*ResponseMeta, error) {
	u := g.ResourceUrl(UserSshKeyApiPath, map[string]string{
		":id":     strconv.Itoa(userId),
		":key_id": strconv.Itoa(keyId),
	})

	return g.deleteSshKey(u)
}
