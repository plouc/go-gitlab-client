package cmd

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/plouc/go-gitlab-client/gogitlab"
	"github.com/spf13/cobra"
)

type Loader interface {
	Start()
	Stop()
}

type FakeLoader struct{}

func (fl *FakeLoader) Start() {}
func (fl *FakeLoader) Stop()  {}

// global
var config *Config
var client *gogitlab.Gitlab
var loader Loader
var writers []io.Writer
var output io.Writer

// options
var configFilePath = "./.glc.yml"
var isInteractive bool
var verbose bool
var silent bool
var noColor bool
var outputFormat, outputDestination string
var page, perPage int
var currentAlias string

func cmdRequireClient(cmdName string) bool {
	if cmdName == "init" || cmdName == "version" {
		return false
	}

	return true
}

var RootCmd = &cobra.Command{
	Use:   "glc",
	Short: "gitlab cli",
	Long:  "gitlab Command Line Application",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config = loadConfig(configFilePath, cmdRequireClient(cmd.Use))

		if cmdRequireClient(cmd.Use) {
			client = gogitlab.NewGitlab(config.Host, config.ApiPath, config.Token)
		}

		if noColor {
			color.NoColor = true
		}

		if isInteractive {
			loader = spinner.New(spinner.CharSets[14], 60*time.Millisecond)
		} else {
			loader = new(FakeLoader)
		}

		if outputFormat == "" {
			if config != nil && config.OutputFormat != "" {
				outputFormat = config.OutputFormat
			} else {
				outputFormat = "text"
			}
		}

		if !silent {
			writers = append(writers, os.Stdout)
		}
		if outputDestination != "" {
			fileWriter, err := os.Create(outputDestination)
			if err != nil {
				fmt.Println(err)
				return
			}

			writers = append(writers, fileWriter)
		}
		output = io.MultiWriter(writers...)
	},
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	RootCmd.PersistentFlags().BoolVar(&silent, "silent", false, "silent mode")
	RootCmd.PersistentFlags().BoolVarP(&isInteractive, "interactive", "i", false, "enable interactive mode when applicable (eg. creation, pagination)")
	RootCmd.PersistentFlags().BoolVar(&noColor, "no-color", false, "disable color output")
	RootCmd.PersistentFlags().StringVarP(&outputFormat, "output-format", "f", "", "Output format, must be one of 'text', 'json'")
	RootCmd.PersistentFlags().StringVarP(&outputDestination, "output-destination", "o", "", "Output result to file if specified")
	RootCmd.PersistentFlags().StringVarP(&currentAlias, "alias", "a", "", "Use resource alias")
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
