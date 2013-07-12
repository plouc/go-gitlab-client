package gogitlab

import (
	"strings"
	"fmt"
	"encoding/json"
)

const (
	user_url = "/users/:id" // Get a single user.
)

type User struct {
    Id            int    `json:"id,omitempty"`
    Username      string `json:"username,omitempty"`
    Email         string `json:"email,omitempty"`
    Name          string `json:"name,omitempty"`
    State         string `json:"state,omitempty"`
    CreatedAt     string `json:"created_at,omitempty"`
    Bio           string `json:"bio,omitempty"`
    Skype         string `json:"skype,omitempty"`
    LinkedIn      string `json:"linkedin,omitempty"`
    Twitter       string `json:"twitter,omitempty"`
    ExternUid     string `json:"extern_uid,omitempty"`
    Provider      string `json:"provider_name,omitempty"`
    ThemeId       int    `json:"theme_id,omitempty"`
    ColorSchemeId int    `json:"theme_id,color_scheme_id"`
}

/*
Get a single user.

    GET /users/:id

Parameters

    id The ID of a user

Usage
	
	user, err := gitlab.User("your_user_id")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v\n", user)
*/
func (g *Gitlab) User(id string) (*User, error) {
	url := strings.Replace(user_url, ":id", id, -1)
	url = g.BaseUrl + g.ApiPath + url + "?private_token=" + g.Token
	fmt.Println(url)

	contents, err := g.buildAndExecRequest("GET", url)
	if err != nil {
		fmt.Println("%s", err)
	}

	user := new(User)
	err = json.Unmarshal(contents, &user)
	if err != nil {
		fmt.Println("%s", err)
	}

	return user, err
}