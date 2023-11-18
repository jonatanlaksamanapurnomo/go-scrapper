package product

import "context"

type DomainItf interface {
	GetTokopediaProducts(ctx context.Context, params TokopediaSearchParams) ([]Product, error)
}

type ResourceItf interface {
	GetTokopediaProducts(ctx context.Context, params TokopediaSearchParams) ([]Product, error)
	GetTokopediaProductsByURL(ctx context.Context, url string) ([]Product, error)
}

type Domain struct {
	resource ResourceItf
}

type Resource struct {
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
	Page      int64
	SortOrder string
}
