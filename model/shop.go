package model

type Shop struct {
	Id          string
	ShopNumber  string `db:"no_siup"`
	Name        string
	Address     string
	PhoneNumber string `db:"phone"`
}
