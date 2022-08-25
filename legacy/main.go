package main

import (
	"fmt"
	"legacy/model"
	"legacy/repository"
	"legacy/service"
	"log"
	"os"
)

func main() {
	//configuration code
	var repo repository.ProductRepository

	if repo = getProductRepository(); repo == nil {
		log.Fatal("Environment not defined...")
	}

	//inMem := repository.NewInMemoryProductRepo()
	// wiring the application
	//db := repository.NewDbProductRepo()
	ps := service.NewProductService(repo)

	// Printing all products
	products := ps.GetAllProducts()

	printProducts(products)

	newlyAddedProduct := ps.AddProduct("new product", model.BOOKS, 109.99)

	// printing the newly added product
	printProduct(newlyAddedProduct.Id, ps)

}

func getProductRepository() repository.ProductRepository {
	if os.Getenv("ENV") == "PROD" {
		return repository.NewDbProductRepo()
	}
	if os.Getenv("ENV") == "DEV" {
		return repository.NewInMemoryProductRepo()
	}
	return nil
	//log.Fatal("Environment not defined...")
}

func printProducts(products []model.Product) {
	fmt.Println("### Printing Products ###")
	for _, p := range products {
		fmt.Println(p)
	}
	fmt.Println("### End ###")
}

func printProduct(id int, ps service.ProductService) {
	if product := ps.GetProductById(id); product != nil {
		fmt.Println("PRODUCT: ", product)
	} else {
		fmt.Println(fmt.Sprintf("Product with id %d does not exist", id))
	}
}
