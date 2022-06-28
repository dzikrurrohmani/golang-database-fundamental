package usecase

import (
	"enigmacamp.com/go-db-fundamnetal/model/dto"
	"enigmacamp.com/go-db-fundamnetal/model/entity"
	"enigmacamp.com/go-db-fundamnetal/repository"
)

type CustomerUseCase interface {
	RegisterNewCustomer(customer *entity.Customer) error
	UpdateCustomer(customer *entity.Customer) error
	DeleteCustomer(id string) error
	FindAllCustomer(page int, totalRow int) ([]entity.Customer, error)
	FindCustomerById(id string) (entity.Customer, error)
	FindCustomerByName(name string) ([]entity.Customer, error)
	GetTotalCustomer() ([]dto.CustomerCount, error)
	GetTotalCustomerBalanceActive() (int, error)
	GetTotalCustomerByAddress() ([]dto.CustomerAddress, error)
}

type customerUseCase struct {
	repo repository.CustomerRepository
}

func (c *customerUseCase) GetTotalCustomer() ([]dto.CustomerCount, error) {
	return c.repo.GetCount()
}

func (c *customerUseCase) GetTotalCustomerByAddress() ([]dto.CustomerAddress, error) {
	return c.repo.GetCountByAddress()
}

func (c *customerUseCase) GetTotalCustomerBalanceActive() (int, error) {
	return c.repo.GetSum()
}

func (c *customerUseCase) FindAllCustomer(page int, totalRow int) ([]entity.Customer, error) {
	return c.repo.GetAll(page, totalRow)
}

func (c *customerUseCase) FindCustomerById(id string) (entity.Customer, error) {
	return c.repo.GetById(id)
}

func (c *customerUseCase) FindCustomerByName(name string) ([]entity.Customer, error) {
	return c.repo.GetByName(name)
}

func (c *customerUseCase) RegisterNewCustomer(customer *entity.Customer) error {
	return c.repo.Insert(customer)
}

func (c *customerUseCase) UpdateCustomer(customer *entity.Customer) error {
	return c.repo.Update(customer)
}

func (c *customerUseCase) DeleteCustomer(id string) error {
	return c.repo.Delete(id)
}

func NewCustomerUseCase(repo repository.CustomerRepository) CustomerUseCase {
	usc := new(customerUseCase)
	usc.repo = repo
	return usc
}
