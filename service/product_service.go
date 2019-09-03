package service

import (
	"context"

	"github.com/adhistria/ijahstore/model"
	"github.com/adhistria/ijahstore/repository"
)

type (
	ProductService interface {
		AddProduct(context.Context, *model.Product) (*model.Product, error)
		GetProducts(context.Context) ([]model.Product, error)
	}

	productServiceImpl struct {
		productRepository repository.ProductRepository
	}
)

func (productService *productServiceImpl) AddProduct(ctx context.Context, product *model.Product) (*model.Product, error) {
	err := productService.productRepository.Add(product)
	return product, err
}

func (productService *productServiceImpl) GetProducts(ctx context.Context) ([]model.Product, error) {
	var products []model.Product
	products, err := productService.productRepository.FindAll()
	return products, err
}

func NewProductService(pr repository.ProductRepository) ProductService {
	return &productServiceImpl{productRepository: pr}
}
