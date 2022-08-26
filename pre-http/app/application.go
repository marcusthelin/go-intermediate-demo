package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"pre-http/repository"
	"pre-http/service"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()

		next.ServeHTTP(w, r)

		log.Println(fmt.Sprintf("Total time %v for request %s: ", time.Since(t1), r.URL.Path))
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)
		log.Println("This is auth....")
	})
}

func Start() {
	r := mux.NewRouter()

	r.Use(loggingMiddleware)
	r.Use(authMiddleware)

	// wiring
	repo := repository.NewInMemoryProductRepo()
	productService := service.NewProductService(repo)
	pr := ProductRoute{service: *productService}

	//registerProductRoutes
	pr.registerRoutes(r)

	err := http.ListenAndServe("localhost:8080", r)
	log.Fatal(err)
}
