package dto

import "database/sql"

type ShopProduct struct {
	ShopId       string `db:"id"`
	SiupNumber   string `db:"no_siup"`
	ShopName     string `db:"name"`
	ProductId    string `db:"product_id"`
	ProductName  string `db:"product_name"`
	ProductPrice int    `db:"price"`
	ProductStock sql.NullInt16    `db:"stock"`
}
