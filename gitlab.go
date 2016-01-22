// Package github implements a simple client to consume gitlab API.
package gogitlab

import (
	"bytes"
	"crypto/tls"
	"errors"
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

	url = g.BaseUrl + g.ApiPath + url

	if strings.Contains(url, "?") {
		url = url + "&private_token=" + g.Token
	} else {
		url = url + "?private_token=" + g.Token
	}

	return url
}

func (g *Gitlab) buildAndExecRequest(method, url string, body []byte) ([]byte, error) {
	return g.buildAndExecRequestEx(method, url, "", body, false)
}

func (g *Gitlab) buildAndExecRequestEx(method, rawurl, opaque string, body []byte, followNextLink bool) ([]byte, error) {

	var req *http.Request
	var err error

	nextUrl, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}

	if len(opaque) > 0 {
		nextUrl.Opaque = opaque
	}

	// Check if both body and followNextLink are set
	if body != nil && followNextLink {
		return nil, errors.New("Cannot body and followNextLink are mutually exclusive")
	}

	baseRequestPath := nextUrl.EscapedPath()
	privateToken := nextUrl.Query().Get("private_token")
	contentsBuffer := &bytes.Buffer{}
	for nextUrl != nil {
		if body != nil {
			reader := bytes.NewReader(body)
			req, err = http.NewRequest(method, nextUrl.String(), reader)
		} else {
			req, err = http.NewRequest(method, nextUrl.String(), nil)
		}
		if err != nil {
			panic("Error while building gitlab request")
		}

		resp, err := g.Client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("Client.Do error: %q", err)
		}
		defer resp.Body.Close()
		partialContents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%s", err)
		}

		if resp.StatusCode >= 400 {
			err = fmt.Errorf("*Gitlab.buildAndExecRequest failed: <%d> %s", resp.StatusCode, req.URL)
		}

		if err != nil {
			return nil, err
		}

		// Clear nextUrl
		nextUrl = nil

		// Check if we need to continue
		linkHeader := resp.Header.Get("Link")
		if followNextLink && linkHeader != "" {
			linkHeaders := strings.Split(linkHeader, ",")
			for _, link := range linkHeaders {
				// Find next link
				if strings.HasSuffix(link, "; rel=\"next\"") {
					nextRawUrl := strings.Trim(strings.TrimSuffix(link, "; rel=\"next\""), " <>")
					next, err := url.Parse(nextRawUrl)
					if err != nil {
						return nil, err
					}

					// Make sure we are targeting the same path
					if next.EscapedPath() != baseRequestPath {
						return nil, fmt.Errorf("Invalid next URL '%s' - path different (original path: %s)",
							nextRawUrl, baseRequestPath)
					}

					// Re-set private_token in next...
					queryValues := next.Query()
					queryValues.Set("private_token", privateToken)
					next.RawQuery = queryValues.Encode()
					nextUrl = next
					break
				}
			}

			// At this point nextUrl might be set. If this is the case, check for a trailing closing bracket
			// in partialContents and replace it with a comma.
			if partialContents[len(partialContents)-1] != byte(']') {
				return nil, errors.New("Cannot follow next URL: partial contents do not seem to be an array.")
			}

			// Remove leading bracket
			partialContents = bytes.TrimPrefix(partialContents, []byte("["))
			// Replace trailing closing bracket with a comma
			partialContents[len(partialContents)-1] = byte(',')
		}
		contentsBuffer.Write(partialContents)
	}

	contents := contentsBuffer.Bytes()
	if contents[0] != byte('[') {
		contents = append([]byte{'['}, contents...)
	}

	if contents[len(contents)-1] == byte(',') {
		contents[len(contents)-1] = byte(']')
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
	u = g.BaseUrl + g.ApiPath + path + "?private_token=" + g.Token
	p, err := url.Parse(u)
	if err != nil {
		return u, ""
	}
	opaque := "//" + p.Host + g.ApiPath + path
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
