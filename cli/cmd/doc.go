package cmd

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

type byName []*cobra.Command

func (s byName) Len() int           { return len(s) }
func (s byName) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s byName) Less(i, j int) bool { return s[i].Name() < s[j].Name() }

func printCmdOptions(buf *bytes.Buffer, cmd *cobra.Command) error {
	flags := cmd.NonInheritedFlags()
	flags.SetOutput(buf)
	if flags.HasAvailableFlags() {
		buf.WriteString("##### Options\n\n```\n")
		flags.PrintDefaults()
		buf.WriteString("```\n\n")
	}

	parentFlags := cmd.InheritedFlags()
	parentFlags.SetOutput(buf)
	if parentFlags.HasAvailableFlags() {
		buf.WriteString("##### Options inherited from parent commands\n\n```\n")
		parentFlags.PrintDefaults()
		buf.WriteString("```\n\n")
	}

	return nil
}

func cmdHasSeeAlso(cmd *cobra.Command) bool {
	if cmd.HasParent() && len(strings.Split(cmd.Parent().CommandPath(), " ")) > 1 {
		return true
	}
	for _, c := range cmd.Commands() {
		if !c.IsAvailableCommand() || c.IsAdditionalHelpTopicCommand() {
			continue
		}
		return true
	}
	return false
}

func cmdDocLink(cmd *cobra.Command) string {
	return fmt.Sprintf(
		"[%s](#%s)",
		cmd.CommandPath(),
		strings.Replace(cmd.CommandPath(), " ", "-", -1),
	)
}

func genCommandDoc(navWriter, treeWriter io.Writer, cmd *cobra.Command) {
	name := cmd.CommandPath()
	parts := strings.Split(name, " ")
	if len(parts) > 1 {
		navBuf := new(bytes.Buffer)
		treeBuf := new(bytes.Buffer)

		cmd.InitDefaultHelpCmd()
		cmd.InitDefaultHelpFlag()

		short := cmd.Short
		long := cmd.Long
		if len(long) == 0 {
			long = short
		}

		navBuf.WriteString(fmt.Sprintf("- %s\t*%s*\n", cmdDocLink(cmd), short))

		treeBuf.WriteString(fmt.Sprintf("#### %s\n\n", name))
		treeBuf.WriteString(fmt.Sprintf("%s\n\n", short))
		treeBuf.WriteString("##### Synopsis\n\n")
		treeBuf.WriteString(long + "\n\n")

		if cmd.Runnable() {
			treeBuf.WriteString(fmt.Sprintf("```\n%s\n```\n\n", cmd.UseLine()))
		}

		if len(cmd.Example) > 0 {
			treeBuf.WriteString("##### Examples\n\n")
			treeBuf.WriteString(fmt.Sprintf("```\n%s\n```\n\n", cmd.Example))
		}

		printCmdOptions(treeBuf, cmd)

		if cmdHasSeeAlso(cmd) {
			treeBuf.WriteString("##### See also\n\n")

			if cmd.HasParent() {
				parent := cmd.Parent()
				if len(strings.Split(parent.CommandPath(), " ")) > 1 {
					treeBuf.WriteString(fmt.Sprintf("- %s\t*%s*\n", cmdDocLink(parent), parent.Short))
					cmd.VisitParents(func(c *cobra.Command) {
						if c.DisableAutoGenTag {
							cmd.DisableAutoGenTag = c.DisableAutoGenTag
						}
					})
				}
			}

			children := cmd.Commands()
			sort.Sort(byName(children))
			for _, child := range children {
				if !child.IsAvailableCommand() || child.IsAdditionalHelpTopicCommand() {
					continue
				}
				treeBuf.WriteString(fmt.Sprintf("- %s\t*%s*\n", cmdDocLink(child), child.Short))
			}

			treeBuf.WriteString("\n")
		}

		treeBuf.WriteString("\n\n")

		navBuf.WriteTo(navWriter)
		treeBuf.WriteTo(treeWriter)
	}

	for _, c := range cmd.Commands() {
		if !c.IsAvailableCommand() || c.IsAdditionalHelpTopicCommand() {
			continue
		}

		genCommandDoc(navWriter, treeWriter, c)
	}
}
