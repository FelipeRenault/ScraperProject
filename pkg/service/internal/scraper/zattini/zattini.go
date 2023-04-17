package zattini

import (
	"PelandoChallenge/pkg/lib/currency"
	"PelandoChallenge/pkg/model"

	"github.com/gocolly/colly/v2"
)

const (
	BaseLink = "https://www.zattini.com.br"

	titleTag       = "main > div > section > div[class=short-showcase-description] > section[class=short-description] > h1[data-productname]"
	imageTag       = "main > div > section > section[class=photo] > figure[class=photo-figure] > img[class=zoom]"
	priceTag       = "main > div > section > section[id=buy-box] > div > div > div[class=default-price] > span > strong"
	descriptionTag = "main > section[id=features] > p[itemprop=description]"
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
