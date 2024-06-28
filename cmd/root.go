package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "atlas",
	Short: "Atlas is a CLI for managing SDK versions",
	Long:  `Atlas is a CLI tool to manage multiple versions of SDKs like JDKs.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	var installCmd = &cobra.Command{
		Use:   "install [version]",
		Short: "Install a specific SDK version",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			version := args[0]
			fmt.Printf("Installing SDK version %s\n", version)
		},
	}

	rootCmd.AddCommand(installCmd)
}
