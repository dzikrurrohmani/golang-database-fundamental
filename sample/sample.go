package sample

import (
	"enigmacamp.com/go-db-fundamnetal/model/dto"
	"enigmacamp.com/go-db-fundamnetal/model/entity"
	"enigmacamp.com/go-db-fundamnetal/repository"
	"enigmacamp.com/go-db-fundamnetal/usecase"
	"fmt"
	"github.com/jmoiron/sqlx"
	generateid "github.com/jutionck/generate-id"
	"log"
)

func CustomerRun(db *sqlx.DB) {
	cstRepo := repository.NewCustomerRepository(db)
	cstUse := usecase.NewCustomerUseCase(cstRepo)

	var customerCount []dto.CustomerCount
	customerCount, _ = cstUse.GetTotalCustomer()

	for _, row := range customerCount {
		fmt.Println(row)
	}

	var customerCountAddress []dto.CustomerAddress
	customerCountAddress, _ = cstUse.GetTotalCustomerByAddress()

	for _, row := range customerCountAddress {
		fmt.Println(row)
	}
	//var customer model.Customer
	//customerId := "cd34aac6-071e-424f-976a-c33124bdd6e4"   // Jution Kirana
	//customerId02 := "cd34aac6-071e-424f-976a-c33124bdd6ex" // Megalodon
	//customer, err := cstUse.FindCustomerById(customerId)
	//if err != nil {
	//	log.Println(err.Error())
	//}
	//fmt.Println(customer)
	//customer, err = cstUse.FindCustomerById(customerId02)
	//if err != nil {
	//	log.Println(err.Error())
	//}
	//fmt.Println(customer)
	//var customers []model.Customer

	// Get ALl
	//customers, err := cstUse.GetAll(1, 5)
	//if err != nil {
	//	log.Println(err.Error())
	//}
	//for _, customer := range customers {
	//	fmt.Println(customer)
	//}

	// Get By Name
	//customers, err := cstUse.FindCustomerByName("mendung")
	//if err != nil {
	//	log.Println(err.Error())
	//}
	//for _, customer := range customers {
	//	fmt.Println(customer)
	//}

	//// INSERT
	//cstId := generateid.GenerateId()
	//customer := model.Customer{
	//	Id:      cstId,
	//	Name:    "Jution Kirana",
	//	Address: "Ragunan",
	//	Phone:   "08292929",
	//	Email:   "jutionck@gmail.com",
	//	Balance: 150000,
	//}
	//cstUse.InsertCustomer(&customer)
	//
	//// DELETE
	//customerId := "C004"
	//err := cstUse.DeleteCustomer(customerId)
	//if err != nil {
	//	fmt.Println("error test")
	//	fmt.Println(err.Error())
	//}

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

func ProductRun(db *sqlx.DB) {
	repo := repository.NewProductRepository(db)
	useCase := usecase.NewProductUseCase(repo)

	product := entity.Product{
		Id:          generateid.GenerateId(),
		Name:        "Mouse",
		Price:       200000,
		Description: "Mouse benda",
		Stock:       20,
		Shop: entity.Shop{
			Id: "7dc92174-82ad-41a0-98d8-44aa44662c13",
		},
	}

	err := useCase.InsertNewProduct(&product)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func ShopRun(db *sqlx.DB) {
	repo := repository.NewShopRepository(db)
	useCase := usecase.NewShopUseCase(repo)

	// Get Shop with Product
	var shopWithproducts []dto.ShopProductDto
	shopWithproducts, err := useCase.FindShopWithProduct(1, 10)
	if err != nil {
		log.Println(err.Error())
	}
	for _, sp := range shopWithproducts {
		fmt.Printf("%#v\n", sp)
	}

	//shop := entity.Shop{
	//	Id:          generateid.GenerateId(),
	//	SiupNumber:  "SIPU/XII/2/2222",
	//	Name:        "Mouse",
	//	Address:     "Jakarta",
	//	PhoneNumber: "20020202",
	//}
	//
	//err := useCase.InsertNewShop(&shop)
	//if err != nil {
	//	return
	//}
}
