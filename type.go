package main

import (
	"math/rand"
	"time"
)

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstname"`
	LasttName string    `json:"lastname"`
	Number    int64     `json:"number"`
	Balance   int64     `json:"balance"`
	CreateAt  time.Time `json:"createAt"`
}

func NewAccount(firstname, lastname string) *Account {

	return &Account{
		FirstName: firstname,
		LasttName: lastname,
		Number:    int64(rand.Intn(1000000)),
		CreateAt:  time.Now().UTC(),
	}
}
