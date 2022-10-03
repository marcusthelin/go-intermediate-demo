package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"pre-http-lab/repository"
	"pre-http-lab/service"

	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tNow := time.Now()
		next.ServeHTTP(w, r)
		fmt.Printf("\nTotal time %v for request %s", time.Since(tNow), r.URL.Path)
	})
}

func Start() {
	// Create a router
	r := mux.NewRouter()

	r.Use(loggingMiddleware)

	// Wire stuff together
	repo := repository.NewInMemoryProductRepo()
	pService := service.NewProductService(repo)
	pRoute := ProductRoute{service: *pService}
	pRoute.registerRoutes(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
