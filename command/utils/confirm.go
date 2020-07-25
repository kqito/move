package utils

import (
	"fmt"
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

func MkdirTargetDir(path string) (bool, error) {
	if !confirm(fmt.Sprintf("'%s' does not exist! mkdir it ?", path)) {
		return false, nil
	}

	err := os.MkdirAll(path, 0755)
	if err != nil {
		return false, err
	}

	fmt.Printf("mkdir: %s", path)

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
