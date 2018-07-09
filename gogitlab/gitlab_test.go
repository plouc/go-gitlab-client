package gogitlab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResourceUrl(t *testing.T) {
	gitlab := NewGitlab("http://base_url/", "api_path", "token")

	assert.Equal(
		t,
		gitlab.ResourceUrl(projectsUrl, nil).String(),
		"http://base_url/api_path/projects",
	)
	assert.Equal(
		t,
		gitlab.ResourceUrl(projectUrl, map[string]string{":id": "123"}).String(),
		"http://base_url/api_path/projects/123",
	)
}

func TestResourceUrlRaw(t *testing.T) {
	gitlab := NewGitlab("http://base_url/", "api_path", "token")
	u, opaque := gitlab.ResourceUrlRaw(projectsUrl, map[string]string{":id": "123"})
	assert.Equal(t, u, "http://base_url/api_path/projects")
	assert.Equal(t, opaque, "//base_url/api_path/projects")

	gitlab = NewGitlab("http://base/url/", "api_path", "token")
	u, opaque = gitlab.ResourceUrlRaw(projectsUrl, nil)
	assert.Equal(t, u, "http://base/url/api_path/projects")
	assert.Equal(t, opaque, "//base/url/api_path/projects")

	u, opaque = gitlab.ResourceUrlRaw(projectsUrl, map[string]string{":id": "123"})
	assert.Equal(t, u, "http://base/url/api_path/projects")
	assert.Equal(t, opaque, "//base/url/api_path/projects")
}
