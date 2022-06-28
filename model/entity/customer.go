package entity

import "database/sql"

type Customer struct {
	Id       string
	Name     string
	Address  sql.NullString
	Phone    sql.NullString
	Email    string
	Balance  int `db:"saldo"`
	IsStatus int `db:"is_status"`
}
