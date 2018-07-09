package gogitlab

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNamespaces(t *testing.T) {
	ts, gitlab := Stub("stubs/namespaces/all.json")
	defer ts.Close()
	namespaces, _, err := gitlab.Namespaces(nil)

	assert.Equal(t, nil, err)
	assert.IsType(t, new(Namespace), namespaces[0])
	assert.Equal(t, 3, len(namespaces))
	assert.Equal(t, 1, namespaces[0].Id)
	assert.Equal(t, "group1", namespaces[1].Path)
}

func TestSearchNamespaces(t *testing.T) {
	ts, gitlab := Stub("stubs/namespaces/search.json")
	defer ts.Close()
	namespaces, _, err := gitlab.Namespaces(&NamespacesOptions{
		Search: "twitter",
	})

	assert.Equal(t, err, nil)
	assert.IsType(t, new(Namespace), namespaces[0])
	assert.Equal(t, 1, len(namespaces))
	assert.Equal(t, 4, namespaces[0].Id)
	assert.Equal(t, "twitter", namespaces[0].Path)
}
