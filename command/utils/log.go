package utils

import (
	"fmt"
)

var Log []string

func AppendLog(operatorPath string, targetPath string, prefix string) {
	log := fmt.Sprintf("[%s] %s ===> %s", prefix, operatorPath, targetPath)

	Log = append(Log, log)
}

func PrintLog() {
	fmt.Println("")
	for _, log := range Log {
		fmt.Println(log)
	}
}
