// Package github implements a simple client to consume gitlab API.
package gogitlab

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"errors"
)

const (
	dasboard_feed_path = "/dashboard.atom"
)

type Gitlab struct {
	BaseUrl      string
	ApiPath      string
	RepoFeedPath string
	Token        string
	Client       *http.Client
}

type Owner struct {
	Id         int
	Username   string
	Email      string
	Name       string
	State      string
	Created_At string
}

type Namespace struct {
	Id          int
	Name        string
	Path        string
	Description string
	Owner_Id    int
	Created_At  string
	Updated_At  string
}

type Branch struct {
	Name      string        `json:"name,omitempty"`
	Protected bool          `json:"protected,omitempty"`
	Commit    *BranchCommit `json:"commit,omitempty"`
}

type Tag struct {
	Name      string        `json:"name,omitempty"`
	Protected bool          `json:"protected,omitempty"`
	Commit    *BranchCommit `json:"commit,omitempty"`
}

type BranchCommit struct {
	Id               string  `json:"id,omitempty"`
	Tree             string  `json:"tree,omitempty"`
	AuthoredDateRaw  string  `json:"authored_date,omitempty"`
	CommittedDateRaw string  `json:"committed_date,omitempty"`
	Message          string  `json:"message,omitempty"`
	Author           *Person `json:"author,omitempty"`
	Committer        *Person `json:"committer,omitempty"`
	/*
	"parents": [
	  {"id": "9b0c4b08e7890337fc8111e66f809c8bbec467a9"},
      {"id": "3ac634dca850cab70ab14b43ad6073d1e0a7827f"}
    ]
    */
}

type Commit struct {
	Id           string
    Short_Id     string
    Title        string
    Author_Name  string
    Author_Email string
    Created_At   string
    CreatedAt    time.Time
}

type Hook struct {
	Id           int    `json:"id,omitempty"`
	Url          string `json:"url,omitempty"`
	CreatedAtRaw string `json:"created_at,omitempty"`
}

type Link struct {
	Rel  string `xml:"rel,attr,omitempty"json:"rel"`
	Href string `xml:"href,attr"json:"href"`
}

type Person struct {
	Name  string `xml:"name"json:"name"`
	Email string `xml:"email"json:"email"`
}

const (
	dateLayout = "2006-01-02T15:04:05-07:00"
)

func NewGitlab(baseUrl string, apiPath string, token string) *Gitlab {

	client := &http.Client{}

	return &Gitlab{
		BaseUrl: baseUrl,
		ApiPath: apiPath,
		Token:   token,
		Client:  client,
	}
}

func (g *Gitlab) buildAndExecRequest(method string, url string) ([]byte, error) {

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic("Error while building gitlab request")
	}

	resp, err := g.Client.Do(req)
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}

	if (resp.StatusCode >= 400) {
		err = errors.New("*Gitlab.buildAndExecRequest failed: " + resp.Status)
	}

	return contents, err
}
