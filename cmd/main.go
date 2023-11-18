package main

import (
	"context"
	"fmt"
	"toped-scrapper/domain/product"
)

var (
	productDomain product.DomainItf
)

func main() {
	initDomains()
	products, err := productDomain.GetTokopediaProducts(context.Background(), product.TokopediaSearchParams{
		Query:     "handphone",
		Page:      1,
		SortOrder: "5",
	})
	fmt.Print(products[0], err)
}

func initDomains() {
	productDomain = product.InitDomain(product.InitResource())
}
