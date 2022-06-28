package main

import (
	"database/sql"
	"log"

	"enigmacamp.com/go-db-fundamnetal/config"
	"enigmacamp.com/go-db-fundamnetal/utils"
	"github.com/jmoiron/sqlx"
)

func main() {
	config := config.NewConfigDB()
	db := config.DbConn()
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			utils.IsError(err)
		}
	}(db)

	//sample.ShopRun(db)
	//sample.CustomerRun(db)

	// Transactional

	// Begin
	defer func() {
		if r := recover(); r != nil {
			log.Println("aplikasi gagal di jalankan....")
		}
	}()
	tx := db.MustBegin()
	cstId := "10001" // Bulan
	cstId02 := 10002
	rslt := tx.MustExec(`update customer set saldo=(saldo-200) where id=$1`, cstId)
	cek(rslt, tx)
	rslt2 := tx.MustExec(`update customer set saldo=(saldo+200) where id=$1`, cstId02)
	r2, _ := rslt2.RowsAffected()
	if r2 == 0 {
		tx.Rollback()
	}
	// Commitgit
	tx.Commit()
}

func cek(rslt sql.Result, tx *sqlx.Tx){
	if r, _ := rslt.RowsAffected(); r == 0 {
		tx.Rollback()
	}
}


/**
# Fetching
-> Get All (x)
-> Get By Id (x)
-> Get By Name (x)
-> Aggregate:
   -> count
   -> sum
   -> etc...
-> Join
*/
