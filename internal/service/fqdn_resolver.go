package service

import (
	"cyberok/internal/resolvers"
)

type FqdnResolverService struct {
	resolver resolvers.FqdnResolver
}

func NewFqdnResolverService(resolver resolvers.FqdnResolver) *FqdnResolverService {
	return &FqdnResolverService{
		resolver: resolver,
	}
}

func (s *FqdnResolverService) LookupIp(host string) ([]string, error) {
	return s.resolver.LookupIp(host)
}
