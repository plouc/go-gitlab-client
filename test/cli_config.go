package test

/*
type CliConfig struct {
	Config
	Dir  string
	File string
}

func NewConfig(dir, file string) *CliConfig {
	c := &CliConfig{
		Dir:  dir,
		File: file,
	}
	c.Host = "http://wiremock:8080"
	c.ApiPath = "/api/v4"

	return c
}

func (c *CliConfig) Path() string {
	return filepath.Join(c.Dir, c.File)
}

func (c *CliConfig) Exists() bool {
	if _, err := os.Stat(c.Path()); os.IsNotExist(err) {
		return false
	}

	return true
}

func (c *CliConfig) Write() {
	c.Config.Write(c.Path())
}
*/
