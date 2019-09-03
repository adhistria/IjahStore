package service

import (
	"context"
	"time"

	"github.com/adhistria/ijahstore/model"
	"github.com/adhistria/ijahstore/repository"
)

type (
	IncomingProductService interface {
		AddIncomingProduct(context.Context, *model.IncomingProduct) (*model.IncomingProduct, error)
	}

	incomingProductServiceImpl struct {
		productRepository         repository.ProductRepository
		incomingProductRepository repository.IncomingProductRepository
	}
)

func (ips *incomingProductServiceImpl) AddIncomingProduct(ctx context.Context, incomingProduct *model.IncomingProduct) (*model.IncomingProduct, error) {

	incomingProduct.SetTotalPurchasePrice()

	err := ips.incomingProductRepository.Add(incomingProduct)
	if err != nil {
		return nil, err
	}
	err = ips.productRepository.AddTotalProduct(incomingProduct)

	if err != nil {
		return nil, err
	}
	product, err := ips.productRepository.GetProduct(incomingProduct.ProductID)

	if err != nil {
		return nil, err
	}
	incomingProduct.Product = product
	incomingProduct.CreatedAt = time.Now()

	return incomingProduct, err
}

func NewIncomingProductService(pr repository.ProductRepository, ipr repository.IncomingProductRepository) IncomingProductService {
	return &incomingProductServiceImpl{productRepository: pr, incomingProductRepository: ipr}
}
