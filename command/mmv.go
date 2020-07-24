package command

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

func RunMmv() func(cmd *cobra.Command, args []string) error {
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

		if err := mv(selectedSources); err != nil {
			return err
		}

		return nil
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

