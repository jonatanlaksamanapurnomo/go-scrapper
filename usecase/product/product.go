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
	page := 1

	for len(products) < int(params.Limit) {
		tempProducts, err := uc.productDomain.GetTokopediaProducts(ctx, product.TokopediaSearchParams{
			Query:     params.Category,
			Page:      page,
			SortOrder: "5",
		})
		if err != nil {
			return nil, err // Stop and return error if unable to fetch products
		}

		// Append products until the limit is reached or exceeded
		products = append(products, tempProducts...)
		if len(products) >= int(params.Limit) {
			break
		}
		page++
	}

	// If we have more products than the limit, trim the slice to the limit
	if len(products) > int(params.Limit) {
		products = products[:params.Limit]
	}

	return products, nil
}
