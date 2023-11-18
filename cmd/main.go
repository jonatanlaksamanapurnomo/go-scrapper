package main

import (
	"context"
	"fmt"
	"toped-scrapper/domain/product"
	ucproduct "toped-scrapper/usecase/product"
)

var (
	productDomain  *product.Domain
	productUsecase *ucproduct.Usecase
)

func main() {
	initDomains()
	initUsecase()

	products, err := productUsecase.GetTokopediaProduct(context.Background(), product.TokopediaSearchParams{
		Query:     "handphone",
		Page:      1,
		SortOrder: "5",
	})
	fmt.Println(len(products), err)
}

func initDomains() {
	productDomain = product.InitDomain(product.InitResource())
}

func initUsecase() {
	productUsecase = ucproduct.InitUsecase(productDomain)
}
