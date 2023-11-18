package ucproduct

import (
	"context"
	"toped-scrapper/domain/product"
)

type ProductInserter struct {
	productDomain product.DomainItf
	workerCount   int
}

func (p *ProductInserter) InsertProducts(ctx context.Context, products []product.Product) error {
	var errorInLoop error
	for _, product := range products {
		if err := p.productDomain.InsertTokopediaProduct(ctx, product); err != nil {
			errorInLoop = err
		}
	}
	return errorInLoop
}
