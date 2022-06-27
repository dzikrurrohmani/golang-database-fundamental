package model

import "database/sql"

type Customer struct {
	Id       string
	Name     string
	Address  sql.NullString
	Phone    string
	Email    sql.NullString
	Balance  int `db:"saldo"`
	IsStatus int `db:"is_status"`
}
