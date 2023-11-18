package main

import (
	"context"
	"fmt"
	"github.com/manifoldco/promptui"
	"log"
	"strconv"
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

	// Interactive prompts
	category := promptForInput("Enter product category (e.g., handphone): ")
	limit := promptForNumber("Enter limit of products to fetch (e.g., 100): ")
	worker := promptForNumber("Enter number of workers (e.g., 5): ")

	products, err := productUsecase.GetTokopediaProduct(context.Background(), ucproduct.GetProductParam{
		Category: category,
		Limit:    limit,
		Worker:   worker,
	})
	if err != nil {
		log.Fatalf("Error fetching products: %v", err)
	}

	fmt.Printf("\nFetched %d products\n", len(products))
	fmt.Printf("Check outputs folder for csv \n")

	for _, p := range products {
		fmt.Printf("Name: %s, Price: %s\n", p.Name, p.Price)
		// Add more details if needed
	}
}

func promptForInput(label string) string {
	prompt := promptui.Prompt{
		Label: label,
	}
	result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	return result
}

func promptForNumber(label string) int {
	validate := func(input string) error {
		_, err := strconv.Atoi(input)
		if err != nil {
			return fmt.Errorf("invalid number")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
	}

	result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	number, _ := strconv.Atoi(result)
	return number
}

func Init() {
	dbHandler, err := postgres.NewPostgresHandler("host=localhost port=5432 user=postgre password=mysecretpassword dbname=go_scrapper sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	err = dbHandler.CheckAndCreateTable()
	if err != nil {
		log.Fatal(err)
	}
	initDomains(*dbHandler)
	initUsecase()
}

func initDomains(db postgres.PostgresHandler) {
	productDomain = product.InitDomain(product.InitResource(db))
}

func initUsecase() {
	productUsecase = ucproduct.InitUsecase(productDomain)
}
