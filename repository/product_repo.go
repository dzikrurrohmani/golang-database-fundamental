package repository

import (
	"errors"

	"enigmacamp.com/go-db-fundamnetal/model"
	"enigmacamp.com/go-db-fundamnetal/utils"
	"github.com/jmoiron/sqlx"
)

type ProductRepository interface {
	Insert(product *model.Product) error
	Update(product *model.Product) error
	Delete(id string) error
}

type productRepository struct {
	db *sqlx.DB
}

func (c *productRepository) Insert(product *model.Product) error {
	_, err := c.db.NamedExec(utils.INSERT_PRODUCT, product)
	err = utils.IsError(err)
	if err != nil {
		return err
	}

	return nil
}

func (c *productRepository) Update(product *model.Product) error {
	_, err := c.db.NamedExec(utils.UPDATE_PRODUCT, product)
	err = utils.IsError(err)
	if err != nil {
		return err
	}

	return nil
}

func (c *productRepository) Delete(id string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New("delete fail")
		}
	}()
	c.db.MustExec(utils.DELETE_PRODUCT, id)
	return
}

func NewProductRepository(db *sqlx.DB) ProductRepository {
	prodRepo := new(productRepository)
	prodRepo.db = db
	return prodRepo
}
