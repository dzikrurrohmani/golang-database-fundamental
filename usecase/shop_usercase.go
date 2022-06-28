package usecase

import (
	"enigmacamp.com/go-db-fundamnetal/model/dto"
	"enigmacamp.com/go-db-fundamnetal/model/entity"
	"enigmacamp.com/go-db-fundamnetal/repository"
)

type ShopUseCase interface {
	InsertNewShop(shop *entity.Shop) error
	UpdateShop(shop *entity.Shop) error
	DeleteShop(id string) error
	FindShopWithProduct(page int, totalRow int) ([]dto.ShopProductDto, error)
}

type shopUseCase struct {
	repo repository.ShopRepository
}

func (s *shopUseCase) FindShopWithProduct(page int, totalRow int) ([]dto.ShopProductDto, error) {
	return s.repo.GetShopProduct(page, totalRow)
}

func (s *shopUseCase) InsertNewShop(shop *entity.Shop) error {
	return s.repo.Insert(shop)
}

func (s *shopUseCase) UpdateShop(shop *entity.Shop) error {
	return s.repo.Update(shop)
}

func (s *shopUseCase) DeleteShop(id string) error {
	return s.repo.Delete(id)
}

func NewShopUseCase(repo repository.ShopRepository) ShopUseCase {
	usc := new(shopUseCase)
	usc.repo = repo
	return usc
}
