package cmd

import (
	"atlas-cli/sdk"
	"fmt"
	"log"
	"runtime"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a specific version of JDK or Node",
	Run: func(cmd *cobra.Command, args []string) {
		dist, _ := cmd.Flags().GetString("dist")
		version, _ := cmd.Flags().GetString("version")

		if dist == "" || version == "" {
			fmt.Println("Please specify both distribution and version.")
			return
		}

		jdk, err := fetchJDKs("https://raw.githubusercontent.com/atlas-sdk/atlas-candidates/main/jdk.yml")
		if err != nil {
			fmt.Println("Error fetching JDKs:", err)
			return
		}

		var downloadURL string
		var versionJDK string
		for _, distribution := range jdk.Distributions {
			if distribution.Name == dist {
				if releases, ok := distribution.Versions[version]; ok {
					for _, release := range releases {
						if runtime.GOOS == "linux" {
							downloadURL = release.URLs.Linux
							versionJDK = release.Version
						} else if runtime.GOOS == "windows" {
							downloadURL = release.URLs.Windows
							versionJDK = release.Version
						}
						break
					}
				}
			}
		}

		if downloadURL == "" {
			fmt.Println("Could not find the specified distribution or version.")
			return
		}

		var nameZip = dist + "-" + version + ".zip"

		err = sdk.DownloadFile(nameZip, downloadURL)
		if err != nil {
			fmt.Println("Error downloading file:", err)
			return
		}

		errZip := sdk.UnzipSource(nameZip, "jdk")
		if errZip != nil {
			log.Fatal(errZip)
		}

		err = sdk.AddToPath("./jdk/" + versionJDK + "bin")
		if err != nil {
			fmt.Printf("Erro ao adicionar o diretório ao PATH: %v\n", err)
			return
		}

		fmt.Println("Download complete.")
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	installCmd.Flags().StringP("dist", "d", "", "Specify distribution to install")
	installCmd.Flags().StringP("version", "v", "", "Specify version to install")
}
