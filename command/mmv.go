package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

func RunMmv() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		fmt.Println(Flag.Copy)
	}
}
