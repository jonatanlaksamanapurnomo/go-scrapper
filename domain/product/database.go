package product

import (
	"context"
	"time"
)

func (rsc Resource) InsertProductDB(ctx context.Context, product Product) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	err := rsc.db.Execute(QueryInsertProduct, product.Name, product.Description, product.ImageLink, product.Rating, product.Price, product.StoreName)
	if err != nil {
		return err
	}

	return nil
}
