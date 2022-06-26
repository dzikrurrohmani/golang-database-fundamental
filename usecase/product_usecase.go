package usecase

import (
	"enigmacamp.com/go-db-fundamnetal/model"
	"enigmacamp.com/go-db-fundamnetal/repository"
)

type ProductUseCase interface {
	InsertProduct(product *model.Product) error
	UpdateProduct(product *model.Product) error
	DeleteProduct(id string) error
}

type productUseCase struct {
	repo repository.ProductRepository
}

func (c *productUseCase) InsertProduct(product *model.Product) error {
	return c.repo.Insert(product)
}

func (c *productUseCase) UpdateProduct(product *model.Product) error {
	return c.repo.Update(product)
}

func (c *productUseCase) DeleteProduct(id string) error {
	return c.repo.Delete(id)
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	usc := new(productUseCase)
	usc.repo = repo
	return usc
}
