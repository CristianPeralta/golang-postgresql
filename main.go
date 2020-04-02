package main

import (
	"log"

	db "./db"
)

func main() {
	log.Printf("Hello world!")
	db.Connect()
}
