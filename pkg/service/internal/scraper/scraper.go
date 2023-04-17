package scraper

import (
	"PelandoChallenge/pkg/lib/collector"
	"PelandoChallenge/pkg/lib/errors"
	"PelandoChallenge/pkg/model"
	"PelandoChallenge/pkg/service/internal/scraper/magalu"
	"PelandoChallenge/pkg/service/internal/scraper/zattini"
	"context"
	"strings"
)

func Scrape(ctx context.Context, product *model.Product) error {
	c, err := collector.NewCollector()
	if err != nil {
		return err
	}

	url := product.URL
	switch {
	case strings.HasPrefix(url, magalu.BaseLink):
		magalu.InfoCollector(product, c)
	case strings.HasPrefix(url, zattini.BaseLink):
		zattini.InfoCollector(product, c)
	default:
		return errors.ErrSiteNotAvailable
	}

	c.Visit(url)
	return nil
}
