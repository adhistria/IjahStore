package service

import (
	"context"

	"github.com/adhistria/ijahstore/model"
	"github.com/adhistria/ijahstore/repository"
)

type (
	ProductService interface {
		AddProduct(context.Context, *model.Product) (*model.Product, error)
	}

	productServiceImpl struct {
		productRepository repository.ProductRepository
	}
)

func (productService *productServiceImpl) AddProduct(ctx context.Context, product *model.Product) (*model.Product, error) {
	err := productService.productRepository.Add(product)
	return product, err
}

func NewProductService(pr repository.ProductRepository) ProductService {
	return &productServiceImpl{productRepository: pr}
}
