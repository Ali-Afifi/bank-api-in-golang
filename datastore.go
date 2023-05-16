package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type DataStore interface {
	CreateAccount(acc *Account) error
	DeleteAccount(id int) error
	UpdateAccount(id int) error
	GetAccountByID(id int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	// connStr := "host=db port=5432 user=postgres password=example dbname=postgres sslmode=disable"
	connStr := "host=localhost port=5432 user=postgres password=example dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{db: db}, nil

}

func (s *PostgresStore) CreateAccount(acc *Account) error {
	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	return nil
}

func (s *PostgresStore) UpdateAccount(id int) error {
	return nil
}

func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}
