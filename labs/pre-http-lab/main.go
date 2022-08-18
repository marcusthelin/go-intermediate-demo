package main

import (
	"fmt"
	"pre-http-lab/model"
	"pre-http-lab/repository"
	"pre-http-lab/service"
)

func main() {
	// creating instance of product service
	ps := service.NewProductService(repository.NewInMemoryProductRepo())

	// Printing all products
	printProducts(ps.GetAllProducts())

	newlyAddedProduct := ps.AddProduct("new product", model.BOOKS, 109.99)

	// printing the newly added product
	printProduct(ps, newlyAddedProduct.Id)

}

func printProducts(products []model.Product) {
	fmt.Println("### Printing Products ###")
	for _, p := range products {
		fmt.Println(p)
	}
	fmt.Println("### End ###")
}

func printProduct(ps *service.ProductService, id int) {
	if product := ps.GetProductById(id); product != nil {
		fmt.Println("PRODUCT: ", product)
	} else {
		fmt.Println(fmt.Sprintf("Product with id %d does not exist", id))
	}
}
