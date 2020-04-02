package db

import (
	"log"
	"time"

	pg "github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type ProductItem struct {
	RefPointer int      `sql:"-"`
	tableName  struct{} `sql:"product_items_collection"`
	ID         int      `sql:"id, pk"`
	Name       string   `sql:"name:unique"`
	Desc       string   `sql:"desc"`
	Image      string   `sql:"image"`
	Price      float64  `sql:"price, type:real"`
	Features   struct {
		Name string
		Desc string
		Imp  int
	} `sql:"features, type:jsonb"`
	CreatedAt time.Time `sql:"created_at"`
	UpdatedAt time.Time `sql:"updated_at"`
	IsActive  bool      `sql:"is_active"`
}

func CreateProdItemsTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&ProductItem{}, opts)
	if createErr != nil {
		log.Printf("Error while creating table ProductItem, Reason %v\n", createErr)
		return createErr
	}
	log.Printf("Table ProductItem created successfully.\n")
	return nil
}
