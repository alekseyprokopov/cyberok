package repository

import (
	"cyberok/internal/model"
	"cyberok/internal/repository/postgres"
	"database/sql"
	_ "github.com/lib/pq"
)

type FqdnItem interface {
	CreateFqdn(fqdn string, ips []string) (int64, error)
	UpdateFqdn(fqdn string, ips []string) (int64, error)
	GetByIPs(ips []string) ([]model.Fqdn, error)
	GetByFQDNs(fqdns []string) ([]model.Fqdn, error)
	GetAll() ([]model.Fqdn, error)
	TruncateIp() error
}

type WhoisItem interface {
	CreateWhois(domain string, whois string) (int64, error)
	GetByDomains(domains []string) ([]model.Whois, error)
	UpdateWhois(domain string, whois string) error
}

type Repository struct {
	FqdnItem
	WhoisItem
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		FqdnItem:  postgres.NewFqdnItemPostgres(db),
		WhoisItem: postgres.NewWhoisItemPostgres(db),
	}
}
