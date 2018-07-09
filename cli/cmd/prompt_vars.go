package cmd

import (
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/plouc/go-gitlab-client/gogitlab"
)

func promptVariable() (*gogitlab.Variable, error) {
	variable := gogitlab.Variable{}

	prompt := promptui.Prompt{
		Label: "Key",
	}
	key, err := prompt.Run()
	if err != nil {
		return nil, err
	}
	variable.Key = key

	prompt = promptui.Prompt{
		Label: "Value",
	}
	value, err := prompt.Run()
	if err != nil {
		return nil, err
	}
	variable.Value = value

	selectPrompt := promptui.Select{
		Label: color.YellowString("Protect variable?"),
		Items: []string{"yes", "no"},
	}
	idx, _, err := selectPrompt.Run()
	if err != nil {
		return nil, err
	}
	if idx == 0 {
		variable.Protected = true
	}

	prompt = promptui.Prompt{
		Label: "EnvironmentScope",
	}
	environmentScope, err := prompt.Run()
	if err != nil {
		return nil, err
	}
	variable.EnvironmentScope = environmentScope

	return &variable, nil
}
