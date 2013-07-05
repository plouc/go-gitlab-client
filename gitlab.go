// Package github implements a simple client to consume gitlab API.
package gogitlab

import (
	"encoding/xml"
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

type ActivityFeed struct {
	Title   string        `xml:"title"json:"title"`
	Id      string        `xml:"id"json:"id"`
	Link    []Link        `xml:"link"json:"link"`
	Updated time.Time     `xml:"updated,attr"json:"updated"`
	Entry   []*FeedCommit `xml:"entry"json:"entries"`
}

type FeedCommit struct {
	Id      string    `xml:"id"json:"id"`
	Title   string    `xml:"title"json:"title"`
	Link    []Link    `xml:"link"json:"link"`
	Updated time.Time `xml:"updated"json:"updated"`
	Author  Person    `xml:"author"json:"author"`
	Summary string    `xml:"summary"json:"summary"`
	//<media:thumbnail width="40" height="40" url="https://secure.gravatar.com/avatar/7070eab7c6206530d3b7820362227fec?s=40&amp;d=mm"/>
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

func (g *Gitlab) Activity() (ActivityFeed, error) {

	url := g.BaseUrl + dasboard_feed_path + "?private_token=" + g.Token
	fmt.Println(url)

	contents, err := g.buildAndExecRequest("GET", url)
	if err != nil {
		fmt.Println("%s", err)
	}

	var activity ActivityFeed
	err = xml.Unmarshal(contents, &activity)
	if err != nil {
		fmt.Println("%s", err)
	}

	return activity, err
}

func (g *Gitlab) RepoActivityFeed(feedPath string) ActivityFeed {

	url := g.BaseUrl + g.RepoFeedPath + "?private_token=" + g.Token

	contents, err := g.buildAndExecRequest("GET", url)
	if err != nil {
		fmt.Println("%s", err)
	}

	var activity ActivityFeed
	err = xml.Unmarshal(contents, &activity)
	if err != nil {
		fmt.Println("%s", err)
	}

	return activity
}
