package service

import (
	"context"
	"time"

	"github.com/adhistria/ijahstore/model"
	"github.com/adhistria/ijahstore/repository"
)

type (
	OutgoingProductService interface {
		AddOutgoingProduct(context.Context, *model.OutgoingProduct) (*model.OutgoingProduct, error)
		GetOutgoingProducts(context.Context) ([]model.OutgoingProduct, error)
	}

	outgoingProductServiceImpl struct {
		productRepository         repository.ProductRepository
		OutgoingProductRepository repository.OutgoingProductRepository
	}
)

func (ops *outgoingProductServiceImpl) AddOutgoingProduct(ctx context.Context, outgoingProduct *model.OutgoingProduct) (*model.OutgoingProduct, error) {

	outgoingProduct.SetTotalSellingPrice()

	err := ops.OutgoingProductRepository.Add(outgoingProduct)
	if err != nil {
		return nil, err
	}
	err = ops.productRepository.SubstractTotalProduct(outgoingProduct)

	if err != nil {
		return nil, err
	}
	product, err := ops.productRepository.GetProduct(outgoingProduct.ProductID)

	if err != nil {
		return nil, err
	}
	outgoingProduct.Product = product
	outgoingProduct.CreatedAt = time.Now()

	return outgoingProduct, err
}

func (ops *outgoingProductServiceImpl) GetOutgoingProducts(ctx context.Context) ([]model.OutgoingProduct, error) {
	var outgoingProducts []model.OutgoingProduct
	outgoingProducts, err := ops.OutgoingProductRepository.FindAll()

	for index, _ := range outgoingProducts {
		outgoingProducts[index].Product, err = ops.productRepository.GetProduct(outgoingProducts[index].ProductID)
		if err != nil {
			return nil, err
		}
	}
	return outgoingProducts, err
}

func NewOutgoingProductService(pr repository.ProductRepository, opr repository.OutgoingProductRepository) OutgoingProductService {
	return &outgoingProductServiceImpl{productRepository: pr, OutgoingProductRepository: opr}
}
