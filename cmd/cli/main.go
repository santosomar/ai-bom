package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var defaultCmd = "bom"

func setDefaultCommand(cmd *cobra.Command) {
	if len(os.Args) > 1 {
		pcmd := os.Args[1]
		if pcmd == "completion" || pcmd == "--version" {
			return
		}
		for _, command := range cmd.Commands() {
			if command.Use == pcmd {
				return
			}
		}
		os.Args = append([]string{os.Args[0], defaultCmd}, os.Args[1:]...)
	}
}

func Execute() {
	ctx := context.Background()

	rootCmd := NewRootCmd()

	// Defaults to `bom` command if no command is specified
	setDefaultCommand(rootCmd)

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
