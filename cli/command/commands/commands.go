package commands

import (
	"github.com/skabashnyuk/cli/cli/command/system"
	"github.com/spf13/cobra"
)

// AddCommands adds all the commands from cli/command to the root command
func AddCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		system.NewVersionCommand(),
	)
}
