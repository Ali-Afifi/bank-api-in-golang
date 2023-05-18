package main

import (
	"math/rand"
	"time"
)

type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Number    int64     `json:"number"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateAccountRequestBody struct {
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
}


func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        0,
		FirstName: firstName,
		LastName:  lastName,
		Number:    rand.Int63n(1000000000),
		CreatedAt: time.Now().UTC(),
	}
}

func (a *Account) SetID(id int) {
	a.ID = id
}
