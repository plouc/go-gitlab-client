// Package github implements a simple client to consume gitlab API.
package gogitlab

import (
	"bytes"
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	dashboardFeedPath = "/dashboard.atom"
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

var (
	skipCertVerify = flag.Bool("gitlab.skip-cert-check", false,
		`If set to true, gitlab client will skip certificate checking for https, possibly exposing your system to MITM attack.`)
)

func NewGitlab(baseUrl, apiPath, token string) *Gitlab {
	config := &tls.Config{InsecureSkipVerify: *skipCertVerify}
	tr := &http.Transport{
		Proxy:           http.ProxyFromEnvironment,
		TLSClientConfig: config,
	}
	client := &http.Client{Transport: tr}

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

	return g.BaseUrl + g.ApiPath + url
}

func (g *Gitlab) execRequest(method, url string, body []byte) (*http.Response, error) {
	var req *http.Request
	var err error

	if body != nil {
		reader := bytes.NewReader(body)
		req, err = http.NewRequest(method, url, reader)
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	req.Header.Add("PRIVATE-TOKEN", g.Token)

	if method == "POST" || method == "PUT" {
		req.Header.Add("Content-Type", "application/json")
	}

	if err != nil {
		panic("Error while building gitlab request")
	}

	resp, err := g.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Client.Do error: %q", err)
	}

	if resp.StatusCode >= http.StatusBadRequest {
		err = fmt.Errorf("*Gitlab.buildAndExecRequest failed: <%d> %s", resp.StatusCode, req.URL)
	}

	return resp, err
}

func (g *Gitlab) buildAndExecRequest(method, url string, body []byte) ([]byte, error) {
	resp, err := g.execRequest(method, url, body)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}

	return contents, err
}

func (g *Gitlab) ResourceUrlRaw(u string, params map[string]string) (string, string) {

	if params != nil {
		for key, val := range params {
			u = strings.Replace(u, key, val, -1)
		}
	}

	path := u
	u = g.BaseUrl + g.ApiPath + path
	p, err := url.Parse(g.BaseUrl)
	if err != nil {
		return u, ""
	}
	opaque := "//" + p.Host + p.Path + g.ApiPath + path

	return u, opaque
}

func (g *Gitlab) buildAndExecRequestRaw(method, url, opaque string, body []byte) ([]byte, error) {

	var req *http.Request
	var err error

	if body != nil {
		reader := bytes.NewReader(body)
		req, err = http.NewRequest(method, url, reader)
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	req.Header.Add("PRIVATE-TOKEN", g.Token)

	if method == "POST" || method == "PUT" {
		req.Header.Add("Content-Type", "application/json")
	}

	if err != nil {
		panic("Error while building gitlab request")
	}

	if len(opaque) > 0 {
		req.URL.Opaque = opaque
	}

	resp, err := g.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Client.Do error: %q", err)
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}

	if resp.StatusCode >= 400 {
		err = fmt.Errorf("*Gitlab.buildAndExecRequestRaw failed: <%d> %s", resp.StatusCode, req.URL)
	}

	return contents, err
}
