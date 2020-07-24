package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "move [operation dir] [target dir]",
	Short: "Commands that can mv multiple files at once",
	Long: `Execute commands for multiple files from the specified directory,
			not only the mv command but also the cp command.`,
	Args: VerifyArgs(),
	RunE: RunMove(),
}

func Execute() {
	RegisterFlags(cmd)

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
