package service

import (
	"context"

	"PelandoChallenge/pkg/lib/db"
	"PelandoChallenge/pkg/service/internal/repository"
)

type Service struct {
	Product ProductInterface
}

func New(ctx context.Context) (*Service, error) {
	db, err := db.Connect()
	if err != nil {
		return nil, err
	}

	productRepository := repository.NewProduct(db)

	productService := &product{
		productRepository: productRepository,
	}

	return &Service{
		Product: productService,
	}, nil
}
