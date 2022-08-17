package repository

import (
	"intermediate-go/legacy/model"
)

type DbProductRepo struct {
	products map[int]model.Product
}

func (r DbProductRepo) GetProductById(id int) *model.Product {
	if p, ok := r.products[id]; ok {
		return &p
	}
	return nil
}

func (r DbProductRepo) GetAllProducts() []model.Product {
	v := make([]model.Product, 0)
	for _, p := range r.products {
		v = append(v, p)
	}
	return v
}

func (r DbProductRepo) AddProduct(newProduct model.Product) *model.Product {
	id := len(r.products) + 1
	newProduct.Id = id
	r.products[id] = newProduct
	return &newProduct
}

func (r DbProductRepo) UpdateProduct(productToBeUpdated model.Product) bool {
	if product := r.GetProductById(productToBeUpdated.Id); product != nil {
		r.products[productToBeUpdated.Id] = productToBeUpdated
		return true
	}
	return false
}

func NewDbProductRepo() DbProductRepo {
	db := DbProductRepo{
		products: map[int]model.Product{
			1: model.NewProduct(1, "Pulses", "GROCESSORY", 100.0),
			2: model.NewProduct(2, "Learning Java Streams", "BOOKS", 55.0),
			3: model.NewProduct(3, "Night light", "BABY", 41.46),
			4: model.NewProduct(4, "The Great Gatsby", "BOOKS", 697.57),
			5: model.NewProduct(5, "Brib", "BABY", 366.90),
			6: model.NewProduct(6, "doll", "TOYS", 95.50),
			7: model.NewProduct(7, "Olive oil", "GROCESSORY", 302.19),
			8: model.NewProduct(8, "Ball", "TOYS", 295.37),
			9: model.NewProduct(9, "Frozen Food", "GROCESSORY", 534.64),
		},
	}
	return db
}
