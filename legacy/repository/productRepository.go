package repository

import "legacy/model"

type ProductRepository interface {
	FindBy(id int) *model.Product
	FindAll() []model.Product
	Save(newProduct model.Product) *model.Product
}
