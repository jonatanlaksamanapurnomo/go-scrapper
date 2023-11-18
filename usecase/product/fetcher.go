package ucproduct

import (
	"context"
	"github.com/schollz/progressbar/v3"
	"sync"
	"toped-scrapper/domain/product"
	"toped-scrapper/pkg/workerpool"
)

type ProductFetcher struct {
	productDomain product.DomainItf
	workerCount   int
}

func (f *ProductFetcher) FetchProducts(ctx context.Context, params GetProductParam) ([]product.Product, error) {
	var allProducts []product.Product
	var totalFetched int
	mu := sync.Mutex{}
	pool := workerpool.New(f.workerCount)

	bar := progressbar.Default(int64(params.Limit))

	totalPages := (int(params.Limit) / 25) + 1
	productsToFetch := int(params.Limit)

	for page := 1; page <= totalPages; page++ {
		currentPage := page

		pool.Submit(func(ctx context.Context) error {
			if totalFetched >= productsToFetch {
				return nil // Skip if already fetched enough products
			}

			searchParams := product.TokopediaSearchParams{
				Query:     params.Category,
				Page:      currentPage,
				SortOrder: "5",
			}

			products, err := f.productDomain.GetTokopediaProducts(ctx, searchParams)
			if err != nil {
				return err
			}

			mu.Lock()
			defer mu.Unlock()

			// Determine the number of products to add
			productsToAdd := min(len(products), productsToFetch-totalFetched)
			allProducts = append(allProducts, products[:productsToAdd]...)
			totalFetched += productsToAdd
			bar.Add(productsToAdd)

			return nil
		})
	}

	if err := pool.Run(ctx); err != nil {
		return nil, err
	}

	bar.Finish()
	return allProducts, nil
}

// Helper function to find the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
