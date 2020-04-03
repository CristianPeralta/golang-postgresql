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
	// SaveProduct(pgDb)
	// db.PlaceHolderDemo(pgDb)
	UpdateItemPrice(pgDb)
}

func SaveProduct(dbRef *pg.DB) {
	newPI := &db.ProductItem{
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
	newPI2 := &db.ProductItem{
		Name:  "Product 2",
		Desc:  "Product 2 Desc",
		Image: "this is an image path",
		Price: 3.5,
		Features: struct {
			Name string
			Desc string
			Imp  int
		}{
			Name: "F2",
			Desc: "Desc F2",
			Imp:  3,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsActive:  true,
	}
	totalItems := []*db.ProductItem{newPI, newPI2}
	// newPI.Save(dbRef)
	// newPI.SaveAndReturn(dbRef)
	newPI.SaveMultiple(dbRef, totalItems)
}

func DeleteItem(dbRef *pg.DB) {
	newPi := &db.ProductItem{
		Name: "Product 5",
	}
	newPi.DeleteItem(dbRef)
}

func UpdateItemPrice(dbRef *pg.DB) {
	newPI := &db.ProductItem{
		ID:    1,
		Price: 5.0,
	}
	newPI.UpdatePrice(dbRef)
}
