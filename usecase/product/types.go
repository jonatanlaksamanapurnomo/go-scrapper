package ucproduct

import (
	"context"
	"toped-scrapper/domain/product"
)

type GetProductParam struct {
	Category string `json:"category"`
	Limit    int    `json:"limit"`
	Worker   int    `json:"worker"`
}

type Fetcher interface {
	FetchProducts(ctx context.Context, params GetProductParam) ([]product.Product, error)
}

// Inserter defines the interface for inserting products
type Inserter interface {
	InsertProducts(ctx context.Context, products []product.Product) error
}
