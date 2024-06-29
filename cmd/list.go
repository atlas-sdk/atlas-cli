package cmd

import (
	"atlas-cli/sdk/service"
	"fmt"
	"github.com/spf13/cobra"
)

var osFlag string
var javaVersionFlag string

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available JDK versions",
	Run: func(cmd *cobra.Command, args []string) {
		versions, err := service.ListAvailableJDKVersions(osFlag, javaVersionFlag)
		if err != nil {
			fmt.Println("Error listing versions:", err)
			return
		}
		fmt.Println("Available JDK versions for OS:", osFlag, "and Java version:", javaVersionFlag)
		for _, version := range versions {
			fmt.Println(" -", version)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&osFlag, "os", "o", "windows", "Specify the operating system (e.g., windows, linux, mac)")
	listCmd.Flags().StringVarP(&javaVersionFlag, "java-version", "j", "", "Specify the Java version (e.g., 8, 11, 17). Leave empty to list all versions starting from Java 8")
}
