package gitlab

import (
	"encoding/json"
	"io"
)

const (
	UsersApiPath       = "/users"
	UserApiPath        = "/users/:id"
	CurrentUserApiPath = "/user"
)

type UserIdentity struct {
	Provider  string `json:"provider,omitempty" yaml:"provider,omitempty"`
	ExternUid string `json:"extern_uid,omitempty" yaml:"extern_uid,omitempty"`
}

type User struct {
	Id               int             `json:"id,omitempty" yaml:"id,omitempty"`
	Username         string          `json:"username,omitempty" yaml:"username,omitempty"`
	Email            string          `json:"email,omitempty" yaml:"email,omitempty"`
	Name             string          `json:"name,omitempty" yaml:"name,omitempty"`
	State            string          `json:"state,omitempty" yaml:"state,omitempty"`
	AvatarUrl        string          `json:"avatar_url,omitempty" yaml:"avatar_url,omitempty"`
	WebUrl           string          `json:"web_url" yaml:"web_url"`
	CreatedAt        string          `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	IsAdmin          bool            `json:"is_admin,omitempty" yaml:"is_admin,omitempty"`
	Bio              string          `json:"bio,omitempty" yaml:"bio,omitempty"`
	Location         string          `json:"location,omitempty" yaml:"location,omitempty"`
	Skype            string          `json:"skype,omitempty" yaml:"skype,omitempty"`
	LinkedIn         string          `json:"linkedin,omitempty" yaml:"linkedin,omitempty"`
	Twitter          string          `json:"twitter,omitempty" yaml:"twitter,omitempty"`
	WebsiteUrl       string          `json:"website_url" yaml:"website_url"`
	Organization     string          `json:"organization" yaml:"organization"`
	LastSignInAt     string          `json:"last_sign_in_at,omitempty" yaml:"last_sign_in_at,omitempty"`
	ConfirmedAt      string          `json:"confirmed_at,omitempty" yaml:"confirmed_at,omitempty"`
	ThemeId          int             `json:"theme_id,omitempty" yaml:"theme_id,omitempty"`
	LastActivityOn   string          `json:"last_activity_on,omitempty" yaml:"last_activity_on,omitempty"`
	ColorSchemeId    int             `json:"color_scheme_id,omitempty" yaml:"color_scheme_id,omitempty"`
	ProjectsLimit    int             `json:"projects_limit,omitempty" yaml:"projects_limit,omitempty"`
	CurrentSignInAt  string          `json:"current_sign_in_at,omitempty" yaml:"current_sign_in_at,omitempty"`
	Identities       []*UserIdentity `json:"identities,omitempty" yaml:"identities,omitempty"`
	CanCreateGroup   bool            `json:"can_create_group,omitempty" yaml:"can_create_group,omitempty"`
	CanCreateProject bool            `json:"can_create_project,omitempty" yaml:"can_create_project,omitempty"`
	TwoFactorEnabled bool            `json:"two_factor_enabled,omitempty" yaml:"two_factor_enabled,omitempty"`
	External         bool            `json:"external,omitempty" yaml:"external,omitempty"`
}

type UserCollection struct {
	Items []*User
}

type UsersOptions struct {
	PaginationOptions

	// Search users by email or username
	Search string `url:"search,omitempty"`

	// Search users by username
	Username string `url:"username,omitempty"`

	// Limit to active users
	Active bool `url:"active,omitempty"`

	// Limit to blocked users
	Blocked bool `url:"blocked,omitempty"`
}

func (u *User) RenderJson(w io.Writer) error {
	return renderJson(w, u)
}

func (u *User) RenderYaml(w io.Writer) error {
	return renderYaml(w, u)
}

func (c *UserCollection) RenderJson(w io.Writer) error {
	return renderJson(w, c.Items)
}

func (c *UserCollection) RenderYaml(w io.Writer) error {
	return renderYaml(w, c.Items)
}

func (g *Gitlab) Users(o *UsersOptions) (*UserCollection, *ResponseMeta, error) {
	u := g.ResourceUrlQ(UsersApiPath, nil, o)

	collection := new(UserCollection)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &collection.Items)
	}

	return collection, meta, err
}

func (g *Gitlab) User(id string) (*User, *ResponseMeta, error) {
	u := g.ResourceUrl(UserApiPath, map[string]string{":id": id})

	user := new(User)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &user)
	}

	return user, meta, err
}

func (g *Gitlab) CurrentUser() (*User, *ResponseMeta, error) {
	u := g.ResourceUrl(CurrentUserApiPath, nil)

	user := new(User)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &user)
	}

	return user, meta, err
}

func (g *Gitlab) RemoveUser(id string) (*ResponseMeta, error) {
	u := g.ResourceUrl(UserApiPath, map[string]string{":id": id})

	_, meta, err := g.buildAndExecRequest("DELETE", u.String(), nil)

	return meta, err
}
