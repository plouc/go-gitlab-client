package hookci

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Commit struct {
	Id        string        `json:"id,omitempty"`
	Message   string        `json:"message,omitempty"`
	Timestamp string        `json:"timestamp,omitempty"`
	Url       string        `json:"url,omitempty"`
	Author    *AuthorCommit `json:"author,omitempty"`
}

type AuthorCommit struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type Repository struct {
	Name        string `json:"name,omitempty"`
	Url         string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
	Homepage    string `json:"homepage,omitempty"`
}

type Hook struct {
	Before   string      `json:"befor,omitempty"`
	After    string      `json:"after,omitempty"`
	Ref      string      `json:"ref,omitempty"`
	UserId   int64       `json:"user_id,omitempty"`
	UserName string      `json:"user_name,omitempty"`
	Repo     *Repository `json:"repository,omitempty"`
	Commits  []Commit    `json:"commits,omitempty"`
	NbCommit int64       `json:"total_commits_count,omitempty"`
}

func StringToHook(s string) (Hook, error) {
	h := Hook{}
	e := json.Unmarshal([]byte(s), &h)
	return h, e
}

func HookToString(h *Hook) (string, error) {
	s, e := json.Marshal(h)
	if e != nil {
		return "", e
	}
	return string(s), e
}

type GitlabHook struct {
	C   chan Hook
	req string
}

func New(req string) (*GitlabHook, chan Hook) {
	gh := &GitlabHook{
		C:   make(chan Hook),
		req: req,
	}
	return gh, gh.C
}

func (h GitlabHook) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.RequestURI, h.req) && r.Method == "POST" {
		hook := Hook{}
		defer r.Body.Close()
		if err := json.NewDecoder(r.Body).Decode(&hook); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		h.C <- hook
		return
	}
	rw.WriteHeader(http.StatusBadRequest)
}
