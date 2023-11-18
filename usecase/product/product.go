package ucproduct

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"sync"
	"time"
	"toped-scrapper/domain/product"
	"toped-scrapper/pkg/workerpool"
)

func (uc *Usecase) GetTokopediaProduct(ctx context.Context, params GetProductParam) ([]product.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var (
		products []product.Product
		mu       sync.Mutex
		pool     = workerpool.New(params.Worker) // Use a worker pool with 5 workers
	)

	page := 1
	productsPerPage := 25

	for len(products) < int(params.Limit) && page <= (int(params.Limit)/productsPerPage)+1 {
		currentPage := page
		pool.Submit(func(ctx context.Context) error {
			tempProducts, err := uc.productDomain.GetTokopediaProducts(ctx, product.TokopediaSearchParams{
				Query:     params.Category,
				Page:      currentPage,
				SortOrder: "5",
			})
			if err != nil {
				return err // Handle error appropriately
			}

			mu.Lock()
			defer mu.Unlock()

			for _, product := range tempProducts {
				products = append(products, product)
			}

			return nil
		})

		page++
	}

	// Run the worker pool
	if err := pool.Run(ctx); err != nil {
		return nil, err
	}

	// If more products were fetched than needed, trim the slice
	if len(products) > int(params.Limit) {
		products = products[:params.Limit]
	}

	for _, product := range products {
		if err := uc.productDomain.InsertTokopediaProduct(ctx, product); err != nil {
			continue
		}
	}
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
