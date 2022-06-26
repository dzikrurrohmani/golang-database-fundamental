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

	tabelProduct(db)
}

func tabelCustomer(db *sqlx.DB) {
	cstRepo := repository.NewCustomerRepository(db)
	cstUse := usecase.NewCustomerUseCase(cstRepo)

	// INSERT
	cstId := utils.GenerateId()
	customer := model.Customer{
		Id:      cstId,
		Name:    "Jution Kirana",
		Address: "Ragunan",
		Phone:   "08292929",
		Email:   "jutionck@gmail.com",
		Balance: 150000,
	}
	cstUse.InsertCustomer(&customer)

	// DELETE
	customerId := "C001"
	err := cstUse.DeleteCustomer(customerId)
	if err != nil {
		fmt.Println("error test")
		fmt.Println(err.Error())
	}

	// UPDATE
	//customerUpdate := model.Customer{
	//	Name:    "Jution Aja",
	//	Address: "Ragunan",
	//	Phone:   "08292929",
	//	Email:   "jutionck@gmail.com",
	//	Id:      "1c63f19d-e896-4116-a847-2dbb75d7eae8",
	//}
	//cstRepo.Update(&customerUpdate)
}

func tabelProduct(db *sqlx.DB) {
	prodRepo := repository.NewProductRepository(db)
	prodUse := usecase.NewProductUseCase(prodRepo)

	// INSERT
	// prodId := utils.GenerateId()
	// strId := utils.GenerateId()
	// product := model.Product{
	// 	Id:          "P002",
	// 	Name:        "Noxtar",
	// 	Price:       2000,
	// 	Description: "Kue coklat",
	// 	Stock:       21,
	// 	Store_id:    strId,
	// }
	// prodUse.InsertProduct(&product)

	// DELETE
	productId := "P002"
	err := prodUse.DeleteProduct(productId)
	if err != nil {
		fmt.Println("error test")
		fmt.Println(err.Error())
	}
}
