package repository

import (
	"enigmacamp.com/go-db-fundamnetal/model/dto"
	"enigmacamp.com/go-db-fundamnetal/model/entity"
	"enigmacamp.com/go-db-fundamnetal/utils"
	"errors"
	"github.com/jmoiron/sqlx"
)

type ShopRepository interface {
	Insert(shop *entity.Shop) error
	Update(shop *entity.Shop) error
	Delete(id string) error
	GetShopProduct(page int, totalRow int) ([]dto.ShopProductDto, error)
}

type shopRepository struct {
	db *sqlx.DB
}

func (s *shopRepository) GetShopProduct(page int, totalRow int) ([]dto.ShopProductDto, error) {
	limit := totalRow
	offset := limit * (page - 1)
	var shopProducts []dto.ShopProductDto
	rows, err := s.db.Queryx(utils.SELECT_SHOP_WITH_PRODUCT, limit, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var sp dto.ShopProductDto
		err := rows.StructScan(&sp)
		if err != nil {
			return nil, err
		}
		shopProducts = append(shopProducts, sp)
	}
	return shopProducts, nil
}

func (s *shopRepository) Insert(shop *entity.Shop) error {
	_, err := s.db.NamedExec(utils.INSERT_SHOP, shop)
	err = utils.IsError(err)
	if err != nil {
		return err
	}
	return nil
}

func (s *shopRepository) Update(shop *entity.Shop) error {
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
	_, err = s.db.Exec(utils.DELETE_SHOP, id)
	if err != nil {
		return err
	}
	return
}

func NewShopRepository(db *sqlx.DB) ShopRepository {
	repo := new(shopRepository)
	repo.db = db
	return repo
}
