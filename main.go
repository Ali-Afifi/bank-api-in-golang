package main

import (
	"log"
)

func main() {

	postgresDataStore, err := NewPostgresStore()

	if err != nil {
		log.Fatalf("An error occurred while trying to connect to database: %v\n", err.Error())
	}

	if err := postgresDataStore.Init(); err != nil {
		log.Fatalf("An error occurred while initializing database: %v\n", err.Error())
	}

	server := NewServer(":8080", postgresDataStore)
	server.Run()

}
