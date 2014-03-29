// Package github implements a simple client to consume gitlab API.
package gogitlab

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

const (
	dateLayout = "2006-01-02T15:04:05-07:00"
)

func NewGitlab(baseUrl, apiPath, token string) *Gitlab {

	client := &http.Client{}

	return &Gitlab{
		BaseUrl: baseUrl,
		ApiPath: apiPath,
		Token:   token,
		Client:  client,
	}
}

func (g *Gitlab) ResourceUrl(url string, params map[string]string) string {

	if params != nil {
		for key, val := range params {
			url = strings.Replace(url, key, val, -1)
		}
	}

	url = g.BaseUrl + g.ApiPath + url + "?private_token=" + g.Token

	return url
}

func (g *Gitlab) buildAndExecRequest(method, url string, body []byte) ([]byte, error) {

	var req *http.Request
	var err error

	if body != nil {
		reader := bytes.NewReader(body)
		req, err = http.NewRequest(method, url, reader)
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		panic("Error while building gitlab request")
	}

	resp, err := g.Client.Do(req)
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}

	if resp.StatusCode >= 400 {
		err = errors.New("*Gitlab.buildAndExecRequest failed: " + resp.Status)
	}

	return contents, err
}
