package utils

import (
	"strings"

	"github.com/fatih/color"
	"github.com/sergi/go-diff/diffmatchpatch"
	"gopkg.in/src-d/go-git.v4/utils/diff"
)

func StringsDiff(a, b string) string {
	ds := []string{}
	diffs := diff.Do(a, b)
	for _, d := range diffs {
		if d.Type != diffmatchpatch.DiffEqual {
			if d.Type == diffmatchpatch.DiffDelete {
				ds = append(ds, color.RedString("- %s", d.Text))
			} else {
				ds = append(ds, color.RedString("+ %s", d.Text))
			}
		}
	}

	return strings.Join(ds, "")
}
