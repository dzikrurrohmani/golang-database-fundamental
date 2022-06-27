package usecase

import (
	"enigmacamp.com/go-db-fundamnetal/model"
	"enigmacamp.com/go-db-fundamnetal/model/dto"
	"enigmacamp.com/go-db-fundamnetal/repository"
)

type ShopUseCase interface {
	InsertNewShop(shop *model.Shop) error
	UpdateShop(shop *model.Shop) error
	DeleteCshop(id string) error
	FindShopWithProduct(page int, totalRow int) ([]dto.ShopProduct, error)
}

type shopUseCase struct {
	repo repository.ShopRepository
}

func (s *shopUseCase) InsertNewShop(shop *model.Shop) error {
	return s.repo.Insert(shop)
}

func (s *shopUseCase) UpdateShop(shop *model.Shop) error {
	return s.repo.Update(shop)
}

func (s *shopUseCase) DeleteCshop(id string) error {
	return s.repo.Delete(id)
}

func (s *shopUseCase) FindShopWithProduct(page int, totalRow int) ([]dto.ShopProduct, error) {
	return s.repo.GetShopProduct(page, totalRow)
}

func NewShopUseCase(repo repository.ShopRepository) ShopUseCase {
	usc := new(shopUseCase)
	usc.repo = repo
	return usc
}
