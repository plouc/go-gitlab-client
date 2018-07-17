// Package github implements a simple client to consume gitlab API.
package gitlab

import (
	"bytes"
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/google/go-querystring/query"
)

const (
	DashboardFeedPath = "/dashboard.atom"
)

type Gitlab struct {
	BaseUrl      string
	ApiPath      string
	RepoFeedPath string
	Token        string
	Client       *http.Client
}

type PaginationOptions struct {
	Page    int `url:"page,omitempty"`
	PerPage int `url:"per_page,omitempty"`
}

type SortDirection string

const (
	SortDirectionAsc  SortDirection = "asc"
	SortDirectionDesc SortDirection = "desc"
)

type SortOptions struct {
	OrderBy string        `url:"order_by,omitempty"`
	Sort    SortDirection `url:"sort,omitempty"`
}

type ResponseWithMessage struct {
	Message string `json:"message"`
}

type ResponseMeta struct {
	Method     string
	Url        string
	StatusCode int
	RequestId  string
	Page       int
	PerPage    int
	PrevPage   int
	NextPage   int
	TotalPages int
	Total      int
	Runtime    float64
}

const (
	dateLayout = "2006-01-02T15:04:05-07:00"
)

var (
	skipCertVerify = flag.Bool("gitlab.skip-cert-check", false,
		`If set to true, gitlab client will skip certificate checking for https, possibly exposing your system to MITM attack.`)
)

// NewGitlab generates a new gitlab service
func NewGitlab(baseUrl, apiPath, token string) *Gitlab {
	config := &tls.Config{InsecureSkipVerify: *skipCertVerify}
	tr := &http.Transport{
		Proxy:           http.ProxyFromEnvironment,
		TLSClientConfig: config,
	}
	client := &http.Client{Transport: tr}

	if apiPath == "" {
		apiPath = "/api/v4"
	}

	return &Gitlab{
		BaseUrl: baseUrl,
		ApiPath: apiPath,
		Token:   token,
		Client:  client,
	}
}

// ResourceUrl builds an url for given resource path.
//
// It replaces path placeholders with values from `params`:
//
//   /whatever/:id => /whatever/1
//
func (g *Gitlab) ResourceUrl(path string, params map[string]string) *url.URL {
	if params != nil {
		for key, val := range params {
			path = strings.Replace(path, key, val, -1)
		}
	}

	u, err := url.Parse(g.BaseUrl + g.ApiPath + path)
	if err != nil {
		panic("Error while building gitlab url, unable to parse generated url")
	}

	return u
}

// ResourceUrlQ generates an url and appends a query string to it if available
func (g *Gitlab) ResourceUrlQ(path string, params map[string]string, qs interface{}) *url.URL {
	u := g.ResourceUrl(path, params)

	if qs != nil {
		v, err := query.Values(qs)
		if err != nil {
			panic("Error while building gitlab url, unable to set query string")
		}

		u.RawQuery = v.Encode()
	}

	return u
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

	req.Header.Add("Private-Token", g.Token)
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
		err = fmt.Errorf("*Gitlab.buildAndExecRequest failed: <%d> %s %s", resp.StatusCode, req.Method, req.URL)
	}

	return resp, err
}

func buildResponseMeta(resp *http.Response, method, u string) ResponseMeta {
	meta := ResponseMeta{}

	meta.Method = method
	meta.Url = u
	meta.StatusCode = resp.StatusCode

	requestId := resp.Header.Get("X-Request-Id")
	if requestId != "" {
		meta.RequestId = requestId
	}

	page := resp.Header.Get("X-Page")
	if page != "" {
		pageInt, err := strconv.Atoi(page)
		if err == nil {
			meta.Page = pageInt
		}
	}

	perPage := resp.Header.Get("X-Per-Page")
	if perPage != "" {
		perPageInt, err := strconv.Atoi(perPage)
		if err == nil {
			meta.PerPage = perPageInt
		}
	}

	prevPage := resp.Header.Get("X-Prev-Page")
	if prevPage != "" {
		prevPageInt, err := strconv.Atoi(prevPage)
		if err == nil {
			meta.PrevPage = prevPageInt
		}
	}

	nextPage := resp.Header.Get("X-Next-Page")
	if nextPage != "" {
		nextPageInt, err := strconv.Atoi(nextPage)
		if err == nil {
			meta.NextPage = nextPageInt
		}
	}

	totalPages := resp.Header.Get("X-Total-Pages")
	if totalPages != "" {
		totalPagesInt, err := strconv.Atoi(totalPages)
		if err == nil {
			meta.TotalPages = totalPagesInt
		}
	}

	total := resp.Header.Get("X-Total")
	if total != "" {
		totalInt, err := strconv.Atoi(total)
		if err == nil {
			meta.Total = totalInt
		}
	}

	runtime := resp.Header.Get("X-Runtime")
	if total != "" {
		runtimeFloat, err := strconv.ParseFloat(runtime, 64)
		if err == nil {
			meta.Runtime = runtimeFloat
		}
	}

	return meta
}

func (g *Gitlab) buildAndExecRequest(method, u string, body []byte) ([]byte, *ResponseMeta, error) {
	resp, err := g.execRequest(method, u, body)
	if err != nil {
		return nil, nil, err
	}

	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}

	meta := buildResponseMeta(resp, method, u)

	return contents, &meta, err
}

func (g *Gitlab) buildAndExecRequestRaw(method, u, opaque string, body []byte) ([]byte, *ResponseMeta, error) {
	var req *http.Request
	var err error

	if body != nil {
		reader := bytes.NewReader(body)
		req, err = http.NewRequest(method, u, reader)
	} else {
		req, err = http.NewRequest(method, u, nil)
	}

	req.Header.Add("Private-Token", g.Token)
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
		return nil, nil, fmt.Errorf("Client.Do error: %q", err)
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}

	if resp.StatusCode >= http.StatusBadRequest {
		err = fmt.Errorf("*Gitlab.buildAndExecRequestRaw failed: <%d> %s %s", resp.StatusCode, req.Method, req.URL)
	}

	meta := buildResponseMeta(resp, method, u)

	return contents, &meta, err
}
