package model

type Customer struct {
	Id       string
	Name     string
	Address  string
	Phone    string
	Email    string
	Balance  int `db:"saldo"`
	IsStatus int `db:"is_status"`
}
