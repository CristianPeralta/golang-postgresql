package db

import (
	"log"

	pg "github.com/go-pg/pg"
)

type Params struct {
	Param1 string
	Param2 string
}

func PlaceHolderDemo(db *pg.DB) error {
	var value string
	var query string = "SELECT ?param1"
	params := Params{
		Param1: "This is param 1",
		Param2: "This is param 2",
	}
	_, selectErr := db.Query(pg.Scan(&value), query, params)
	if selectErr != nil {
		log.Printf("Error while running the select query, Reason: %v\n", selectErr)
		return selectErr
	}
	log.Printf("Scan successful, scanned value: %s\n", value)
	return nil
}
