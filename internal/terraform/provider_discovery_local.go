package terraform

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type StringFilterFunc func(string) bool

var StringFilterMatchAll = func(string) bool {
	return true
}

type LocalProviderDiscoveryOpts struct {
	PluginsPath    string
	RegistryFilter StringFilterFunc
}

type localProviderDiscovery struct {
	pluginsPath string
}

func DefaultPluginsPath() string {
	return fmt.Sprintf("%s/.terraform.d/plugins", os.Getenv("HOME"))
}

func NewLocalProviderDiscovery(opts LocalProviderDiscoveryOpts) ProviderDiscovery {
	return &localProviderDiscovery{
		pluginsPath: opts.PluginsPath,
	}
}

var _ ProviderDiscovery = (*localProviderDiscovery)(nil)

func (pd *localProviderDiscovery) Discover() (providers []Provider, err error) {
	err = filepath.WalkDir(pd.pluginsPath, func(path string, d fs.DirEntry, err error) (e error) {
		if path == pd.pluginsPath {
			return
		}

		if !strings.HasPrefix(path, pd.pluginsPath+"/") {
			return
		}

		id := strings.TrimPrefix(path, pd.pluginsPath+"/")
		segments := strings.Split(id, "/")
		if len(segments) != 6 {
			return
		}

		osArch := strings.Split(segments[4], "_")
		os := osArch[0]
		arch := osArch[1]

		providers = append(providers, Provider{
			ID:        id,
			URL:       "file://" + path,
			Namespace: segments[1],
			Type:      segments[2],
			Version:   segments[3],
			OS:        os,
			Arch:      arch,
		})

		return
	})

	return
}
