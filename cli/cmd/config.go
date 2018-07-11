package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/fatih/color"
	"gopkg.in/yaml.v2"
)

type Alias struct {
	Alias        string            `yaml:"alias"`
	ResourceType string            `yaml:"resource_type"`
	ResourceIds  map[string]string `yaml:"resource_ids"`
}

func (a *Alias) IdValues() []string {
	var ids []string
	for _, id := range a.ResourceIds {
		ids = append(ids, id)
	}

	return ids
}

func (a *Alias) IdsString() string {
	out := []string{}
	for key, id := range a.ResourceIds {
		out = append(out, fmt.Sprintf("%s: %s", key, id))
	}

	return strings.Join(out, ", ")
}

type Config struct {
	Host         string   `yaml:"host"`
	ApiPath      string   `yaml:"api_path"`
	Token        string   `yaml:"token"`
	OutputFormat string   `yaml:"output_format,omitempty"`
	Aliases      []*Alias `yaml:"aliases"`
}

func (c *Config) findAlias(alias, resourceType string) (int, *Alias) {
	for idx, a := range c.Aliases {
		if a.Alias == alias && a.ResourceType == resourceType {
			return idx, a
		}
	}

	return 0, nil
}

func (c *Config) removeAliasAt(idx int) {

}

func (c *Config) findAliasE(alias, resourceType string) (int, *Alias) {
	idx, a := c.findAlias(alias, resourceType)
	if a == nil {
		color.Red("✘ Unable to find alias for %s: %s", resourceType, alias)
		os.Exit(1)
	}

	return idx, a
}

func loadConfig(path string, isMandatory bool) *Config {
	var c *Config

	file, err := ioutil.ReadFile(path)
	if err != nil && isMandatory {
		color.Red("✘ An error occurred while loading config file:")
		fmt.Printf("  %v\n\n", err)
		os.Exit(1)
	}

	err = yaml.Unmarshal(file, &c)
	if err != nil && isMandatory {
		color.Red("✘ An error occurred while parsing config file:")
		fmt.Printf("  %v\n\n", err)
		os.Exit(1)
	}

	if c != nil && c.Aliases == nil {
		c.Aliases = []*Alias{}
	}

	return c
}

func (c *Config) Write(path string) {
	configYaml, err := yaml.Marshal(c)
	if err != nil {
		color.Red("✘ An error occurred while serializing config:")
		fmt.Printf("  %v\n\n", err)
		os.Exit(1)
	}

	err = ioutil.WriteFile(path, configYaml, 0644)
	if err != nil {
		color.Red("✘ An error occurred while writing config file:")
		fmt.Printf("  %v\n\n", err)
		os.Exit(1)
	}
}

func (c *Config) aliasIdsOrArgs(alias, resourceType string, args []string) (map[string]string, error) {
	if !isValidResourceType(resourceType) {
		return nil, fmt.Errorf("✘ Invalid resource type %s!", resourceType)
	}

	idKeys := resources[resourceType]
	requiredArgCount := len(idKeys)

	_, a := c.findAlias(alias, resourceType)
	if a == nil {
		if len(args) < requiredArgCount {
			return nil, fmt.Errorf(
				color.RedString("✘ %d argument(s) required for %s resource but got %d:\n  - %s\n"),
				requiredArgCount,
				resourceType,
				len(args),
				strings.Join(idKeys, "\n  - "),
			)
		}

		ids := map[string]string{}
		for idx, key := range idKeys {
			ids[key] = args[idx]
		}

		return ids, nil
	}

	if len(a.ResourceIds) != requiredArgCount {
		return nil, fmt.Errorf(
			color.RedString("✘ %s alias for %s seems to be corrupted,\nexpected %d id(s) but got %d, required id(s):\n  - %s\n"),
			a.Alias,
			resourceType,
			requiredArgCount,
			len(a.ResourceIds),
			strings.Join(idKeys, "\n  - "),
		)
	}

	return a.ResourceIds, nil
}
