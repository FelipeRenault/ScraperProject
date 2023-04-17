package service

import (
	"PelandoChallenge/pkg/model"
	"PelandoChallenge/pkg/service/internal/repository"
	"PelandoChallenge/pkg/service/internal/scraper"
	"context"
	"fmt"
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
		fmt.Println("Retornando informações do DB")
		return product, nil
	}

	scrapedProduct := &model.Product{URL: url}
	err = scraper.Scrape(ctx, scrapedProduct)
	if err != nil {
		return nil, err
	}
	fmt.Println("Informações buscadas no site fonte")

	return scrapedProduct, p.productRepository.SaveProduct(ctx, scrapedProduct)
}
