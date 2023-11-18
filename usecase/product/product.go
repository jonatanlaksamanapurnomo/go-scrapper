package ucproduct

import (
	"context"
	"time"
	"toped-scrapper/domain/product"
)

func (uc *Usecase) GetTokopediaProduct(ctx context.Context, params GetProductParam) ([]product.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var products []product.Product
	productsChan := make(chan []product.Product)
	errChan := make(chan error)
	doneChan := make(chan bool)

	// Start 5 worker goroutines
	for i := 0; i < 5; i++ {
		go func(page int) {
			for {
				select {
				case <-ctx.Done():
					return // Exit the goroutine if context is cancelled
				case <-doneChan:
					return // Exit the goroutine if work is done
				default:
					tempProducts, err := uc.productDomain.GetTokopediaProducts(ctx, product.TokopediaSearchParams{
						Query:     params.Category,
						Page:      page,
						SortOrder: "5",
					})
					if err != nil {
						errChan <- err
						return
					}
					productsChan <- tempProducts
				}
			}
		}(i + 1)
	}

	// Collect results and handle errors
	remaining := params.Limit
	for remaining > 0 {
		select {
		case err := <-errChan:
			close(doneChan) // Stop all workers on error
			return nil, err
		case p := <-productsChan:
			products = append(products, p...)
			remaining -= len(p)
			if remaining <= 0 {
				close(doneChan) // Signal workers to stop
				break
			}
		}
	}

	// If we have more products than the limit, trim the slice to the limit
	if len(products) > params.Limit {
		products = products[:params.Limit]
	}

	return products, nil
}
