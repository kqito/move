package command

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

func RunMove() func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		sources, err := ioutil.ReadDir(Args.OperationDir)
		if err != nil {
			return err
		}

		if len(sources) == 0 {
			fmt.Printf("%s does not exist sources.", Args.OperationDir)
		}

		var sourceNames []string

		for _, source := range sources {
			var prefix string
			prefix = "[File] "
			if source.IsDir() {
				prefix = "[Dir ] "
			}

			sourceInfo := prefix + source.Name()
			sourceNames = append(sourceNames, sourceInfo)
		}

		selectedSources := []string{}
		prompt := &survey.MultiSelect{
			Message: "Please select the target source: ",
			Options: sourceNames,
		}
		survey.AskOne(prompt, &selectedSources)

		// Execute
		mkdirAll(Args.TargetDir)

		if Flag.Copy {
			return cp(selectedSources)
		}

		return mv(selectedSources)
	}
}

func removePrefix(str string) string {
	var prefixReg = regexp.MustCompile(`^\[.*\]\s`)
	return prefixReg.ReplaceAllLiteralString(str, "")
}

func mv(sources []string) error {
	for _, source := range sources {
		oldPath := fmt.Sprintf("%s/%s", Args.OperationDir, removePrefix(source))
		newPath := fmt.Sprintf("%s/%s", Args.TargetDir, removePrefix(source))

		if err := os.Rename(oldPath, newPath); err != nil {
			return err
		}
	}

	return nil
}

func cp(sources []string) error {
	for _, source := range sources {
		sourcePath := fmt.Sprintf("%s/%s", Args.OperationDir, removePrefix(source))
		targetPath := fmt.Sprintf("%s/%s", Args.TargetDir, removePrefix(source))

		sourceInfo, err := os.Stat(sourcePath)
		if err != nil {
			return err
		}

		if sourceInfo.IsDir() {
			if err := copyDir(sourcePath, targetPath); err != nil {
				return err
			}
		} else {
			if err := copyFile(sourcePath, targetPath); err != nil {
				return err
			}
		}
	}

	return nil
}

func copyFile(sourcePath string, targetPath string) error {
	src, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	return nil
}

func copyDir(sourcePath string, targetPath string) error {
	src := filepath.Clean(sourcePath)
	dst := filepath.Clean(targetPath)

	si, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !si.IsDir() {
		return fmt.Errorf("source is not a directory")
	}

	_, err = os.Stat(dst)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	err = os.MkdirAll(dst, si.Mode())
	if err != nil {
		return err
	}

	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			err = copyDir(srcPath, dstPath)
			if err != nil {
				return err
			}
		} else {
			if entry.Mode()&os.ModeSymlink != 0 {
				continue
			}

			err = copyFile(srcPath, dstPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func mkdirAll(path string) error {
	_, statErr := os.Stat(path)
	if statErr == nil {
		return nil
	}

	err := os.MkdirAll(path, 0755)
	if err != nil {
		return err
	}

	fmt.Printf("mkdir: %s", path)

	return nil
}
