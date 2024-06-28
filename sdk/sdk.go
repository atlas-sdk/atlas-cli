package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AdoptiumRelease struct {
	Version struct {
		Semver string `json:"semver"`
	} `json:"version_data"`
}

func ListAvailableJDKVersions(os, javaVersion string) ([]string, error) {
	var versions []string
	if javaVersion != "" {
		vs, err := fetchJDKVersions(os, javaVersion)
		if err != nil {
			return nil, err
		}
		versions = append(versions, vs...)
	} else {
		javaVersions := []string{"8", "11", "17"}
		for _, v := range javaVersions {
			vs, err := fetchJDKVersions(os, v)
			if err != nil {
				return nil, err
			}
			versions = append(versions, vs...)
		}
	}
	return versions, nil
}

func fetchJDKVersions(os, javaVersion string) ([]string, error) {
	url := fmt.Sprintf("https://api.adoptium.net/v3/assets/feature_releases/%s/ga?architecture=x64&heap_size=normal&image_type=jdk&jvm_impl=hotspot&os=%s&vendor=adoptium", javaVersion, os)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %s", resp.Status)
	}

	var releases []AdoptiumRelease
	if err := json.NewDecoder(resp.Body).Decode(&releases); err != nil {
		return nil, err
	}

	var versions []string
	for _, release := range releases {
		versions = append(versions, release.Version.Semver)
	}

	return versions, nil
}
