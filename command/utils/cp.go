package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ExeCp(sources []string, operationDir string, targetDir string) error {
	for _, source := range sources {
		sourcePath := fmt.Sprintf("%s/%s", operationDir, RemovePrefix(source))
		targetPath := fmt.Sprintf("%s/%s", targetDir, RemovePrefix(source))

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

