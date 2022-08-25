package main

import (
	"encoding/json"
	"log"
	"net/http"
	"pre-http/model"
	"pre-http/repository"
	"pre-http/service"
)

func manageProducts(w http.ResponseWriter, r *http.Request) {
	// handle add user flow
	ps := service.NewProductService(repository.NewInMemoryProductRepo())

	if r.Method == http.MethodPost {
		var request model.Product
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			log.Println("Not able to parse request: ", err.Error())
			w.WriteHeader(400)
			return
		}
		newProduct := ps.AddProduct(request.Name, request.Category, request.Price)

		writeResponse(w, newProduct)
	} else {
		products := ps.GetAllProducts()
		writeResponse(w, products)
	}
}

func writeResponse(w http.ResponseWriter, data any) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	http.HandleFunc("/products", manageProducts)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
