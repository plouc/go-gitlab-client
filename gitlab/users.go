package gitlab

import (
	"encoding/json"
)

const (
	UsersApiPath       = "/users"     // Get users list
	UserApiPath        = "/users/:id" // Get a single user.
	CurrentUserApiPath = "/user"      // Get current user
)

type UserIdentity struct {
	Provider  string `json:"provider,omitempty"`
	ExternUid string `json:"extern_uid,omitempty"`
}

type User struct {
	Id               int             `json:"id,omitempty"`
	Username         string          `json:"username,omitempty"`
	Email            string          `json:"email,omitempty"`
	Name             string          `json:"name,omitempty"`
	State            string          `json:"state,omitempty"`
	AvatarUrl        string          `json:"avatar_url,omitempty"`
	WebUrl           string          `json:"web_url"`
	CreatedAt        string          `json:"created_at,omitempty"`
	IsAdmin          bool            `json:"is_admin,omitempty"`
	Bio              string          `json:"bio,omitempty"`
	Location         string          `json:"location,omitempty"`
	Skype            string          `json:"skype,omitempty"`
	LinkedIn         string          `json:"linkedin,omitempty"`
	Twitter          string          `json:"twitter,omitempty"`
	WebsiteUrl       string          `json:"website_url"`
	Organization     string          `json:"organization"`
	LastSignInAt     string          `json:"last_sign_in_at,omitempty"`
	ConfirmedAt      string          `json:"confirmed_at,omitempty"`
	ThemeId          int             `json:"theme_id,omitempty"`
	LastActivityOn   string          `json:"last_activity_on,omitempty"`
	ColorSchemeId    int             `json:"color_scheme_id,omitempty"`
	ProjectsLimit    int             `json:"projects_limit,omitempty"`
	CurrentSignInAt  string          `json:"current_sign_in_at,omitempty"`
	Identities       []*UserIdentity `json:"identities,omitempty"`
	CanCreateGroup   bool            `json:"can_create_group,omitempty"`
	CanCreateProject bool            `json:"can_create_project,omitempty"`
	TwoFactorEnabled bool            `json:"two_factor_enabled,omitempty"`
	External         bool            `json:"external,omitempty"`
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

func (g *Gitlab) Users(o *UsersOptions) ([]*User, *ResponseMeta, error) {
	u := g.ResourceUrlQ(UsersApiPath, nil, o)

	var users []*User

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &users)
	}

	return users, meta, err
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
