package service

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
