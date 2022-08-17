package service

import (
	"intermediate-go/legacy/model"
	"intermediate-go/legacy/repository"
)

func GetAllProducts() []model.Product {
	repo := repository.NewInMemoryProductRepo()
	return repo.FindAll()
}

func GetProductById(id int) *model.Product {
	repo := repository.NewInMemoryProductRepo()
	return repo.FindBy(id)
}

func AddProduct(name string, productType model.Category, price float32) *model.Product {
	repo := repository.NewInMemoryProductRepo()
	newProduct := model.Product{Name: name, ProductType: productType, Price: price}

	return repo.Save(newProduct)
}
