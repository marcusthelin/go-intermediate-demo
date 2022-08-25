package main

import (
	"encoding/json"
	"log"
	"net/http"
	"pre-http/model"
	"pre-http/repository"
	"pre-http/service"
)

type Struct1 struct {
	Name string
}

type Struct2 struct {
	Struct1 Struct1 `json:"some_struct"`
	City    string  `json:"my_city"`
}

func manageProducts(w http.ResponseWriter, r *http.Request) {
	// handle add user flow
	ps := service.NewProductService(repository.NewInMemoryProductRepo())

	if r.Method == http.MethodPost {
		var request model.Product
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			log.Println("Not able to parse request: ", err.Error())
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}
		newProduct := ps.AddProduct(request.Name, request.Category, request.Price)

		writeResponse(w, newProduct)
	} else {
		//products := ps.GetAllProducts()
		struct2 := Struct2{
			Struct1: Struct1{Name: "test"},
			City:    "Delhi",
		}
		writeResponse(w, struct2)
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
