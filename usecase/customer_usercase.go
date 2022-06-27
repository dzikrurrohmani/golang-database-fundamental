package usecase

import (
	"enigmacamp.com/go-db-fundamnetal/model"
	"enigmacamp.com/go-db-fundamnetal/repository"
)

type CustomerUseCase interface {
	InsertCustomer(customer *model.Customer) error
	UpdateCustomer(customer *model.Customer) error
	DeleteCustomer(id string) error
	GetAll(page int, totalRow int) ([]model.Customer,error)
	GetById(id string) (model.Customer, error)
	GetByName(name string) ([]model.Customer, error)
}

type customerUseCase struct {
	repo repository.CustomerRepository
}

func (c *customerUseCase) InsertCustomer(customer *model.Customer) error {
	return c.repo.Insert(customer)
}

func (c *customerUseCase) UpdateCustomer(customer *model.Customer) error {
	return c.repo.Update(customer)
}

func (c *customerUseCase) DeleteCustomer(id string) error {
	return c.repo.Delete(id)
}

func (c *customerUseCase) GetAll(page int, totalRow int) ([]model.Customer,error){
	return c.repo.GetAll(page, totalRow)
}

func (c *customerUseCase) GetById(id string) (model.Customer, error){
	return c.repo.GetById(id)
}

func (c *customerUseCase) GetByName(name string) ([]model.Customer, error){
	return c.repo.GetByName(name)
}

func NewCustomerUseCase(repo repository.CustomerRepository) CustomerUseCase {
	usc := new(customerUseCase)
	usc.repo = repo
	return usc
}
