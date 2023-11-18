package product

import (
	"context"
)

func (d *Domain) GetTokopediaProducts(ctx context.Context, params TokopediaSearchParams) ([]Product, error) {
	products, err := d.resource.GetTokopediaProducts(ctx, params)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (d *Domain) InsertTokopediaProduct(ctx context.Context, product Product) error {
	// Call the InsertProductDB function from the Resource layer
	err := d.resource.InsertProductDB(ctx, product)
	if err != nil {
		return err // Propagate the error upward for handling
	}

	return nil
}
