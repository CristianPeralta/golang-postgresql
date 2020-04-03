package db

import (
	"log"

	pg "github.com/go-pg/pg"
)

func PlaceHolderDemo(db *pg.DB) error {
	var value int
	var query string = "SELECT ?"
	_, selectErr := db.Query(pg.Scan(&value), query, 42)
	if selectErr != nil {
		log.Printf("Error while running the select query, Reason: %v\n", selectErr)
		return selectErr
	}
	log.Printf("Scan successful, scanned value: %v\n", value)
	return nil
}
