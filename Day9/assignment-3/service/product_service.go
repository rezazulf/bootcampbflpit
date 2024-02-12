package service

import (
	"belajar-mock/entity"
	"belajar-mock/repository"
	"errors"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (service ProductService) GetOneProduct(id string) (*entity.Product, error) {
	product := service.Repository.FindById(id)
	if product == nil {
		return nil, errors.New("Product not found!")
	}
	return product, nil
}

// func (service ProductService) GetAllProducts() ([]*entity.Product, error) {
// 	products := service.Repository.FindAll()
// 	if products == nil {
// 		return nil, errors.New("no products found")
// 	}
// 	return products, nil
// }

func (service ProductService) GetAllProducts() ([]*entity.Product, error) {
	return service.Repository.FindAll(), nil
}
