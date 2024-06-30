package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available versions of JDK and Node",
	Run: func(cmd *cobra.Command, args []string) {
		jdk, err := fetchJDKs("https://raw.githubusercontent.com/atlas-sdk/atlas-candidates/main/jdk.yml")
		if err != nil {
			fmt.Println("Error fetching JDKs:", err)
			return
		}

		for _, distribution := range jdk.Distributions {
			fmt.Printf("Distribution: %s\n", distribution.Name)
			versions := make([]string, 0, len(distribution.Versions))
			for version := range distribution.Versions {
				versions = append(versions, version)
			}
			sort.Slice(versions, func(i, j int) bool {
				return compareVersionPrefixes(versions[i], versions[j])
			})
			for _, version := range versions {
				releases := distribution.Versions[version]
				sort.Slice(releases, func(i, j int) bool {
					return compareVersions(releases[i].Version, releases[j].Version)
				})
				for _, release := range releases {
					fmt.Printf("  Version: %s (%s)\n", release.Version, version)
					//fmt.Printf("    Linux URL: %s\n", release.URLs.Linux)
					//fmt.Printf("    Windows URL: %s\n", release.URLs.Windows)
				}
			}
		}
	},
}

func compareVersionPrefixes(v1, v2 string) bool {
	order := map[string]int{
		"jdk8":  1,
		"jdk11": 2,
		"jdk17": 3,
	}
	return order[v1] < order[v2]
}

func compareVersions(v1, v2 string) bool {
	return strings.Compare(v1, v2) < 0
}

func init() {
	rootCmd.AddCommand(listCmd)
}
