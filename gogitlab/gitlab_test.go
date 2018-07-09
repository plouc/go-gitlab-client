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
