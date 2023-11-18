package ucproduct

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"time"
	"toped-scrapper/domain/product"
)

func (uc *Usecase) GetTokopediaProduct(ctx context.Context, params GetProductParam) ([]product.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	fetcher := &ProductFetcher{
		productDomain: uc.productDomain,
		workerCount:   params.Worker,
	}

	products, err := fetcher.FetchProducts(ctx, params)
	if err != nil {
		return nil, err
	}

	inserter := &ProductInserter{
		productDomain: uc.productDomain,
		workerCount:   params.Worker,
	}

	inserter.InsertProducts(ctx, products)

	// Generate CSV after all products have been inserted and fetched
	return uc.GenerateCSV(products)
}

func (uc *Usecase) GenerateCSV(products []product.Product) ([]product.Product, error) {
	// Generating a timestamped filename
	timestamp := time.Now().Format("20060102-150405") // YYYYMMDD-HHMMSS format
	outputPath := "outputs"
	filename := fmt.Sprintf("%s/products-%s.csv", outputPath, timestamp)

	// Ensure the output directory exists
	err := os.MkdirAll(outputPath, os.ModePerm) // os.ModePerm is 0777, which means read, write & execute for everyone
	if err != nil {
		return nil, err
	}

	// Create a new CSV file with the timestamped name
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header
	header := []string{"Name", "Description", "ImageLink", "Rating", "Price", "StoreName"}
	if err := writer.Write(header); err != nil {
		return nil, err
	}

	// Write product data
	for _, p := range products {
		row := []string{p.Name, p.Description, p.ImageLink, p.Rating, p.Price, p.StoreName}
		if err := writer.Write(row); err != nil {
			return nil, err
		}
	}

	return products, nil
}
