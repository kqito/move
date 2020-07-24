package command

import (
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
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		if len(args) != 2 {
			return fmt.Errorf("Require two directory args.")
		}

		if err := isExistDir(args[0], "first"); err != nil {
			return err
		}

		operationDir, err := filepath.Rel(cwd, args[0])
		if err != nil {
			return err
		}

		targetDir, err := filepath.Rel(cwd, args[1])
		if err != nil {
			return err
		}

		if operationDir == targetDir {
			return fmt.Errorf("The args is same dir.")
		}

		Args.OperationDir = operationDir
		Args.TargetDir = targetDir

		return nil
	}
}

func isExistDir(path string, dirName string) error {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("The %s directory may not exist.", dirName)
	}

	fileInfo, err := os.Stat(absPath)
	if err != nil {
		return fmt.Errorf("The %s directory may not exist.", dirName)
	}

	if !fileInfo.IsDir() {
		return fmt.Errorf("The %s args is not a directory", dirName)
	}

	if os.IsNotExist(err) {
		return fmt.Errorf("The directory specified in the %s argument does not exist.", dirName)
	}

	return nil
}
