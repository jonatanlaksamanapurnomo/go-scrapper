package ucproduct

import (
	"context"
	"sync"
	"toped-scrapper/domain/product"
	"toped-scrapper/pkg/workerpool"
)

type TokopediaFetcher struct {
	productDomain product.DomainItf
	workerCount   int
}

func (f *TokopediaFetcher) FetchProducts(ctx context.Context, params GetProductParam) ([]product.Product, error) {
	var (
		allProducts []product.Product
		mu          sync.Mutex
		pool        = workerpool.New(f.workerCount) // Initialize the worker pool with the given number of workers
	)

	productsPerPage := 25 // Assumed number of products per page
	totalPageCount := (int(params.Limit) / productsPerPage) + 1

	for page := 1; page <= totalPageCount; page++ {
		currentPage := page // Capture the current page for the closure
		pool.Submit(func(ctx context.Context) error {
			searchParams := product.TokopediaSearchParams{
				Query:     params.Category,
				Page:      currentPage,
				SortOrder: "5", // Assuming "5" is the sort order you want to use
			}

			products, err := f.productDomain.GetTokopediaProducts(ctx, searchParams)
			if err != nil {
				return err // Handle error appropriately
			}

			mu.Lock()
			allProducts = append(allProducts, products...)
			mu.Unlock()

			return nil
		})
	}

	// Wait for all the workers to finish
	if err := pool.Run(ctx); err != nil {
		return nil, err
	}

	// Trim the results if more products were fetched than needed
	if len(allProducts) > int(params.Limit) {
		allProducts = allProducts[:params.Limit]
	}

	return allProducts, nil
}
