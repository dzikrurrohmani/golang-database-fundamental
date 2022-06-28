package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "dzikrurrohmani"
	dbPassword := "password"
	dbName := "golang_db_fundamental"

	// urutan url koneksi ke db postgres
	// postgres://dbUser:dbPassword@dbHost:dbPort/dbName
	// dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		panic(err.Error())
	} else {
		log.Println("Connected")
	}

	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			panic(err.Error())
		}
	}(db)
	
}