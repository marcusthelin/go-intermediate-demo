package service

import (
	"reflect"
	"testing"

	"pre-http-lab/mocks/repository"
	"pre-http-lab/model"

	"github.com/golang/mock/gomock"
)

func TestAddProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := repository.NewMockProductRepository(ctrl)
	ps := ProductService{repo: mockRepo}

	expectedProduct := model.Product{
		Name:     "Hello World",
		Category: model.BOOKS,
		Price:    100.22,
	}

	mockRepo.EXPECT().Save(expectedProduct).Return(&expectedProduct)
	mockRepo.EXPECT().FindAll().Return([]model.Product{expectedProduct})

	ps.AddProduct("Hello World", model.BOOKS, 100.22)
	allProducts := ps.GetAllProducts()
	if len(allProducts) != 1 {
		t.Errorf("Expected product list should have 6 products, got %v", len(allProducts))
	}
}

func TestGetAllProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := repository.NewMockProductRepository(ctrl)
	ps := ProductService{repo: mockRepo}

	mockRepo.EXPECT().FindAll().MinTimes(1)

	if products := ps.GetAllProducts(); products != nil {
		if len(products) != 5 {
			t.Errorf("Expected products list should have 5 products")
		}
	}
}

func TestGetProductById(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := repository.NewMockProductRepository(ctrl)
	ps := ProductService{repo: mockRepo}

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
			mockRepo.EXPECT().FindBy(tt.id).Return(tt.want)
			if got := ps.GetProductById(tt.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProductById() = %v, want %v", got, tt.want)
			}
		})
	}
}
