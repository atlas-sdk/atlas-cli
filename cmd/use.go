package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
    Use:   "use [version]",
    Short: "Use a specific SDK version",
    Args:  cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        version := args[0]
        fmt.Printf("Using SDK version %s\n", version)
    },
}

func init() {
    rootCmd.AddCommand(useCmd)
}
