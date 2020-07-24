package utils

import (
	"os"

	"github.com/AlecAivazis/survey/v2"
)

func OverwriteSource(message string, removeSourcePath string) (bool, error) {
	if !confirm(message) {
		return false, nil
	}

	err := os.RemoveAll(removeSourcePath)
	if err != nil {
		return false, err
	}

	return true, nil
}

func confirm(message string) bool {
	isConfirm := false
	prompt := &survey.Confirm{
		Message: message,
	}
	survey.AskOne(prompt, &isConfirm)

	return isConfirm
}
