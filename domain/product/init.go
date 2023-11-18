package product

import "context"

type DomainItf interface {
	GetTokopediaProducts(ctx context.Context, params TokopediaSearchParams) ([]Product, error)
}

type ResourceItf interface {
	GetTokopediaProducts(ctx context.Context, params TokopediaSearchParams) ([]Product, error)
	GetTokopediaProductsByURL(ctx context.Context, url string) ([]Product, error)
}

func InitDomain(rsc ResourceItf) *Domain {
	return &Domain{
		resource: rsc,
	}
}

func InitResource() Resource {
	return Resource{}
}
