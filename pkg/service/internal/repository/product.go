package repository

import (
	"context"
	"database/sql"
	"errors"

	"PelandoChallenge/pkg/model"

	"github.com/go-gorp/gorp"
)

type Product interface {
	GetProductByURL(ctx context.Context, url string) (*model.Product, error)
	SaveProduct(ctx context.Context, product *model.Product) error
}

func NewProduct(db *gorp.DbMap) Product {
	return &productRepository{
		db,
	}
}

type productRepository struct {
	db *gorp.DbMap
}

var (
	execSaveProduct = `
		INSERT INTO pelando.products (url, title, image, price, description)
		VALUES (?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
			title = VALUES(title),
			image = VALUES(image),
			price = VALUES(price),
			description = VALUES(description);
	`

	queryGetProduct = `
		SELECT
			id,
			url,
			title,
			image,
			price,
			description,
			created_at,
			updated_at
		FROM pelando.products
		WHERE url = ?
		LIMIT 1;
	`
)

func (p *productRepository) GetProductByURL(ctx context.Context, url string) (*model.Product, error) {
	var product model.Product

	err := p.db.SelectOne(&product, queryGetProduct, url)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *productRepository) SaveProduct(ctx context.Context, product *model.Product) error {
	_, err := p.db.Exec(execSaveProduct, product.URL, product.Title)
	return err
}
