package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DataStore interface {
	CreateAccount(acc *Account) (int, error)
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

func (s *PostgresStore) CreateAccount(acc *Account) (int, error) {

	query := `
	INSERT INTO account (first_name, last_name, number, balance, created_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id;`

	rows, err := s.db.Query(query, acc.FirstName, acc.LastName, acc.Number, acc.Balance, acc.CreatedAt)

	if err != nil {
		return 0, err
	}

	for rows.Next() {

		err := rows.Scan(&acc.ID)

		if err != nil {
			log.Printf("account creation error:%+v", err.Error())
			return 0, fmt.Errorf("an error occurred will creating account")
		}
	}

	return acc.ID, nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	query := "DELETE FROM account WHERE id = $1"

	_, err := s.db.Query(query, id)

	if err != nil {
		return fmt.Errorf("account with id:%d not found", id)
	}

	return nil
}

func (s *PostgresStore) UpdateAccount(id int) error {
	return nil
}

func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {

	query := "SELECT * FROM account WHERE id = $1"

	rows, err := s.db.Query(query, id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {

		return rowToAccount(rows)

	}
	return nil, fmt.Errorf("account with id:%d not found", id)

}

func (s *PostgresStore) GetAllAccounts() ([]*Account, error) {

	query := "SELECT * FROM account"

	rows, err := s.db.Query(query)

	if err != nil {
		return nil, err
	}

	accounts := []*Account{}

	for rows.Next() {
		account, err := rowToAccount(rows)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)

	}

	return accounts, nil
}

func rowToAccount(rows *sql.Rows) (*Account, error) {
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

	return account, nil
}
