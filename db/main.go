package db

import (
	pg "github.com/go-pg/pg"
)

func Connect() {
	opts := &pg.Options{
		User:     "postgres",
		Password: "123456",
		Addr:     "localhost:5432",
	}
	var db *pg.DB = pg.Connect(opts)
	return
}
