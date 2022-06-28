package entity

import "time"

type Transaction struct {
	Id           string
	Customer     Customer `db:"customer_id"`
	Product      Product  `db:"product_id"`
	Shop         Shop     `db:"store_id"`
	PurchaseDate time.Time
	Quantity     int
}
