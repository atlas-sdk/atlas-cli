package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available SDK versions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("List of available SDK versions")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
