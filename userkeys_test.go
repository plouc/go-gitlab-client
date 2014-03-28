package gogitlab

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserKeys(t *testing.T) {
	keys_stub, err := ioutil.ReadFile("stubs/userkeys/index.json")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(keys_stub))
	}))
	defer ts.Close()

	gitlab := NewGitlab(ts.URL, "", "")
	keys, err := gitlab.UserKeys()
	key := keys[0]

	assert.Equal(t, err, nil)
	assert.Equal(t, len(keys), 2)
	assert.Equal(t, key.Title, "git@example.com:diaspora/diaspora-client.git")
	assert.Equal(t, key.Key, "http://example.com/diaspora/diaspora-client.git")
}
