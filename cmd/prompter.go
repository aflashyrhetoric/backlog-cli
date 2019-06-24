package cmd

import (
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
)

var boop bool

// PromptInput ... Prompts user for an assignee and returns a best guess at an assigneeID
func PromptInput(label, defaultValue string) (string, error) {

	validate := func(input string) error {
		if len(input) < 4 && input != "me" {
			return errors.New("too short for easy identification, please type more")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    fmt.Sprintf("%s ", label),
		Validate: validate,
		Default:  defaultValue,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}
	return result, nil
}

// PromptInput ... Prompts user for an assignee and returns a best guess at an assigneeID
func PromptInputHidden(label, defaultValue, mask string) (string, error) {

	validate := func(input string) error {
		if len(input) < 4 && input != "me" {
			return errors.New("too short for easy identification, please type more")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    fmt.Sprintf("%s ", label),
		Validate: validate,
		Default:  defaultValue,
		Mask:     0x2318,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}
	return result, nil
}

// PromptConfirm ... Prompts user for a y/n question
func PromptConfirm(message string) (bool, error) {
	fmt.Println(message)

	prompt := promptui.Prompt{
		Label:     "Assignee ",
		IsConfirm: true,
		Default:   "n",
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return false, err
	}

	return result == "y", nil
}

// AskFormField ... Requests and returns a field for a url.Values form
func AskFormField(label string) (string, error) {
	validate := func(input string) error {
		if len(input) < 4 && input != "me" {
			return errors.New("too short for easy identification, please type more")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    fmt.Sprintf("%s ", label),
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	return result, nil
}

// AssigneeSelect  ... Initiates a select prompt for the user to select  a user
func AssigneeSelect(matches []User) User {
	fmt.Print(matches)
	return User{}
}
