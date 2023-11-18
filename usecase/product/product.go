package ucproduct

import (
	"context"
	"time"
	"toped-scrapper/domain/product"
)

func (uc *Usecase) GetTokopediaProduct(ctx context.Context, params product.TokopediaSearchParams) (resp []product.Product, err error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	resp, err = uc.productDomain.GetTokopediaProducts(ctx, params)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
