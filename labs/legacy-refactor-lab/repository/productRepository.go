package repository

import "legacy-refactor-lab/model"

type ProductRepository interface {
	FindBy(id int) *model.Product
	FindAll() []model.Product
	Save(newProduct model.Product) *model.Product
}
