package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var lsCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed Eclipse Che versions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
