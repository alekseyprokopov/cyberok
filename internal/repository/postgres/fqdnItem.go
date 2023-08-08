package postgres

import (
	"cyberok/internal/model"
	"database/sql"
	"fmt"
	"github.com/jackskj/carta"
	"github.com/lib/pq"
)

type FqdnItemPostgres struct {
	db *sql.DB
}

func NewFqdnItemPostgres(db *sql.DB) *FqdnItemPostgres {
	return &FqdnItemPostgres{db: db}
}

func (s *FqdnItemPostgres) CreateFqdn(fqdn string, ips []string) (int64, error) {
	const op = "repository.postgres.fqdnItem.Create"
	var fqdnId int64
	stmtFqdn, err := s.db.Prepare("INSERT INTO fqdn(name) VALUES($1) RETURNING id")
	if err != nil {
		return 0, fmt.Errorf("%s: prepare statement: %w", op, err)
	}
	err = stmtFqdn.QueryRow(fqdn).Scan(&fqdnId)
	if err != nil {
		return 0, fmt.Errorf("%s:scan %w", op, err)
	}

	stmtIp, err := s.db.Prepare("INSERT INTO ip(fqdn_id, ip) VALUES($1,$2)")
	for _, ip := range ips {
		_, err = stmtIp.Exec(fqdnId, ip)
		if err != nil {
			return 0, fmt.Errorf("%s:insert %w", op, err)
		}
	}

	return fqdnId, nil
}

func (s *FqdnItemPostgres) UpdateFqdn(fqdn string, ips []string) (int64, error) {
	const op = "repository.postgres.fqdnItem.Update"

	stmtFqdn, err := s.db.Prepare("SELECT id FROM fqdn WHERE name=$1")
	if err != nil {
		return 0, fmt.Errorf("%s: prepare statement: %w", op, err)
	}

	var fqdnId int64

	err = stmtFqdn.QueryRow(fqdn).Scan(&fqdnId)
	if err != nil {
		return 0, fmt.Errorf("%s:query exec %w", op, err)
	}

	stmtDelete, err := s.db.Prepare("DELETE FROM ip WHERE fqdn_id=$1;")
	if err != nil {
		return 0, fmt.Errorf("%s: prepare statement: %w", op, err)
	}
	_, err = stmtDelete.Exec(fqdnId)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := s.db.Prepare("INSERT INTO ip(fqdn_id, ip) VALUES($1,$2)")

	for _, ip := range ips {
		_, err = stmt.Exec(fqdnId, ip)
		if err != nil {
			return 0, fmt.Errorf("%s: %w", op, err)
		}
	}

	return fqdnId, nil
}

func (s *FqdnItemPostgres) GetByIPs(ips []string) ([]model.Fqdn, error) {
	const op = "repository.postgres.GetByIPs"

	q := `
			SELECT name AS fqdn_name, ip 
			FROM fqdn 
			LEFT JOIN ip ON ip.fqdn_id = fqdn.id 
			WHERE ip=ANY($1)
			`

	rows, err := s.db.Query(q, pq.Array(ips))
	if err != nil {
		return nil, fmt.Errorf("%s:query exec %w", op, err)
	}

	var result []model.Fqdn

	if err = carta.Map(rows, &result); err != nil {
		return nil, fmt.Errorf("%s:carta's scan %w", op, err)
	}

	return result, nil
}

func (s *FqdnItemPostgres) GetByFQDNs(fqdns []string) ([]model.Fqdn, error) {
	const op = "repository.postgres.GetByFQDNs"

	q := `
			SELECT name AS fqdn_name, ip 
			FROM fqdn 
			LEFT JOIN ip ON ip.fqdn_id = fqdn.id 
			WHERE name = ANY($1)
			`

	rows, err := s.db.Query(q, pq.Array(fqdns))
	if err != nil {
		return nil, fmt.Errorf("%s:query exec %w", op, err)
	}

	var result []model.Fqdn
	if err = carta.Map(rows, &result); err != nil {
		return nil, fmt.Errorf("%s:carta's scan %w", op, err)
	}

	return result, nil
}

func (s *FqdnItemPostgres) GetAll() ([]model.Fqdn, error) {
	const op = "repository.postgres.GetAll"

	q := `
			SELECT name AS fqdn_name, ip 
			FROM fqdn 
			LEFT JOIN ip ON ip.fqdn_id = fqdn.id 
			`

	rows, err := s.db.Query(q)
	if err != nil {
		return nil, fmt.Errorf("%s:query exec %w", op, err)
	}

	var result []model.Fqdn
	if err = carta.Map(rows, &result); err != nil {
		return nil, fmt.Errorf("%s:carta's scan %w", op, err)
	}

	return result, nil
}

func (s *FqdnItemPostgres) TruncateIp() error {
	const op = "repository.postgres.TruncateIp"
	q := `TRUNCATE ip`

	_, err := s.db.Exec(q)
	if err != nil {
		return fmt.Errorf("%s:query exec %w", op, err)
	}

	return nil
}
