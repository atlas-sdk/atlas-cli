package model

type JDK struct {
	Distributions []struct {
		Name     string `yaml:"name"`
		Versions map[string][]struct {
			Version string `yaml:"version"`
			URLs    struct {
				Linux   string `yaml:"linux"`
				Windows string `yaml:"windows"`
			} `yaml:"urls"`
		} `yaml:"versions"`
	} `yaml:"distributions"`
}
