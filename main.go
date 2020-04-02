package main

import (
	"log"

	db "./db"
)

func main() {
	log.Printf("Hello world!")
	pgDb := db.Connect()
}
