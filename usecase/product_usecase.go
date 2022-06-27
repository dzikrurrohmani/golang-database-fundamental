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

func (p *productUseCase) InsertProduct(product *model.Product) error {
	return p.repo.Insert(product)
}

func (p *productUseCase) UpdateProduct(product *model.Product) error {
	return p.repo.Update(product)
}

func (p *productUseCase) DeleteProduct(id string) error {
	return p.repo.Delete(id)
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	usc := new(productUseCase)
	usc.repo = repo
	return usc
}
