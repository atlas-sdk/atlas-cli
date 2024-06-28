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
	// Aqui você pode adicionar subcomandos como list, install, use, etc.
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List available SDK versions",
		Run: func(cmd *cobra.Command, args []string) {
			// Lógica para listar versões disponíveis
			fmt.Println("List of available SDK versions")
		},
	}

	var installCmd = &cobra.Command{
		Use:   "install [version]",
		Short: "Install a specific SDK version",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			version := args[0]
			// Lógica para instalar a versão especificada
			fmt.Printf("Installing SDK version %s\n", version)
		},
	}

	var useCmd = &cobra.Command{
		Use:   "use [version]",
		Short: "Use a specific SDK version",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			version := args[0]
			// Lógica para configurar a versão especificada como padrão
			fmt.Printf("Using SDK version %s\n", version)
		},
	}

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(useCmd)
}
