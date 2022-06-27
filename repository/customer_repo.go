package repository

import (
	"errors"

	"enigmacamp.com/go-db-fundamnetal/model"
	"enigmacamp.com/go-db-fundamnetal/utils"
	"github.com/jmoiron/sqlx"
)

type CustomerRepository interface {
	Insert(customer *model.Customer) error
	Update(customer *model.Customer) error
	Delete(id string) error
	GetAll(page int, totalRow int) ([]model.Customer, error)
	GetById(id string) (model.Customer, error)
	GetByName(name string) ([]model.Customer, error)
}

type customerRepository struct {
	db *sqlx.DB
}

func (c *customerRepository) Insert(customer *model.Customer) error {
	_, err := c.db.NamedExec(utils.INSERT_CUSTOMER, customer)
	err = utils.IsError(err)
	if err != nil {
		return err
	}

	return nil
}

func (c *customerRepository) Update(customer *model.Customer) error {
	_, err := c.db.NamedExec(utils.UPDATE_CUSTOMER, customer)
	err = utils.IsError(err)
	if err != nil {
		return err
	}
	return nil
}

func (c *customerRepository) Delete(id string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New("delete fail")
		}
	}()
	c.db.MustExec(utils.DELETE_CUSTOMER, id)
	return
}

func (c *customerRepository) GetAll(page int, totalRow int) ([]model.Customer, error) {
	// rumus pagination
	limit := totalRow
	offset := limit * (page - 1)

	// // penggunaan fetch all data: select ($1), named query (:column_name), queryx
	// customers := []model.Customer{}
	// fmt.Println(customers)

	// if err := c.db.Select(&customers, utils.SELECT_CUSTOMER_ALL, limit, offset); err != nil {
	// 	utils.IsError(err)
	// 	return nil, err
	// }
	// return customers, nil
	// penggunaan fetch all data: select ($1), named query (:column_name), queryx()
	customer := []model.Customer{}
	err := c.db.Select(&customer, utils.SELECT_CUSTOMER_ALL, limit, offset)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (c *customerRepository) GetById(id string) (model.Customer, error) {
	panic("sek")
}

func (c *customerRepository) GetByName(name string) ([]model.Customer, error) {
	panic("sek")
}

func NewCustomerRepository(db *sqlx.DB) CustomerRepository {
	cstRepo := new(customerRepository)
	cstRepo.db = db
	return cstRepo
}
