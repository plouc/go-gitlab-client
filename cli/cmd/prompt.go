package cmd

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/plouc/go-gitlab-client/gitlab"
)

func confirmAction(question, abortMessage string, autoConfirm bool) bool {
	if !isInteractive || autoConfirm {
		return true
	}

	prompt := promptui.Select{
		Label: color.YellowString(question),
		Items: []string{"yes", "no"},
	}

	_, answer, err := prompt.Run()
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	if answer == "yes" {
		return true
	}

	color.Red("  %s", abortMessage)

	return false
}

func handlePaginatedResult(meta *gitlab.ResponseMeta, fetch func()) {
	if !isInteractive {
		return
	}

	if meta.PrevPage != 0 || meta.NextPage != 0 {
		actions := []string{"quit"}
		if meta.PrevPage != 0 {
			actions = append(actions, fmt.Sprintf("get previous page (%d)", meta.PrevPage))
		}
		if meta.NextPage != 0 {
			actions = append(actions, fmt.Sprintf("get next page (%d)", meta.NextPage))
		}

		prompt := promptui.Select{
			Label: color.YellowString("There's more results available, what do you want to do?"),
			Items: actions,
		}

		_, action, err := prompt.Run()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if action == "quit" {
			return
		}

		if strings.Contains(action, "previous") {
			page = meta.PrevPage
			fetch()
		}

		if strings.Contains(action, "next") {
			page = meta.NextPage
			fetch()
		}
	}
}
