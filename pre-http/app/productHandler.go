package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"pre-http/model"
	"pre-http/service"
	"strconv"
)

type ProductRoute struct {
	service service.ProductService
}

func (pr ProductRoute) getAllProducts(w http.ResponseWriter, r *http.Request) {
	// handle add user flow
	products := pr.service.GetAllProducts()
	writeResponse(w, products, http.StatusOK)
}

func (pr ProductRoute) saveProduct(w http.ResponseWriter, r *http.Request) {
	var request model.Product
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println("Not able to parse request: ", err.Error())

		errMessage := map[string]string{"message": err.Error()}

		writeResponse(w, errMessage, http.StatusBadRequest)
		return
	}
	newProduct := pr.service.AddProduct(request.Name, request.Category, request.Price)

	writeResponse(w, newProduct, http.StatusOK)
}

func (pr ProductRoute) registerRoutes(r *mux.Router) {
	r.HandleFunc("/products", pr.getAllProducts).Methods(http.MethodGet)
	r.HandleFunc("/products", pr.saveProduct).Methods(http.MethodPost)
	r.HandleFunc("/products/{id:[0-9]+}", pr.getProduct).Methods(http.MethodGet)
}

func (pr ProductRoute) getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	productId, err := strconv.Atoi(id)
	if err != nil {
		writeResponse(w, "invalid id", http.StatusBadRequest)
		return
	}
	product := pr.service.GetProductById(productId)
	writeResponse(w, product, http.StatusOK)
}

func writeResponse(w http.ResponseWriter, data any, code int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
