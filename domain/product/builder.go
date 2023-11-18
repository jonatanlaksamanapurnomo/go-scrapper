package product

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type ProductBuilder interface {
	SetName(name string) ProductBuilder
	SetDescription(description string) ProductBuilder
	SetImageLink(imageLink string) ProductBuilder
	SetRating(rating string) ProductBuilder
	SetPrice(price string) ProductBuilder
	SetStoreName(storeName string) ProductBuilder
	Build() Product
}

type ConcreteProductBuilder struct {
	product Product
}

func (b *ConcreteProductBuilder) SetName(name string) ProductBuilder {
	b.product.Name = name
	return b
}

func (b *ConcreteProductBuilder) SetDescription(description string) ProductBuilder {
	b.product.Description = description
	return b
}

func (b *ConcreteProductBuilder) SetImageLink(imageLink string) ProductBuilder {
	b.product.ImageLink = imageLink
	return b
}

func (b *ConcreteProductBuilder) SetRating(rating string) ProductBuilder {
	b.product.Rating = rating
	return b
}

func (b *ConcreteProductBuilder) SetPrice(price string) ProductBuilder {
	b.product.Price = price
	return b
}

func (b *ConcreteProductBuilder) SetStoreName(storeName string) ProductBuilder {
	b.product.StoreName = storeName
	return b
}

func (b *ConcreteProductBuilder) Build() Product {
	return b.product
}

func buildProductFromHTML(html string) (resp []Product, err error) {
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
