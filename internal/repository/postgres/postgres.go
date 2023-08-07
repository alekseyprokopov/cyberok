package postgres

import (
	"cyberok/internal/config"
	"database/sql"
	"fmt"
)

type Storage struct {
	config *config.DBConfig
	db     *sql.DB
}

func NewPostgresDB(config *config.DBConfig) (*sql.DB, error) {
	const op = "repository.postgres.new"

	fmt.Println(fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.Host, config.Port, config.Username, config.DBName, config.Password, config.SSLMode))

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.Host, config.Port, config.Username, config.DBName, config.Password, config.SSLMode))

	if err != nil {

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return db, nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}
