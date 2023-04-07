package main

import (
	"flag"
	"fmt"
	"log"
)

func seedAccount(store Storage, fname, lname, pw string) *Account {
	acc, err := NewAccount(fname, lname, pw)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}

	fmt.Println("New account => ", acc.Number)

	return acc
}

func seedAccounts(s Storage) {
	seedAccount(s, "Fábio", "Valinhos", "123456")
}

func main() {

	seed := flag.Bool("seed", false, "seed the bd")
	flag.Parse()

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.init(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		fmt.Println("seeding the database")

		// seed stuff
		seedAccounts(store)
	}

	server := NewAPIServer(":3000", store)
	server.Run()
}
