package command

import (
	"github.com/spf13/cobra"
)

type FlagType struct {
	Copy bool
}

var Flag = &FlagType{
	Copy: false,
}

func RegisterFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&Flag.Copy, "copy", "c", Flag.Copy, "run as 'cp' command")
}
