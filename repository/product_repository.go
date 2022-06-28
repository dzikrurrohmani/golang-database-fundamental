package repository

import (
	"enigmacamp.com/go-db-fundamnetal/model/entity"
	"enigmacamp.com/go-db-fundamnetal/utils"
	"errors"
	"github.com/jmoiron/sqlx"
)

type ProductRepository interface {
	Insert(product *entity.Product) error
	Update(product *entity.Product) error
	Delete(id string) error
}

type productRepository struct {
	db *sqlx.DB
}

func (p *productRepository) Insert(product *entity.Product) error {
	stmt, err := p.db.Preparex(utils.INSERT_PRODUCT)
	err = utils.IsError(err)
	if err != nil {
		return err
	}
	stmt.MustExec(product.Id, product.Name, product.Price, product.Description, product.Stock, product.Shop.Id)

	return nil
}

func (p *productRepository) Update(product *entity.Product) error {
	_, err := p.db.NamedExec(utils.UPDATE_PRODUCT, product)
	err = utils.IsError(err)
	if err != nil {
		return err
	}

	return nil
}

func (p *productRepository) Delete(id string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New("delete fail")
		}
	}()
	_, err = p.db.Exec(utils.DELETE_PRODUCT, id)
	if err != nil {
		return err
	}
	return
}

func NewProductRepository(db *sqlx.DB) ProductRepository {
	repo := new(productRepository)
	repo.db = db
	return repo
}
