package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install Eclipse Che",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Install called")
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
