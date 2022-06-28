package entity

type Product struct {
	Id          string
	Name        string
	Price       int
	Description string
	Stock       int
	Shop        Shop `db:"store_id"`
}
