package cmd

import (
	"github.com/spf13/cobra"
	"github.com/skabashnyuk/cli/cli/command"
)

// AddCommands adds all the commands from cli/command to the root command
func AddCommands(cmd *cobra.Command, cheCli *command.CheCli) {
	cmd.AddCommand(
		NewVersionCommand(cheCli),
	)
}
