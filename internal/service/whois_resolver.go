package service

import (
	"cyberok/internal/resolvers"
)

type WhoisResolverService struct {
	resolver resolvers.WhoisResolver
}

func NewWhoisResolverService(resolver resolvers.WhoisResolver) *WhoisResolverService {
	return &WhoisResolverService{
		resolver: resolver,
	}
}

func (r *WhoisResolverService) LookupWhois(host string) (string, error) {
	return r.resolver.LookupWhois(host)
}
