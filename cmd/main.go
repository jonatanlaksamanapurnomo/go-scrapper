package main

import (
	"context"
	"fmt"
	"log"
	"toped-scrapper/domain/product"
	"toped-scrapper/pkg/database/postgres"
	ucproduct "toped-scrapper/usecase/product"
)

var (
	productDomain  *product.Domain
	productUsecase *ucproduct.Usecase
)

func main() {
	Init()
	products, err := productUsecase.GetTokopediaProduct(context.Background(), ucproduct.GetProductParam{
		Category: "handphone",
		Limit:    100,
		Worker:   5,
	})
	fmt.Println(len(products), err)
}

func Init() {
	dbHandler, err := postgres.NewPostgresHandler("host=localhost port=5432 user=postgre password=mysecretpassword dbname=go_scrapper sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer dbHandler.Close()

	initDomains()
	initUsecase()
}

func initDomains() {
	productDomain = product.InitDomain(product.InitResource())
}

func initUsecase() {
	productUsecase = ucproduct.InitUsecase(productDomain)
}
