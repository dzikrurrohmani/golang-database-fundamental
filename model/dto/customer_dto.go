package dto

import "database/sql"

type CustomerCount struct {
	IsStatus int `db:"is_status"`
	Total    int
}

type CustomerAddress struct {
	Address sql.NullString
	Total   int
}
