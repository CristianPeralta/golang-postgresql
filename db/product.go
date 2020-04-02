package db

import "time"

type ProductItem struct {
	tableName struct{}  `sql:"product_items_collection"`
	ID        int       `sql:"id, pk"`
	Name      string    `sql:"name:unique"`
	Desc      string    `sql:"desc"`
	Image     string    `sql:"image"`
	Price     float64   `sql:"price, type:real"`
	Features  string    `sql:"features"`
	CreatedAt time.Time `sql:"created_at"`
	UpdatedAt time.Time `sql:"updated_at"`
	IsActive  bool      `sql:"is_active"`
}
