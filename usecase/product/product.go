package ucproduct

import (
	"context"
	"math"
	"sync"
	"time"
	"toped-scrapper/domain/product"
)

func (uc *Usecase) GetTokopediaProduct(ctx context.Context, params GetProductParam) ([]product.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// Channel to collect products
	productsCh := make(chan []product.Product)
	// Channel to communicate errors
	errCh := make(chan error)

	// WaitGroup to synchronize goroutines
	var wg sync.WaitGroup

	// Determining the number of pages to fetch based on the limit
	pagesToFetch := int(math.Ceil(float64(params.Limit) / float64(25)))
	for i := 0; i < pagesToFetch; i++ {
		wg.Add(1)
		go func(page int) {
			defer wg.Done()
			fetchedProducts, err := uc.productDomain.GetTokopediaProducts(ctx, product.TokopediaSearchParams{
				Query:     params.Category,
				Page:      page,
				SortOrder: "5",
			})
			if err != nil {
				errCh <- err
				return
			}
			productsCh <- fetchedProducts
		}(i + 1) // page starts from 1
	}

	// Close channels once all goroutines are done
	go func() {
		wg.Wait()
		close(productsCh)
		close(errCh)
	}()

	var products []product.Product
	for {
		select {
		case fetchedProducts := <-productsCh:
			products = append(products, fetchedProducts...)
			if len(products) >= int(params.Limit) {
				return products[:params.Limit], nil
			}
		case err := <-errCh:
			return nil, err
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}
