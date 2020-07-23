package command

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

type ArgsType struct {
	OperationDir string
	TargetDir    string
}

var Args = &ArgsType{
	OperationDir: "",
	TargetDir:    "",
}

func VerifyArgs() cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("Require two directory args.")
		}

		if err := isExistDir(args[0], "first"); err != nil {
			return err
		}

		if err := isExistDir(args[1], "second"); err != nil {
			return err
		}

		return nil
	}
}

func isExistDir(path string, dirName string) error {
	absPath, absErr := filepath.Abs(path)
	if absErr != nil {
		return errors.New(fmt.Sprintf("The %s directory may not exist.", dirName))
	}

	fileInfo, statErr := os.Stat(absPath)
	if statErr != nil {
		return errors.New(fmt.Sprintf("It couldn't get the %s directory information.", dirName))
	}

	if !fileInfo.IsDir() {
		return errors.New(fmt.Sprintf("The %s args is not a directory", dirName))
	}

	if os.IsNotExist(statErr) {
		return errors.New(fmt.Sprintf("The directory specified in the %s argument does not exist.", dirName))
	}

	return nil
}
