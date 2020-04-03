package db

import (
	"log"
	"os"
	"time"

	pg "github.com/go-pg/pg"
)

func Connect() *pg.DB {
	opts := &pg.Options{
		User:         "postgres",
		Password:     "123456",
		Addr:         "localhost:5432",
		Database:     "mydb",
		DialTimeout:  30 * time.Second,
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
		IdleTimeout:  30 * time.Minute,
		MaxConnAge:   1 * time.Minute,
		PoolSize:     20,
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
