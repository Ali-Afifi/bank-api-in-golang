package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type DataStore interface {
	CreateAccount(acc *Account) error
	DeleteAccount(id int) error
	UpdateAccount(id int) error
	GetAllAccounts() ([]*Account, error)
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

func (s *PostgresStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error {

	query := `CREATE TABLE IF NOT EXISTS account (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(255),
		last_name VARCHAR(255),
		number SERIAL,
		balance INT,
		created_at TIMESTAMP
	);`

	_, err := s.db.Exec(query)

	return err

}

func (s *PostgresStore) CreateAccount(acc *Account) error {

	query := `
	INSERT INTO account (first_name, last_name, number, balance, created_at)
	VALUES ($1, $2, $3, $4, $5);`

	rows, err := s.db.Query(query, acc.FirstName, acc.LastName, acc.Number, acc.Balance, acc.CreatedAt)

	if err != nil {
		return err
	}

	log.Printf("%+v\n", rows)

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

func (s *PostgresStore) GetAllAccounts() ([]*Account, error) {

	query := `SELECT * FROM account`

	rows, err := s.db.Query(query)

	if err != nil {
		return nil, err
	}

	accounts := []*Account{}

	for rows.Next() {
		account := &Account{}

		err := rows.Scan(
			&account.ID,
			&account.FirstName,
			&account.LastName,
			&account.Number,
			&account.Balance,
			&account.CreatedAt)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)

	}

	return accounts, nil
}
