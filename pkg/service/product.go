package service

import (
	"PelandoChallenge/pkg/model"
	"PelandoChallenge/pkg/service/internal/repository"
	"context"
)

type ProductInterface interface {
	GetProduct(ctx context.Context, url string) (*model.Product, error)
}

type product struct {
	productRepository repository.Product
}

func (p *product) GetProduct(ctx context.Context, url string) (*model.Product, error) {
	product, err := p.productRepository.GetProductByURL(ctx, url)
	if err != nil {
		return nil, err
	}

	if product != nil && product.IsWithinTimeThreshold() {
		return product, nil
	}

	// Fetch from website

	return nil, nil
}
