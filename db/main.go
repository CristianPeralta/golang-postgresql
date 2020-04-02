package db

import (
	"log"
	"os"

	pg "github.com/go-pg/pg"
)

func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "postgres",
		Password: "123456",
		Addr:     "localhost:5432",
		Database: "mydb",
	}
	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect to database.\n")
		os.Exit(100)
	}
	log.Printf("Connecttion to database Successful.\n")
	CreateProdItemsTable(db)
	return db
}
