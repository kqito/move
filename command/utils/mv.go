package utils

import (
	"fmt"
	"os"
)

func ExeMv(sources []string, operationDir string, targetDir string) error {
	for _, source := range sources {
		oldPath := fmt.Sprintf("%s/%s", operationDir, RemovePrefix(source))
		newPath := fmt.Sprintf("%s/%s", targetDir, RemovePrefix(source))

		_, statErr := os.Stat(newPath)
		if !os.IsNotExist(statErr) {
			isConfirmed, removeErr := OverwriteSource(fmt.Sprintf("%s is exist! Overwrite?", newPath), newPath)
			if removeErr != nil {
				return removeErr
			}

			if !isConfirmed {
				continue
			}
		}

		if err := os.Rename(oldPath, newPath); err != nil {
			return err
		}
	}

	return nil
}

