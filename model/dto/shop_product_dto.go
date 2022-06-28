package dto

type ShopProductDto struct {
	ShopId       string `db:"id"`
	ShopNumber   string `db:"no_siup"`
	ShopName     string `db:"name"`
	ProductId    string `db:"product_id"`
	ProductName  string `db:"product_name"`
	ProductPrice int    `db:"price"`
	ProductStock int    `db:"stock"`
}
