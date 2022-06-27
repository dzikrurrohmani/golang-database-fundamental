package repository

import (
	"errors"

	"enigmacamp.com/go-db-fundamnetal/model"
	"enigmacamp.com/go-db-fundamnetal/model/dto"
	"enigmacamp.com/go-db-fundamnetal/utils"
	"github.com/jmoiron/sqlx"
)

type ShopRepository interface {
	Insert(shop *model.Shop) error
	Update(shop *model.Shop) error
	Delete(id string) error
	GetShopProduct(page int, totalRow int) ([]dto.ShopProduct, error)
}

type shopRepository struct {
	db *sqlx.DB
}

func (s *shopRepository) Insert(shop *model.Shop) error {
	_, err := s.db.NamedExec(utils.INSERT_SHOP, shop)
	err = utils.IsError(err)
	if err != nil {
		return err
	}

	return nil
}

func (s *shopRepository) Update(shop *model.Shop) error {
	_, err := s.db.NamedExec(utils.UPDATE_SHOP, shop)
	err = utils.IsError(err)
	if err != nil {
		return err
	}
	return nil
}

func (s *shopRepository) Delete(id string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New("delete fail")
		}
	}()
	s.db.MustExec(utils.DELETE_SHOP, id)
	return
}

func (s *shopRepository) GetShopProduct(page int, totalRow int) ([]dto.ShopProduct, error) {
	// rumus pagination
	limit := totalRow
	offset := limit * (page - 1)

	var shopProduct []dto.ShopProduct
	rows, err := s.db.Queryx((utils.SELECT_SHOP_WITH_PRODUCT), limit, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var sp dto.ShopProduct
		err := rows.StructScan(&sp)
		if err != nil {
			return nil, err
		}
		shopProduct = append(shopProduct, sp)
	}
	return shopProduct, nil
}

func NewShopRepository(db *sqlx.DB) ShopRepository {
	shpRepo := new(shopRepository)
	shpRepo.db = db
	return shpRepo
}

