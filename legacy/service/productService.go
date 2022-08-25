package service

import (
	"legacy/model"
	"legacy/repository"
)

//var (
//  //repo repository.ProductRepository = repository.NewInMemoryProductRepo()
//  repo repository.ProductRepository = repository.NewDbProductRepo()
//)

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return ProductService{
		repo: productRepo,
	}
}

func (p ProductService) GetAllProducts() []model.Product {
	//repo := repository.NewInMemoryProductRepo()
	return p.repo.FindAll()
	//return repo.GetAllProducts()
}

func (p ProductService) GetProductById(id int) *model.Product {
	//repo := repository.NewInMemoryProductRepo()
	return p.repo.FindBy(id)
	//return repo.GetProductById(id)
}

func (p ProductService) AddProduct(name string, category model.Category, price float32) *model.Product {
	//repo := repository.NewInMemoryProductRepo()
	newProduct := model.Product{Name: name, Category: category, Price: price}

	return p.repo.Save(newProduct)
	//return repo.AddProduct(newProduct)
}
