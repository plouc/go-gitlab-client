package gogitlab

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUserKeys(t *testing.T) {
	stub, err := ioutil.ReadFile("stubs/public_keys/index.json")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(stub))
	}))
	defer ts.Close()

	gitlab := NewGitlab(ts.URL, "", "")
	keys, err := gitlab.UserKeys()

	assert.Equal(t, err, nil)
	assert.Equal(t, len(keys), 2)
}

func TestGetUserKey(t *testing.T) {
	stub, err := ioutil.ReadFile("stubs/public_keys/show.json")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(stub))
	}))
	defer ts.Close()

	gitlab := NewGitlab(ts.URL, "", "")
	key, err := gitlab.UserKey("1")

	assert.Equal(t, err, nil)
	assert.IsType(t, new(PublicKey), key)
	assert.Equal(t, key.Title, "Public key")
	assert.Equal(t, key.Key, "ssh-rsa AAAAB3NzaC1yc2EAAAABJQAAAIEAiPWx6WM4lhHNedGfBpPJNPpZ7yKu+dnn1SJejgt4596k6YjzGGphH2TUxwKzxcKDKKezwkpfnxPkSMkuEspGRt/aZZ9wa++Oi7Qkr8prgHc4soW6NUlfDzpvZK2H5E7eQaSeP3SAwGmQKUFHCddNaP0L+hM7zhFNzjFvpaMgJw0=")
}

func TestAddKey(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(""))
	}))
	defer ts.Close()

	gitlab := NewGitlab(ts.URL, "", "")
	err := gitlab.AddKey("Public key", "stubbed key")

	assert.Equal(t, err, nil)
}

func TestAddUserKey(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(""))
	}))
	defer ts.Close()

	gitlab := NewGitlab(ts.URL, "", "")
	err := gitlab.AddUserKey("1", "Public key", "stubbed key")

	assert.Equal(t, err, nil)
}

func TestRemoveKey(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(""))
	}))
	defer ts.Close()

	gitlab := NewGitlab(ts.URL, "", "")
	err := gitlab.RemoveKey("1")

	assert.Equal(t, err, nil)
}
