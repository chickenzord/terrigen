package terraform

type Platform struct {
	OS   string `json:"os"`
	Arch string `json:"arch"`
}

type VersionInfo struct {
	Version   string     `json:"version"`
	Platforms []Platform `json:"platforms"`
	Protocols []string   `json:"-"`
}

func FindVersions(providers []Provider, providerNamespace, providerType string) (versions []VersionInfo, errr error) {
	tmpVersions := make(map[string][]Platform)

	for _, provider := range providers {
		if provider.Namespace != providerNamespace || provider.Type != providerType {
			continue
		}

		if _, ok := tmpVersions[provider.Version]; !ok {
			tmpVersions[provider.Version] = []Platform{}
		}

		tmpVersions[provider.Version] = append(tmpVersions[provider.Version], Platform{
			OS:   provider.OS,
			Arch: provider.Arch,
		})
	}

	for version, platforms := range tmpVersions {
		versions = append(versions, VersionInfo{
			Version:   version,
			Platforms: platforms,
		})
	}

	return
}
