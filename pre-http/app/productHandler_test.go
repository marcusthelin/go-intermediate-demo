package app

import (
	"net/http"
	"net/http/httptest"
	"pre-http/repository"
	"pre-http/service"
	"testing"
)

func Test_get_all_products_ok(t *testing.T) {
	repo := repository.NewInMemoryProductRepo()
	productService := service.NewProductService(repo)
	pr := ProductRoute{service: *productService}

	//router := mux.NewRouter()
	//router.HandleFunc("/products", pr.getAllProducts)

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/products", nil)

	//router.ServeHTTP(recorder, req)

	pr.getAllProducts(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Error("unexpected status code, expected 200 and got 404")
	}
}
