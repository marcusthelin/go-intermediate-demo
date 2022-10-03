package app

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"pre-http-lab/repository"
	"pre-http-lab/service"

	"github.com/gorilla/mux"
)

func Test_get_all_products(t *testing.T) {
	repo := repository.NewInMemoryProductRepo()
	pService := service.NewProductService(repo)
	pr := ProductRoute{
		service: *pService,
	}

	router := mux.NewRouter()
	router.HandleFunc("/products", pr.getAllProducts)

	recorder := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	router.ServeHTTP(recorder, req)

	fmt.Println(recorder.Code)

	if recorder.Code != http.StatusOK {
		t.Error("Not an ok status code returned")
	}
}
