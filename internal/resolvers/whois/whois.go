package whois

import (
	"cyberok/internal/config"
	"fmt"
	"github.com/domainr/whois"
)

type Resolver struct {
	client *whois.Client
}

func NewWhoisResolver(config *config.DNSConfig) *Resolver {
	const op = "resolvers.getWhois.new"
	duration := config.WhoisTimeout
	client := whois.NewClient(duration)
	return &Resolver{client: client}
}

func (r *Resolver) LookupWhois(host string) (string, error) {
	const op = "resolvers.getWhois.lookup"
	request, err := whois.NewRequest(host)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	response, err := r.client.Fetch(request)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	res := response.String()

	return res, nil
}
