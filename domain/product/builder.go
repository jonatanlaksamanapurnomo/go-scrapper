package product

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func parseHTML(ctx context.Context, html string) (resp []Product, err error) {
	var products []Product

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return resp, err
	}

	builder := &ConcreteProductBuilder{}

	content := doc.Find("div[data-testid='divSRPContentProducts']")
	content.Find(".css-qa82pd[data-testid='master-product-card']").Each(func(i int, s *goquery.Selection) {
		imageLink, _ := s.Find(".pcv3_img_container img").Attr("src")
		storeName := s.Find(".prd_link-shop-name").Text()
		rating := s.Find(".prd_rating-average-text").Text()

		builder.SetName(s.Find(".prd_link-product-name").Text()).
			SetPrice(s.Find(".prd_link-product-price").Text()).
			SetImageLink(imageLink).
			SetStoreName(storeName).
			SetRating(rating)
		product := builder.Build()
		products = append(products, product)
	})

	return products, nil
}
