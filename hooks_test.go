package gogitlab

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHook(t *testing.T) {
	stub, err := ioutil.ReadFile("stubs/hooks/show.json")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(stub))
	}))
	defer ts.Close()

	gitlab := NewGitlab(ts.URL, "", "")
	hook, err := gitlab.ProjectHook("1", "2")

	assert.Equal(t, err, nil)
	assert.IsType(t, new(Hook), hook)
	assert.Equal(t, hook.Url, "http://example.com/hook")
}
