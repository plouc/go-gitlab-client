package gitlab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGitlab_ResourceUrl(t *testing.T) {
	gitlab := NewGitlab("http://base_url/", "api_path", "token")

	assert.Equal(
		t,
		gitlab.ResourceUrl(ProjectsApiPath, nil).String(),
		"http://base_url/api_path/projects",
	)
	assert.Equal(
		t,
		gitlab.ResourceUrl(ProjectApiPath, map[string]string{":id": "123"}).String(),
		"http://base_url/api_path/projects/123",
	)
}
