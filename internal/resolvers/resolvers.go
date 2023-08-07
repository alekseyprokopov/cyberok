package resolvers

type FqdnResolver interface {
	LookupIp(host string) ([]string, error)
}

type WhoisResolver interface {
	LookupWhois(host string) (string, error)
}
