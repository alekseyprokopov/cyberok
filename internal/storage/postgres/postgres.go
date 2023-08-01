package postgres

import (
	"cyberok/internal/config"
	"database/sql"
	_ "github.com/lib/pq"
)

type Storage struct {
	config *config.DBConfig
	db     *sql.DB
}

func New(config *config.DBConfig) *Storage {
	return &Storage{
		config: config,
	}
}

func (s *Storage) Open() error {
	db, err := sql.Open("postgres", s.config.DbUrl)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *Storage) Close() {

}
