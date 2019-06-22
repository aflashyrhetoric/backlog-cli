package cmd

import (
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
)

var boop bool

// PromptInput ... Prompts user for an assignee and returns a best guess at an assigneeID
func PromptInput(message string) (string, error) {

	fmt.Println(message)

	validate := func(input string) error {
		if len(input) < 4 && input != "me" {
			return errors.New("too short for easy identification, please type more")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Assignee ",
		Validate: validate,
		Default:  "me",
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}
	return result, nil
}

// AskFormField
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
