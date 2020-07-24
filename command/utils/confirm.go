package utils

import (
	"os"

	"github.com/AlecAivazis/survey/v2"
)

func ConfirmOverride(message string, removeSourcePath string) (bool, error) {
	if !confirm(message) {
		return false, nil
	}

	removeAllErr := os.RemoveAll(removeSourcePath)
	if removeAllErr != nil {
		return false, removeAllErr
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

