package terraform

type Provider struct {
	ID  string
	URL string

	Namespace string
	Type      string
	Version   string

	OS   string
	Arch string
}
