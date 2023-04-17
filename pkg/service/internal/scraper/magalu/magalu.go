package magalu

import (
	"PelandoChallenge/pkg/lib/currency"
	"PelandoChallenge/pkg/model"

	"github.com/gocolly/colly/v2"
)

const (
	BaseLink = "https://www.magazineluiza.com.br"

	titleTag       = "main > section > div[data-testid=mod-headingproduct] > h1"
	imageTag       = "main > section > div[data-testid=mod-mediagallery] > div > div > div > img[data-testid=image-selected-thumbnail]"
	priceTag       = "main > section > div[data-testid=mod-productprice] > div[data-testid=product-price] > div > div[data-testid=price-default] > p[data-testid=price-value]"
	descriptionTag = "main > section > div[data-testid=product-detail] > div[data-testid=product-detail-description] > div[data-testid=rich-content-container]"
)

func InfoCollector(product *model.Product, c *colly.Collector) {
	c.OnHTML(titleTag, func(e *colly.HTMLElement) {
		title := e.Text
		product.Title = title
	})

	c.OnHTML(imageTag, func(e *colly.HTMLElement) {
		img := e.Attr("src")
		product.Image = img
	})

	c.OnHTML(priceTag, func(e *colly.HTMLElement) {
		priceStr := e.Text

		price, _ := currency.ParsePrice(priceStr)
		product.Price = price
	})

	c.OnHTML(descriptionTag, func(e *colly.HTMLElement) {
		description := e.Text
		product.Description = description
	})
}
