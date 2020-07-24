package utils

import (
	"regexp"
)

func RemovePrefix(str string) string {
	var prefixReg = regexp.MustCompile(`^\[.*\]\s`)
	return prefixReg.ReplaceAllLiteralString(str, "")
}
