package utils

import (
	"os"

	"github.com/AlecAivazis/survey/v2"
)

func SelectSources(sources []os.FileInfo, selected *[]string) {
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

	prompt := &survey.MultiSelect{
		Message: "Please select the target source: ",
		Options: sourceNames,
	}

	survey.AskOne(prompt, selected)
}
