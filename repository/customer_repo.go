package repository

import (
	"enigmacamp.com/go-db-fundamnetal/model"
	"enigmacamp.com/go-db-fundamnetal/utils"
	"errors"
	"github.com/jmoiron/sqlx"
)

type CustomerRepository interface {
	Insert(customer *model.Customer) error
	Update(customer *model.Customer) error
	Delete(id string) error
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

func NewCustomerRepository(db *sqlx.DB) CustomerRepository {
	cstRepo := new(customerRepository)
	cstRepo.db = db
	return cstRepo
}
