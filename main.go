package main

import (
	"log"
	"time"

	db "./db"
	"github.com/go-pg/pg"
)

func main() {
	log.Printf("Hello world!")
	pgDb := db.Connect()
	SaveProduct(pgDb)
}

func SaveProduct(dbRef *pg.DB) {
	newPI := db.ProductItem{
		Name:  "Product 1",
		Desc:  "Product 1 Desc",
		Image: "this is an image path",
		Price: 3.5,
		Features: struct {
			Name string
			Desc string
			Imp  int
		}{
			Name: "Fi",
			Desc: "Desc F1",
			Imp:  3,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsActive:  true,
	}
	newPI.SaveAndReturn(dbRef)
}
