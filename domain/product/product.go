package product

import (
	"context"
)

func (d Domain) GetTokopediaProducts(ctx context.Context, params TokopediaSearchParams) ([]Product, error) {
	products, err := d.resource.GetTokopediaProducts(ctx, params)
	if err != nil {
		return nil, err
	}

	return products, nil
}
