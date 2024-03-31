package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func promptPassword() (string, bool) {
	validate := func(input string) error {
		if len(input) < 1 {
			return fmt.Errorf("key length cannot be zero")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Key",
		Validate: validate,
		Mask:     '*',
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", false
	}

	return result, true
}

func promptData() (string, bool) {
	validate := func(input string) error {
		if len(input) < 1 {
			return fmt.Errorf("data length cannot be zero")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Data",
		Validate: validate,
		Mask:     '*',
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", false
	}

	return result, true
}
