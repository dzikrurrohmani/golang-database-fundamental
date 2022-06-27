package main

import (
	"fmt"

	"enigmacamp.com/go-db-fundamnetal/config"
	"enigmacamp.com/go-db-fundamnetal/model"
	"enigmacamp.com/go-db-fundamnetal/repository"
	"enigmacamp.com/go-db-fundamnetal/usecase"
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

	// Transactional

	// Begin
	tx := db.MustBegin()
	// cstId1 := "10001" // Bulan
	// cstId2 := 10002   // Bintang
	// tx.MustExec(`insert into customer (id, name, saldo) values ($1, $2, $3)`, cstId1, "Bulan", 5000)
	// tx.MustExec(`insert into customer (id, name, saldo) values ($1, $2, $3)`, cstId2, "Bintang", 5000)
	cstId1 := "10001" // Bulan
	cstId2 := 10002

	rslt1 := tx.MustExec(`update customer set saldo=saldo+200 where id=$1`, cstId1)
	rslt2 := tx.MustExec(`update customer set saldo=saldo-200 where id=$1`, cstId2)
	// Rollback
	r1, _ := rslt1.RowsAffected()
	if r1 == 0 {
		tx.Rollback()
	}
	r2, _ := rslt2.RowsAffected()
	if r2 == 0 {
		tx.Rollback()
	}
	// Commit
	tx.Commit()
}

func ShopRun(db *sqlx.DB) {
	shpRepo := repository.NewShopRepository(db)
	shpUse := usecase.NewShopUseCase(shpRepo)

	shopWithProduct, err := shpUse.FindShopWithProduct(1, 4)

	utils.IsError(err)
	for _, shpPro := range shopWithProduct {
		fmt.Println(shpPro)
	}
}

func CustRun(db *sqlx.DB) {
	cstRepo := repository.NewCustomerRepository(db)
	cstUse := usecase.NewCustomerUseCase(cstRepo)

	// INSERT
	//cstId := utils.GenerateId()
	//customer := model.Customer{
	//	Id:      cstId,
	//	Name:    "Jution Kirana",
	//	Address: "Ragunan",
	//	Phone:   "08292929",
	//	Email:   "jutionck@gmail.com",
	//	Balance: 150000,
	//}
	//cstUse.InsertCustomer(&customer)

	// DELETE
	// customerId := "C001"
	// err := cstUse.DeleteCustomer(customerId)
	// if err != nil {
	// 	fmt.Println("error test")
	// 	fmt.Println(err.Error())
	// }

	// UPDATE
	//customerUpdate := model.Customer{
	//	Name:    "Jution Aja",
	//	Address: "Ragunan",
	//	Phone:   "08292929",
	//	Email:   "jutionck@gmail.com",
	//	Id:      "1c63f19d-e896-4116-a847-2dbb75d7eae8",
	//}
	//cstRepo.Update(&customerUpdate)

	// GET ALL
	var customers []model.Customer

	customers, err := cstUse.FindAllCustomer(1, 5)
	utils.IsError(err)

	for _, customer := range customers {
		fmt.Println(customer)
	}

	customers, err = cstUse.FindCustomerByName("bulan")
	utils.IsError(err)

	for _, customer := range customers {
		fmt.Println(customer)
	}

	customer, err := cstUse.FindCustomerById("C004")
	utils.IsError(err)
	fmt.Println(customer)

	customerBalance, _ := cstUse.GetTotalBalanceActiveCustomer()
	fmt.Println("Total Saldo Customer Active:", customerBalance)

	customerCount, _ := cstUse.GetTotalCustomer()
	fmt.Println("Total Customer Active:")
	for _, count := range customerCount {
		fmt.Println(count)
	}

	customerCountAddress, _ := cstUse.GetTotalCustomerByAddress()
	fmt.Println("Total Customer Address:")
	for _, count := range customerCountAddress {
		fmt.Println(count)
	}

}
