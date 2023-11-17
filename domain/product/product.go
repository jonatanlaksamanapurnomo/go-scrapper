package product

import (
	"context"
)

func (d Domain) GetProduct(ctx context.Context, url string) (resp []Product, err error) {
	products, err := d.resource.GetTokopediaProducts(ctx, url)
	if err != nil {
		return resp, err
	}

	return products, nil
}
