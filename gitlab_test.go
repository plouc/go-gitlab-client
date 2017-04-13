package gogitlab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResourceUrl(t *testing.T) {
	gitlab := NewGitlab("http://base_url/", "api_path", "token")

	assert.Equal(t, gitlab.ResourceUrl(projects_url, nil), "http://base_url/api_path/projects?private_token=token")
	assert.Equal(t, gitlab.ResourceUrl(project_url, map[string]string{":id": "123"}), "http://base_url/api_path/projects/123?private_token=token")
}

func TestResourceUrlRaw(t *testing.T) {
	gitlab := NewGitlab("http://base_url/", "api_path", "token")
	u, opaque := gitlab.ResourceUrlRaw(projects_url, map[string]string{":id": "123"})
	assert.Equal(t, u, "http://base_url/api_path/projects?private_token=token")
	assert.Equal(t, opaque, "//base_url/api_path/projects")

	gitlab = NewGitlab("http://base/url/", "api_path", "token")
	u, opaque = gitlab.ResourceUrlRaw(projects_url, nil)
	assert.Equal(t, u, "http://base/url/api_path/projects?private_token=token")
	assert.Equal(t, opaque, "//base/url/api_path/projects")

	u, opaque = gitlab.ResourceUrlRaw(projects_url, map[string]string{":id": "123"})
	assert.Equal(t, u, "http://base/url/api_path/projects?private_token=token")
	assert.Equal(t, opaque, "//base/url/api_path/projects")
}
