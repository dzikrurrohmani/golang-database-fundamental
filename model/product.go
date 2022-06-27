package model

type Product struct {
	Id          string
	Name        string
	Price       int
	Description string
	Stoct       int
	Shop        Shop `db:"store_id"`
}
