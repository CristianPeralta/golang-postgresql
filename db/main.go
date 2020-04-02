package db

import (
	"log"
	"os"

	pg "github.com/go-pg/pg"
)

func Connect() {
	opts := &pg.Options{
		User:     "postgres",
		Password: "123456",
		Addr:     "localhost:5432",
	}
	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect to database.\n")
		os.Exit(100)
	}
	log.Printf("Connecttion to database Successful.\n")
	closeErr := db.Close()
	if closeErr != nil {
		log.Printf("Error while closing the connection, Reason: %v\n", closeErr)
		os.Exit(100)
	}
	log.Printf("Connecttion closed Successfully.\n")
	return
}
