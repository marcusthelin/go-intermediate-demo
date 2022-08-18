package service

import (
	"go-intermediate/pre-http/model"
	"go-intermediate/pre-http/repository"
)

type ProductService struct {
	repo repository.ProductRepository
}

func (ps *ProductService) GetAllProducts() []model.Product {

	return ps.repo.FindAll()
}

func (ps *ProductService) GetProductById(id int) *model.Product {
	return ps.repo.FindBy(id)
}

func (ps *ProductService) AddProduct(name string, category model.Category, price float32) *model.Product {

	newProduct := model.Product{Name: name, Category: category, Price: price}

	return ps.repo.Save(newProduct)
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{repo}
}
