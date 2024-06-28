package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install [version]",
	Short: "Install a specific SDK version",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		version := args[0]
		fmt.Printf("Installing SDK version %s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
