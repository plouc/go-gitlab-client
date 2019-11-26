package gitlab

import (
	"testing"

	"github.com/edison-moreland/go-gitlab-client/test"
	"github.com/stretchr/testify/assert"
)

func TestGitlab_Namespaces(t *testing.T) {
	ts := test.CreateMockServer(t, []string{
		"namespaces/namespaces",
	})
	defer ts.Close()
	gitlab := NewGitlab(ts.URL, "", "")

	namespaces, meta, err := gitlab.Namespaces(nil)

	assert.NoError(t, err)

	assert.IsType(t, new(Namespace), namespaces.Items[0])
	assert.Equal(t, 3, len(namespaces.Items))
	assert.Equal(t, 1, namespaces.Items[0].Id)
	assert.Equal(t, "group1", namespaces.Items[1].Path)

	assert.IsType(t, new(ResponseMeta), meta)
	assert.Equal(t, 1, meta.Page)
	assert.Equal(t, 10, meta.PerPage)
}

func TestGitlab_SearchNamespaces(t *testing.T) {
	ts := test.CreateMockServer(t, []string{
		"namespaces/namespaces_search",
	})
	defer ts.Close()
	gitlab := NewGitlab(ts.URL, "", "")

	namespaces, meta, err := gitlab.Namespaces(&NamespacesOptions{
		Search: "twitter",
	})

	assert.NoError(t, err)

	assert.IsType(t, new(Namespace), namespaces.Items[0])
	assert.Equal(t, 1, len(namespaces.Items))
	assert.Equal(t, 4, namespaces.Items[0].Id)
	assert.Equal(t, "twitter", namespaces.Items[0].Path)

	assert.IsType(t, new(ResponseMeta), meta)
	assert.Equal(t, 1, meta.Page)
	assert.Equal(t, 10, meta.PerPage)
}
