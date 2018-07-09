package utils

import (
	"os"
	"path/filepath"

	"github.com/plouc/go-gitlab-client/cli/cmd"
)

type Config struct {
	cmd.Config
	Dir  string
	File string
}

func NewConfig(dir, file string) *Config {
	c := &Config{
		Dir:  dir,
		File: file,
	}
	c.Host = "http://wiremock:8080"
	c.ApiPath = "/api/v4"

	return c
}

func (c *Config) Path() string {
	return filepath.Join(c.Dir, c.File)
}

func (c *Config) Exists() bool {
	if _, err := os.Stat(c.Path()); os.IsNotExist(err) {
		return false
	}

	return true
}

func (c *Config) Write() {
	c.Config.Write(c.Path())
}
