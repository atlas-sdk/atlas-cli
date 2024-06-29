package service

import (
	"atlas-cli/sdk/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func fetchJDKVersions(os, javaVersion string) ([]string, error) {
	url := fmt.Sprintf("https://api.adoptium.net/v3/assets/feature_releases/%s/ga?architecture=x64&heap_size=normal&image_type=jdk&jvm_impl=hotspot&os=%s&vendor=adoptium", javaVersion, os)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %s", resp.Status)
	}

	var releases []model.AdoptiumRelease
	if err := json.NewDecoder(resp.Body).Decode(&releases); err != nil {
		return nil, err
	}

	var versions []string
	for _, release := range releases {
		versions = append(versions, release.Version.Semver)
	}

	return versions, nil
}
