package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(completionCmd)
}

var completionCmd = &cobra.Command{
	Use:     "completion [shell]",
	Aliases: []string{"c"},
	Short:   "Output shell completion code for the specified shell (bash or zsh)",
	Long: `Output shell completion code for the specified shell (bash or zsh).
The shell code must be evaluated to provide interactive
completion of glc commands.
This can be done by sourcing it from the .bash_profile or .zshrc.
For bash you can run:

  echo "source <(kubectl completion bash)" >> ~/.bashrc
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("you must specify a shell")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		shell := args[0]

		if shell == "bash" {
			RootCmd.GenBashCompletion(output)
		} else if shell == "zsh" {
			RootCmd.GenZshCompletion(output)
		}
	},
}
