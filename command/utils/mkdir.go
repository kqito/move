package utils

import (
	"fmt"
	"os"
)

func MkdirAll(path string) error {
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
