package terraform

type ProviderDiscovery interface {
	Discover() ([]Provider, error)
}
