package main

import (
	"fmt"
	"intermediate-go/legacy/model"
	"intermediate-go/legacy/service"
)

func main() {

	// Printing all products
	printProducts(service.GetAllProducts())

	newlyAddedProduct := service.AddProduct("new product", "BOOKS", 109.99)

	// printing the newly added product
	printProduct(newlyAddedProduct.Id)

}

func printProducts(products []model.Product) {
	fmt.Println("### Printing Products ###")
	for _, p := range products {
		fmt.Println(p)
	}
	fmt.Println("### End ###")
}

func printProduct(id int) {
	if product := service.GetProductById(id); product != nil {
		fmt.Println("PRODUCT: ", product)
	} else {
		fmt.Println(fmt.Sprintf("Product with id %d does not exist", id))
	}
}
