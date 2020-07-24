package utils

import (
	"fmt"
)

type LogType struct {
	prefix   string
	operator string
	target   string
}

var logs []LogType

func AppendLog(operatorPath string, targetPath string, prefix string) {
	log := LogType{
		prefix:   prefix,
		operator: operatorPath,
		target:   targetPath,
	}

	logs = append(logs, log)
}

func PrintLog() {
	var maxLen = 0
	for _, log := range logs {
		if maxLen < len(log.operator) {
			maxLen = len(log.operator)
		}
	}

	// "'%-20s'\n"
	for _, log := range logs {
		fmt.Printf("[%s] %-*s  ====>  %s\n", log.prefix, maxLen, log.operator, log.target)
	}
}
