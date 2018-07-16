package cmd

import (
	"bytes"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(docCmd)
}

var docCmd = &cobra.Command{
	Use:   "doc",
	Short: "Generate CLI documentation in markdown format",
	Run: func(cmd *cobra.Command, args []string) {
		navBuf := new(bytes.Buffer)
		treeBuf := new(bytes.Buffer)

		navBuf.WriteString("\n\n")
		genCommandDoc(navBuf, treeBuf, RootCmd)
		navBuf.WriteString("\n\n\n")

		navBuf.WriteTo(os.Stdout)
		treeBuf.WriteTo(os.Stdout)
	},
}
