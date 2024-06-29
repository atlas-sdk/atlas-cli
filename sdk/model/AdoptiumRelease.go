package model

type AdoptiumRelease struct {
	Version struct {
		Semver string `json:"semver"`
	} `json:"version_data"`
}
