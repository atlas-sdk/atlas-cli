package cmd

import (
	"atlas-cli/sdk"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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
		for _, distribution := range jdk.Distributions {
			if distribution.Name == dist {
				if releases, ok := distribution.Versions[version]; ok {
					for _, release := range releases {
						if runtime.GOOS == "linux" {
							downloadURL = release.URLs.Linux
						} else if runtime.GOOS == "windows" {
							downloadURL = release.URLs.Windows
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

		err = downloadFile(nameZip, downloadURL)
		if err != nil {
			fmt.Println("Error downloading file:", err)
			return
		}

		errZip := sdk.UnzipSource(nameZip, "jdk")
		if errZip != nil {
			log.Fatal(errZip)
		}

		fmt.Println("Download complete.")
	},
}

func downloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(installCmd)

	installCmd.Flags().StringP("dist", "d", "", "Specify distribution to install")
	installCmd.Flags().StringP("version", "v", "", "Specify version to install")
}
