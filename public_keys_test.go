package gogitlab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUserKeys(t *testing.T) {
	ts, gitlab := Stub("stubs/public_keys/index.json")
	keys, err := gitlab.UserKeys()

	assert.NoError(t, err)
	assert.Equal(t, len(keys), 2)
	defer ts.Close()
}

func TestListKeys(t *testing.T) {
	ts, gitlab := Stub("stubs/public_keys/index.json")
	keys, err := gitlab.ListKeys("1")

	assert.NoError(t, err)
	assert.Equal(t, len(keys), 2)
	defer ts.Close()
}

func TestGetUserKey(t *testing.T) {
	ts, gitlab := Stub("stubs/public_keys/show.json")
	key, err := gitlab.UserKey("1")

	assert.NoError(t, err)
	assert.IsType(t, new(PublicKey), key)
	assert.Equal(t, key.Title, "Public key")
	assert.Equal(t, key.Key, "ssh-rsa AAAAB3NzaC1yc2EAAAABJQAAAIEAiPWx6WM4lhHNedGfBpPJNPpZ7yKu+dnn1SJejgt4596k6YjzGGphH2TUxwKzxcKDKKezwkpfnxPkSMkuEspGRt/aZZ9wa++Oi7Qkr8prgHc4soW6NUlfDzpvZK2H5E7eQaSeP3SAwGmQKUFHCddNaP0L+hM7zhFNzjFvpaMgJw0=")
	defer ts.Close()
}

func TestAddKey(t *testing.T) {
	ts, gitlab := Stub("")
	err := gitlab.AddKey("Public key", "stubbed key")

	assert.NoError(t, err)
	defer ts.Close()
}

func TestAddUserKey(t *testing.T) {
	ts, gitlab := Stub("")
	err := gitlab.AddUserKey("1", "Public key", "stubbed key")

	assert.NoError(t, err)
	defer ts.Close()
}

func TestDeleteKey(t *testing.T) {
	ts, gitlab := Stub("")
	err := gitlab.DeleteKey("1")

	assert.NoError(t, err)
	defer ts.Close()
}
