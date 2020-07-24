package utils

import (
	"fmt"
	"os"
)

func MkdirAll(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}

	err = os.MkdirAll(path, 0755)
	if err != nil {
		return err
	}

	fmt.Printf("mkdir: %s", path)

	return nil
}
