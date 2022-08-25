package main

import (
	"encoding/json"
	"log"
	"net/http"
	"pre-http/repository"
	"pre-http/service"
)

func getAllProducts(w http.ResponseWriter, r *http.Request) {
	ps := service.NewProductService(repository.NewInMemoryProductRepo())
	products := ps.GetAllProducts()

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func main() {
	http.HandleFunc("/products", getAllProducts)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
