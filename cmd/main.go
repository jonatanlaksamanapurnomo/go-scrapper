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
	products, err := productDomain.GetProduct(context.Background(), "https://www.tokopedia.com/search?q=handphone&source=search&srp_component_id=04.06.00.00&st=product")
	fmt.Print(fmt.Sprintf("%+v", products[0]), len(products), err)
}

func initDomains() {
	productDomain = product.InitDomain(product.InitResource())
}
