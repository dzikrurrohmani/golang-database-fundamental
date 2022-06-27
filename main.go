package main

import (
	"fmt"
	"time"

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
	cstRepo := repository.NewCustomerRepository(db)
	cstUse := usecase.NewCustomerUseCase(cstRepo)
	time.Sleep(time.Second)

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

	customers, err := cstUse.FindAllCustomer(1, 2)
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
}
