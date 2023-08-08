package postgres

import (
	"cyberok/internal/model"
	"database/sql"
	"fmt"
	"github.com/jackskj/carta"
	"github.com/lib/pq"
)

type WhoisItemPostgres struct {
	db *sql.DB
}

func NewWhoisItemPostgres(db *sql.DB) *WhoisItemPostgres {
	return &WhoisItemPostgres{db: db}
}

func (s *WhoisItemPostgres) CreateWhois(domain string, whois string) (int64, error) {
	const op = "repository.postgres.whoisItem.SaveWhois"
	var id int64
	err := s.db.QueryRow("INSERT INTO whois(domain, info) VALUES($1,$2) RETURNING id", domain, whois).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *WhoisItemPostgres) UpdateWhois(domain string, whois string) error {
	const op = "repository.postgres.whoisItem.SaveWhois"
	err := s.db.QueryRow("UPDATE whois SET info=$1 WHERE domain=$2", whois, domain)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *WhoisItemPostgres) GetByDomains(domains []string) ([]model.Whois, error) {
	const op = "repository.postgres.whoisItem.GetWhoisByDomains"

	q := "SELECT domain,info FROM whois WHERE domain=ANY($1)"

	rows, err := s.db.Query(q, pq.Array(domains))
	if err != nil {
		return nil, fmt.Errorf("%s:query exec %w", op, err)
	}

	var res []model.Whois

	if err = carta.Map(rows, &res); err != nil {
		return nil, fmt.Errorf("%s:carta's scan %w", op, err)
	}

	return res, nil
}
