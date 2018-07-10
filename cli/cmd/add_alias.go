package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func init() {
	addCmd.AddCommand(aliasCmd)
}

var aliasCmd = &cobra.Command{
	Use:     "alias ALIAS RESOURCE_TYPE [...resource ids]",
	Aliases: []string{"a"},
	Short:   "Create resource alias",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("you must specify an alias")
		}

		if len(args) < 2 {
			return fmt.Errorf("you must specify a resource type")
		}
		resourceType := args[1]
		if !isValidResourceType(resourceType) {
			return errors.New(color.RedString("'%s' is not a valid resource type, must be one of:\n- %s\n", resourceType, strings.Join(resourceTypes, "\n- ")))
		}

		resourceTypeIds := resources[resourceType]
		if len(args) < 2+len(resourceTypeIds) {
			return errors.New(color.RedString("missing resource id %s requires %s\n", resourceType, strings.Join(resourceTypeIds, ", ")))
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		alias := args[0]
		resourceType := args[1]
		resourceIds := args[2:]

		color.Yellow("Creating alias for %s => %s…", resourceType, alias)

		action := "created"

		_, a := config.findAlias(alias, resourceType)
		if a != nil {
			color.Red("✘ An alias already exists for %s! (%s - %s)", a.Alias, a.ResourceType, a.IdsString())

			prompt := promptui.Select{
				Label: "Do you want to overwrite current alias?",
				Items: []string{"no", "yes"},
			}
			_, answer, err := prompt.Run()
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			if answer == "no" {
				color.Red("  Aborted alias creation")
				return
			}

			action = "updated"
		}

		idsMap := map[string]string{}
		for idx, key := range resources[resourceType] {
			idsMap[key] = resourceIds[idx]
		}

		config.Aliases = append(config.Aliases, &Alias{
			ResourceType: resourceType,
			ResourceIds:  idsMap,
			Alias:        alias,
		})

		config.Write(configFile)

		color.Green("✔ Alias '%s' for resource '%s' was successfully %s", alias, resourceType, action)
	},
}
