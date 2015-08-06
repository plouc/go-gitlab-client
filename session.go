package gogitlab

import (
	"encoding/json"
	"net/url"
)

const (
	session_url = "/session" // Create a new session
)

type Session struct {
	Id               int    `json:"id,omitempty"`
	Username         string `json:"username,omitempty"`
	Email            string `json:"email,omitempty"`
	Name             string `json:"name,omitempty"`
	PrivateToken     string `json:"private_token,omitempty"`
	Blocked          bool   `json:"blocked,omitempty"`
	CreatedAt        string `json:"created_at,omitempty"`
	Bio              string `json:"bio,omitempty"`
	Skype            string `json:"skype,omitempty"`
	LinkedIn         string `json:"linkedin,omitempty"`
	Twitter          string `json:"twitter,omitempty"`
	WebsiteUrl       string `json:"website_url,omitempty"`
	DarkScheme       bool   `json:"dark_scheme,omitempty"`
	ThemeId          int    `json:"theme_id,omitempty"`
	IsAdmin          bool   `json:"is_admin,omitempty"`
	CanCreateGroup   bool   `json:"can_create_group,omitempty"`
	CanCreateTeam    bool   `json:"can_create_team,omitempty"`
	CanCreateProject bool   `json:"can_create_project,omitempty"`
}

func (g *Gitlab) NewSession(login string, email string, password string) (*Session, error) {
	path := g.ResourceUrl(session_url, nil)

	v := url.Values{}
	if login != "" {
		v.Set("login", login)
	}
	if email != "" {
		v.Set("email", email)
	}
	v.Set("password", password)

	body := v.Encode()

	var session *Session
	contents, err := g.buildAndExecRequest("POST", path, []byte(body))
	if err == nil {
		err = json.Unmarshal(contents, &session)
	}

	return session, err
}
