package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"pre-http/repository"
	"pre-http/service"
)

func Start() {
	r := mux.NewRouter()

	// wiring
	repo := repository.NewInMemoryProductRepo()
	productService := service.NewProductService(repo)
	pr := ProductRoute{service: *productService}

	//registerProductRoutes
	pr.registerRoutes(r)

	err := http.ListenAndServe("localhost:8080", r)
	log.Fatal(err)
}
