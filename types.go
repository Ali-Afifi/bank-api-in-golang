package main

import (
	"math/rand"
	"time"
)

type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Number    int       `json:"number"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateAccountRequestBody struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type TransferRequestBody struct {
	ToAccount int `json:"toAccount"`
	Amount    int `json:"amount"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        0,
		FirstName: firstName,
		LastName:  lastName,
		Number:    rand.Intn(1000000000),
		CreatedAt: time.Now().UTC(),
	}
}

func (a *Account) SetID(id int) {
	a.ID = id
}
