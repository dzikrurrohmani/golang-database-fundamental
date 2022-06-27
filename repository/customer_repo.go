package repository

import (
	"errors"
	"enigmacamp.com/go-db-fundamnetal/model"
	"enigmacamp.com/go-db-fundamnetal/model/dto"
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
	GetSum() (int, error)
	GetCount() ([]dto.CustomerCount, error)
	GetCountByAddress() ([]dto.CustomerAddress, error)
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

	// penggunaan fetch all data: select ($1), named query (:column_name), queryx
	customers := []model.Customer{}
	if err := c.db.Select(&customers, utils.SELECT_CUSTOMER_ALL, limit, offset); err != nil {
		utils.IsError(err)
		return nil, err
	}
	return customers, nil
}

func (c *customerRepository) GetById(id string) (model.Customer, error) {
	customer := model.Customer{}
	if err := c.db.Get(&customer, utils.SELECT_CUSTOMER_BY_ID, id); err != nil {
		utils.IsError(err)
		return model.Customer{}, err
	}
	return customer, nil
}

func (c *customerRepository) GetByName(name string) ([]model.Customer, error) {
	customers := []model.Customer{}
	if err := c.db.Select(&customers, utils.SELECT_CUSTOMER_BY_NAME, "%"+name+"%"); err != nil {
		utils.IsError(err)
		return nil, err
	}
	return customers, nil
}

func (c *customerRepository) GetSum() (int, error) {
	var customerBalance int
	if err := c.db.Get(&customerBalance, utils.SELECT_SUM_SALDO_CUSTOMER); err != nil {
		utils.IsError(err)
		return 0, err
	}
	return customerBalance, nil
}
func (c *customerRepository) GetCount() ([]dto.CustomerCount, error) {
	var customerCount []dto.CustomerCount
	if err := c.db.Select(&customerCount, utils.SELECT_COUNT_CUSTOMER); err != nil {
		utils.IsError(err)
		return nil, err
	}
	return customerCount, nil
}

func (c *customerRepository) GetCountByAddress() ([]dto.CustomerAddress, error) {
	var customerCount []dto.CustomerAddress
	if err := c.db.Select(&customerCount, utils.SELECT_COUNT_CUSTOMER_ADDRESS); err != nil {
		utils.IsError(err)
		return nil, err
	}
	return customerCount, nil
}

func NewCustomerRepository(db *sqlx.DB) CustomerRepository {
	cstRepo := new(customerRepository)
	cstRepo.db = db
	return cstRepo
}
