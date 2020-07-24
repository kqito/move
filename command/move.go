package command

import (
	"fmt"
	"io/ioutil"

	"github.com/kqito/move/command/utils"
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

		var selectedSources []string
		utils.SelectSources(sources, selectedSources)

		// Execute
		utils.MkdirAll(Args.TargetDir)

		if Flag.Copy {
			return utils.ExeCp(selectedSources, Args.OperationDir, Args.TargetDir)
		}

		return utils.ExeMv(selectedSources, Args.OperationDir, Args.TargetDir)
	}
}
