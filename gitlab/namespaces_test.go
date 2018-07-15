package gitlab

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNamespaces(t *testing.T) {
	ts, gitlab := mockServerFromMapping(t, "namespaces/namespaces.json")
	defer ts.Close()

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

func TestSearchNamespaces(t *testing.T) {
	ts, gitlab := mockServerFromMapping(t, "namespaces/namespaces_search.json")
	defer ts.Close()

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
