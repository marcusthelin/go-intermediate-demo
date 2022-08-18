package repository

import "pre-http-lab/model"

type ProductRepository interface {
	FindBy(id int) *model.Product
	FindAll() []model.Product
	Save(newProduct model.Product) *model.Product
	Update(productToBeUpdated model.Product) bool
}
