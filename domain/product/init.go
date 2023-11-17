package product

import "context"

type DomainItf interface {
	GetProduct(ctx context.Context, url string) (resp []Product, err error)
}

type ResourceItf interface {
	GetTokopediaProducts(ctx context.Context, url string) (resp []Product, err error)
}

func InitDomain(rsc ResourceItf) *Domain {
	return &Domain{
		resource: rsc,
	}
}

func InitResource() Resource {
	return Resource{}
}
