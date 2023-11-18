package ucproduct

import (
	"context"
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
		pool     = workerpool.New(5) // Use a worker pool with 5 workers
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
			products = append(products, tempProducts...)
			mu.Unlock()

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

	return products, nil
}
