package main

import (
	"toped-scrapper/domain/product"
)

var (
	productDomain product.DomainItf
)

func main() {
	initDomains()
}

func initDomains() {
	productDomain = product.InitDomain(product.InitResource())
}
