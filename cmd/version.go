package cmd

import (
	"fmt"
	"github.com/skabashnyuk/cli/cli/command"
	"github.com/spf13/cobra"
)

// NewVersionCommand creates a new cobra.Command for `che version`
func NewVersionCommand(cheCli *command.CheCli) *cobra.Command {

	return &cobra.Command{
		Use:   "version [OPTIONS]",
		Short: "Show the Eclipse Che cli version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Eclipse Che cli 1.0")
		},
	}
}
