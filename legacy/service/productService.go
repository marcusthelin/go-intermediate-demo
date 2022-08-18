package service

import (
	"go-intermediate/legacy/model"
	"go-intermediate/legacy/repository"
)

func GetAllProducts() []model.Product {
	repo := repository.NewInMemoryProductRepo()
	return repo.FindAll()
}

func GetProductById(id int) *model.Product {
	repo := repository.NewInMemoryProductRepo()
	return repo.FindBy(id)
}

func AddProduct(name string, category model.Category, price float32) *model.Product {
	repo := repository.NewInMemoryProductRepo()
	newProduct := model.Product{Name: name, Category: category, Price: price}

	return repo.Save(newProduct)
}
