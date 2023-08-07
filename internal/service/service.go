package service

import (
	"cyberok/internal/model"
	"cyberok/internal/repository"
	"cyberok/internal/resolvers"
)

type FqdnItem interface {
	CreateFqdn(fqdn string, ip []string) (int64, error)
	UpdateFqdn(fqdn string, ips []string) (int64, error)
	GetByIPs(ips []string) ([]model.Fqdn, error)
	GetByFQDNs(fqdns []string) ([]model.Fqdn, error)
	GetAll() ([]model.Fqdn, error)
	TruncateIp() error
}

type WhoisItem interface {
	CreateWhois(domain string, whois string) (int64, error)
	UpdateWhois(domain string, whois string) error
	GetByDomains(domains []string) ([]model.Whois, error)
}

type FqdnResolver interface {
	LookupIp(host string) ([]string, error)
}

type WhoisResolver interface {
	LookupWhois(host string) (string, error)
}

type Service struct {
	FqdnItem
	WhoisItem
	FqdnResolver
	WhoisResolver
}

func New(rep *repository.Repository, FqdnResolver resolvers.FqdnResolver, WhoisResolver WhoisResolver) *Service {
	return &Service{
		FqdnItem:      NewFqdnItemService(rep.FqdnItem),
		WhoisItem:     NewWhoisItemService(rep.WhoisItem),
		FqdnResolver:  FqdnResolver,
		WhoisResolver: WhoisResolver,
	}
}
