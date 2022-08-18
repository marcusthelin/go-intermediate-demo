package service

import (
	"pre-http-lab/model"
	"pre-http-lab/repository"
	"reflect"
	"testing"
)

func TestAddProduct(t *testing.T) {
	ps := ProductService{repo: repository.NewInMemoryProductRepo()}
	newProduct := ps.AddProduct("new test product", model.TOYS, 100.23)

	if newProduct.Name != "new test product" {
		t.Errorf("Not the same product ")
	}

	if len(ps.GetAllProducts()) != 6 {
		t.Errorf("Expected product list should have 6 products")
	}
}

func TestGetAllProducts(t *testing.T) {
	ps := ProductService{repo: repository.NewInMemoryProductRepo()}

	if got := ps.GetAllProducts(); got != nil {
		if len(got) != 5 {
			t.Errorf("Expected products list should have 5 products")
		}
	}
}

func TestGetProductById(t *testing.T) {
	ps := ProductService{repo: repository.NewInMemoryProductRepo()}

	tests := []struct {
		name string
		id   int
		want *model.Product
	}{
		{
			"With product Id 1",
			1,
			&model.Product{Id: 1, Name: "InMem - Grocessory product", Category: model.GROCESSORY, Price: 100.0},
		},
		{
			"With product Id 2",
			2,
			&model.Product{Id: 2, Name: "InMem - Scala programming", Category: model.BOOKS, Price: 55.0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ps.GetProductById(tt.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProductById() = %v, want %v", got, tt.want)
			}
		})
	}
}
