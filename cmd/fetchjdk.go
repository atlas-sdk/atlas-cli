package cmd

import (
	"atlas-cli/cmd/model"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
)

func fetchJDKs(url string) (*model.JDK, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch URL: %s, status code: %d", url, resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var jdk model.JDK
	err = yaml.Unmarshal(body, &jdk)
	if err != nil {
		return nil, err
	}

	return &jdk, nil
}
