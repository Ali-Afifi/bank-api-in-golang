package main

import (
	"fmt"
	"log"
)

func main() {

	dataStore, err := NewPostgresStore()

	if err != nil {
		log.Fatalf("An error occurred while trying to connect to database: %v\n", err.Error())
	}

	fmt.Printf("%+v\n", dataStore)

	server := NewServer(":8080", dataStore)
	server.Run()

}
