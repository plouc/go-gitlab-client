package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"i"},
	Short:   "Init glc config",
	Run: func(cmd *cobra.Command, args []string) {
		color.Yellow("Initializing config…")

		action := "updated"
		if config == nil {
			config = new(Config)
			action = "initialized"
		}

		hostLabel := "host"
		if config.Host != "" {
			hostLabel = fmt.Sprintf("host (leave blank to use current value %s)", config.Host)
		}
		prompt := promptui.Prompt{
			Label: hostLabel,
		}
		host, err := prompt.Run()
		if err != nil {
			color.Red("Aborting init:\n  %v\n", err)
			return
		}
		if host != "" {
			config.Host = host
		}

		apiPathLabel := "api_path (leave empty to use default)"
		if config.ApiPath != "" {
			apiPathLabel = fmt.Sprintf("api_path (leave blank to use current value %s)", config.ApiPath)
		}
		prompt = promptui.Prompt{
			Label: apiPathLabel,
		}
		apiPath, err := prompt.Run()
		if err != nil {
			fmt.Printf("Aborting init:\n  %v\n", err)
			return
		}
		if apiPath != "" {
			config.ApiPath = apiPath
		} else {
			config.ApiPath = "/api/v4"
		}

		tokenLabel := "token"
		if config.Token != "" {
			tokenLabel = "token (leave empty to use current value)"
		}
		prompt = promptui.Prompt{
			Label: tokenLabel,
		}
		token, err := prompt.Run()
		if err != nil {
			color.Red("Aborting init:\n  %v\n", err)
			return
		}
		if token != "" {
			config.Token = token
		}

		selectPrompt := promptui.Select{
			Label: "default output format",
			Items: []string{"text", "json"},
		}
		_, format, err := selectPrompt.Run()
		if err != nil {
			color.Red("Aborting init:\n  %v\n", err)
			return
		}
		config.OutputFormat = format

		config.Write(configFile)

		color.Green("✔ Configuration file was successfully %s", action)
	},
}
