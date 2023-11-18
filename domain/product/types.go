package product

import (
	"context"
	"toped-scrapper/pkg/database/postgres"
)

type DomainItf interface {
	GetTokopediaProducts(ctx context.Context, params TokopediaSearchParams) ([]Product, error)
	InsertTokopediaProduct(ctx context.Context, product Product) error
}

type ResourceItf interface {
	GetTokopediaProducts(ctx context.Context, params TokopediaSearchParams) ([]Product, error)
	GetTokopediaProductsByURL(ctx context.Context, url string) ([]Product, error)
	InsertProductDB(ctx context.Context, product Product) error
}

type Domain struct {
	resource ResourceItf
}

type Resource struct {
	db postgres.PostgresHandler
}

type Product struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageLink   string `json:"image_link"`
	Rating      string `json:"rating"`
	Price       string `json:"price"`
	StoreName   string `json:"store_name"`
}

type TokopediaSearchParams struct {
	Query     string
	Page      int
	SortOrder string
}
