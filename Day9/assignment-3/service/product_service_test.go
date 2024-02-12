package service

import (
	"belajar-mock/entity"
	"belajar-mock/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{Repository: productRepository}

func TestProductServiceGetOneProductNotFound(t *testing.T) {
	productRepository.Mock.On("FindById", "1").Return(nil)

	product, err := productService.GetOneProduct("1")

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "Product not found!", err.Error(), "Error response has to be 'product not found!'")
}

func TestProductServiceGetOneProductNotFounds(t *testing.T) {
	product := entity.Product{
		Id:   "2",
		Name: "Kaca mata",
	}

	productRepository.Mock.On("FindById", "2").Return(product)

	result, err := productService.GetOneProduct("2")

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, product.Id, result.Id, "result has to be '2")
	assert.Equal(t, product.Name, result.Name, "result has to be 'Kaca mata'")
	assert.Equal(t, &product, result, "resut has to be a product data with id '2'")
}

func TestFindAllProducts(t *testing.T) {
	products := []*entity.Product{
		{Id: "1", Name: "Product 1"},
		{Id: "2", Name: "Product 2"},
		{Id: "3", Name: "Product 3"},
	}
	productRepository.Mock.On("FindAll").Return(products)

	result, err := productService.GetAllProducts()

	assert.Nil(t, err)

	// assert.Equal(t, len(products), len(result)+1)
	assert.Equal(t, products, result)
}
