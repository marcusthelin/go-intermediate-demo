package repository

import "pre-http/model"

//go:generate mockgen -destination=../mocks/repository/mockProductRepository.go -package=repository pre-http/repository ProductRepository

type ProductRepository interface {
	FindBy(id int) *model.Product
	FindAll() []model.Product
	Save(newProduct model.Product) *model.Product
	Update(productToBeUpdated model.Product) bool
}
