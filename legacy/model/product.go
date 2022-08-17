package model

import "fmt"

type Product struct {
	Id          int
	Name        string
	ProductType Category
	Price       float32
}

func (p Product) String() string {
	return fmt.Sprintf("{ id:%d, name:\"%s\", productType:\"%s\", price:%5.2f }", p.Id, p.Name, p.ProductType, p.Price)
}

func NewProduct(id int, name string, productType Category, price float32) Product {
	return Product{id, name, productType, price}
}
